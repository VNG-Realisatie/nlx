// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package monitor

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-monitor/health"
)

const monitorHTTPTimeout = 30 * time.Second

// HealthChecker checks the inways of a StoredService and modifies it's health state directly in the StoredService struct.
type HealthChecker struct {
	logger     *zap.Logger
	httpClient *http.Client

	availabilitiesLock sync.RWMutex
	availabilities     map[uint64]*availability

	stmtSelectAvailabilities  *sqlx.Stmt
	stmtUpdateHealth          *sqlx.Stmt
	stmtCleanUpAvailabilities *sqlx.Stmt
	stmtUpdateInwayVersion    *sqlx.Stmt

	listener                     *pq.Listener
	shutdownNotificationListener chan struct{}
	shutdown                     chan struct{}
}

type availability struct {
	ID                       uint64 `json:"id"`
	OrganizationSerialNumber string `json:"organization_serial_number"`
	ServiceName              string `json:"service_name"`
	Address                  string `json:"address"`
	InwayID                  uint64 `json:"inway_id"`
}

func (a *availability) healthCheckURL() string {
	return fmt.Sprintf("https://%s/.nlx/health/%s", a.Address, a.ServiceName)
}

type databaseAction struct {
	Action       string        `json:"action"`
	Availability *availability `json:"availability"`
}

const (
	dbNotificationChannel              = "availabilities"
	cleanupPostgresConnectionsInterval = 90 * time.Second
	cleanupOfflineServicesInterval     = 5 * time.Minute

	healthCheckInterval = 1 * time.Minute

	postgresListenerReconnectTimeoutMin = 10 * time.Second
	postgresListenerReconnectTimeoutMax = 1 * time.Minute

	idleConnectionTimeout = 1 * time.Second
)

func New(logger *zap.Logger, certificate *common_tls.CertificateBundle) *HealthChecker {
	return &HealthChecker{
		logger:         logger,
		availabilities: make(map[uint64]*availability),
		httpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: certificate.TLSConfig(),
				IdleConnTimeout: idleConnectionTimeout,
			},
			Timeout: monitorHTTPTimeout,
		},
	}
}

func (h *HealthChecker) Run(
	db *sqlx.DB,
	postgresDNS string,
	ttlOfflineService int,
) error {
	var err error
	h.stmtSelectAvailabilities, err = db.Preparex(`
		SELECT
			availabilities.id,
			availabilities.inway_id,
			organizations.serial_number AS organization_serial_number,
			services.name AS service_name,
			inways.address
		FROM directory.availabilities
			INNER JOIN directory.inways
				ON availabilities.inway_id = inways.id
			INNER JOIN directory.services
				ON availabilities.service_id = services.id
			INNER JOIN directory.organizations
				ON services.organization_id = organizations.id
	`)

	if err != nil {
		return errors.Wrap(err, "failed to prepare stmtSelectAvailabilities")
	}

	h.stmtUpdateHealth, err = db.Preparex(`
		UPDATE
			directory.availabilities
		SET
			healthy = $2,
			unhealthy_since =
				CASE WHEN $2 = false THEN
					NOW()
				ELSE
					NULL
				END
		WHERE
			id = $1 AND healthy != $2`)
	if err != nil {
		return errors.Wrap(err, "failed to prepare stmtUpdateHealth")
	}

	h.stmtUpdateInwayVersion, err = db.Preparex(`
		UPDATE
			directory.inways
		SET
			version = $2
		WHERE id = $1
	`)
	if err != nil {
		return errors.Wrap(err, "failed to prepare stmtUpdateInwayVersion")
	}

	h.stmtCleanUpAvailabilities, err = db.Preparex(fmt.Sprintf(`
		DELETE FROM directory.availabilities
		WHERE NOW() - INTERVAL '%d seconds' > availabilities.last_announced`,
		ttlOfflineService))
	if err != nil {
		return errors.Wrap(err, "failed to prepare stmtCleanUpAvailabilities")
	}

	return h.run(postgresDNS)
}

func (h *HealthChecker) run(postgressDNS string) error {
	listenerErrorCallback := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			h.logger.Error("error listening for db events", zap.Error(err))
		}
	}
	listener := pq.NewListener(postgressDNS, postgresListenerReconnectTimeoutMin, postgresListenerReconnectTimeoutMax, listenerErrorCallback)

	err := listener.Listen(dbNotificationChannel)
	if err != nil {
		return err
	}

	h.listener = listener

	err = h.loadAvailabilities()
	if err != nil {
		return err
	}

	shutdownNotificationListener := make(chan struct{})
	h.shutdownNotificationListener = shutdownNotificationListener

	go h.waitForNotification(listener, shutdownNotificationListener)

	shutdown := make(chan struct{})
	h.shutdown = shutdown

	go h.runCleanUpServices(shutdown)
	h.runHealthChecker(shutdown)

	return nil
}

func (h *HealthChecker) loadAvailabilities() error {
	newAvailabilities := []*availability{}

	err := h.stmtSelectAvailabilities.Select(&newAvailabilities)
	if err != nil {
		return err
	}

	h.availabilitiesLock.Lock()
	for _, availability := range newAvailabilities {
		h.availabilities[availability.ID] = availability
	}
	h.availabilitiesLock.Unlock()

	return nil
}

func (h *HealthChecker) Shutdown() error {
	h.listener.Close()
	close(h.shutdownNotificationListener)
	close(h.shutdown)

	return nil
}

func (h *HealthChecker) runCleanUpServices(shutdown chan struct{}) {
	h.logger.Debug("initial cleaning up stale services")

	servicesRemoved, err := h.cleanUpServices()
	if err != nil {
		h.logger.Error("error cleaning up offline services", zap.Error(err))
	}

	h.logger.Debug("cleanup complete", zap.Int64("services removed", servicesRemoved))

	for {
		select {
		case <-time.After(cleanupOfflineServicesInterval):
			h.logger.Debug("cleaning up offline services")

			servicesRemoved, err := h.cleanUpServices()
			if err != nil {
				h.logger.Error("error cleaning up offline services", zap.Error(err))
				continue
			}

			h.logger.Debug("cleanup complete", zap.Int64("services removed", servicesRemoved))
		case <-shutdown:
			return
		}
	}
}

func (h *HealthChecker) cleanUpServices() (int64, error) {
	res, err := h.stmtCleanUpAvailabilities.Exec()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (h *HealthChecker) waitForNotification(l *pq.Listener, c <-chan struct{}) {
	for {
		select {
		case n := <-l.Notify:
			if n != nil {
				h.onDatabaseNotification(n.Extra)
			}
		case <-time.After(cleanupPostgresConnectionsInterval):
			// Check connection after 90 seconds without a notification
			go func() {
				if err := l.Ping(); err != nil {
					h.logger.Error("error pinging DB listener", zap.Error(err))
				}
			}()
		case <-c:
			return
		}
	}
}

func (h *HealthChecker) onDatabaseNotification(payload string) {
	dbAction := &databaseAction{}

	err := json.Unmarshal([]byte(payload), dbAction)
	if err != nil {
		h.logger.Error("Error processing JSON", zap.Error(err))
		return
	}

	h.logger.Debug("received DB action", zap.String("action", dbAction.Action))

	switch dbAction.Action {
	case "INSERT", "UPDATE":
		h.addAvailability(dbAction.Availability)
	case "DELETE":
		h.removeAvailability(dbAction.Availability.ID)
	default:
		h.logger.Error("unknown database action", zap.String("database action", dbAction.Action))
	}
}

func (h *HealthChecker) addAvailability(a *availability) {
	h.availabilitiesLock.Lock()
	h.availabilities[a.ID] = a
	h.availabilitiesLock.Unlock()
}

func (h *HealthChecker) removeAvailability(id uint64) {
	h.availabilitiesLock.Lock()
	delete(h.availabilities, id)
	h.availabilitiesLock.Unlock()
}

func (h *HealthChecker) runHealthChecker(shutdown chan struct{}) {
	for {
		select {
		case <-time.After(healthCheckInterval):
			h.availabilitiesLock.RLock()

			h.logger.Debug("running health checks", zap.Int("availability count", len(h.availabilities)))

			for _, av := range h.availabilities {
				go h.checkInwayStatus(*av)
			}
			h.availabilitiesLock.RUnlock()

			h.logger.Debug("running health checks complete")
		case <-shutdown:
			return
		}
	}
}

func (h *HealthChecker) checkInwayStatus(av availability) {
	logger := h.logger.With(zap.Uint64("id", av.ID), zap.Uint64("inway-id", av.InwayID), zap.String("organization-serial-number", av.OrganizationSerialNumber), zap.String("service", av.ServiceName), zap.String("inway-address", av.Address))
	healthCheckURL := av.healthCheckURL()

	logger.Debug("checking inway status", zap.String("health check URL", healthCheckURL))

	resp, err := h.httpClient.Get(healthCheckURL)
	if err != nil {
		logger.Error("failed to check health", zap.Error(err))
		h.updateAvailabilityHealth(av, false)

		return
	}
	defer resp.Body.Close()

	logger.Debug("done checking inway status", zap.Int("http status code", resp.StatusCode))

	if resp.StatusCode != http.StatusOK {
		logger.Info(fmt.Sprintf("inway /health endpoint returned non-200 http status: %d", resp.StatusCode))
		h.updateAvailabilityHealth(av, false)

		return
	}

	status := &health.Status{}

	err = json.NewDecoder(resp.Body).Decode(&status)
	if err != nil {
		logger.Info("failed to parse json returned by the inway", zap.Error(err))
		h.updateAvailabilityHealth(av, false)

		return
	}

	logger.Debug("updating availability health")
	h.updateAvailabilityHealth(av, status.Healthy)
	logger.Debug("updating availability health done")

	logger.Debug("finished checking inway status", zap.Bool("health", status.Healthy), zap.String("version", status.Version))
}

func (h *HealthChecker) updateAvailabilityHealth(av availability, newHealth bool) {
	res, err := h.stmtUpdateHealth.Exec(av.ID, newHealth)
	if err != nil {
		h.logger.Error("failed to update health in db", zap.Error(err))
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		h.logger.Error("failed to get rows affected update health in db", zap.Error(err))
		return
	}

	if rowsAffected == 1 {
		if !newHealth {
			h.logger.Info(fmt.Sprintf("inway %s.%s>%s became unhealthy", av.OrganizationSerialNumber, av.ServiceName, av.Address))
		} else {
			h.logger.Info(fmt.Sprintf("inway %s.%s>%s became healthy", av.OrganizationSerialNumber, av.ServiceName, av.Address))
		}
	}
}
