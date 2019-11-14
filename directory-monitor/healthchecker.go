// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package monitor

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"go.nlx.io/nlx/common/process"
	"go.nlx.io/nlx/directory-monitor/health"
)

// HealthChecker checks the inways of a StoredService and modifies it's health state directly in the StoredService struct.
type healthChecker struct {
	logger     *zap.Logger
	httpClient *http.Client

	availabilitiesLock sync.RWMutex
	availabilities     map[uint64]*availability

	stmtSelectAvailabilities  *sqlx.Stmt
	stmtUpdateHealth          *sqlx.Stmt
	stmtCleanUpAvailabilities *sqlx.Stmt
	stmtUpdateInwayVersion    *sqlx.Stmt
}

type availability struct {
	ID               uint64 `json:"id"`
	OrganizationName string `json:"organization_name"`
	ServiceName      string `json:"service_name"`
	Address          string `json:"address"`
	InwayID          uint64 `json:"inway_id"`
}

type databaseAction struct {
	Action       string        `json:"action"`
	Availability *availability `json:"availability"`
}

var RunningHealthChecker *healthChecker

const dbNotificationChannel = "availabilities"
const cleanupInterval = 90 * time.Second

// RunHealthChecker starts a healthchecker process
func RunHealthChecker(
	proc *process.Process,
	logger *zap.Logger,
	db *sqlx.DB,
	postgresDNS string,
	caCertPool *x509.CertPool,
	certKeyPair *tls.Certificate,
	ttlOfflineService int,
) error {
	h := &healthChecker{
		logger:         logger,
		availabilities: make(map[uint64]*availability),
		httpClient: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{*certKeyPair},
			},
		}},
	}

	var err error
	h.stmtSelectAvailabilities, err = db.Preparex(`
		SELECT
			availabilities.id,
			availabilities.inway_id,
			organizations.name AS organization_name,
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

	return h.run(proc, postgresDNS)
}

func (h *healthChecker) run(proc *process.Process, postgressDNS string) error {
	listenerErrorCallback := func(ev pq.ListenerEventType, err error) {
		if err != nil {
			h.logger.Error("error listening for db events", zap.Error(err))
		}
	}
	listener := pq.NewListener(postgressDNS, 10*time.Second, time.Minute, listenerErrorCallback)
	err := listener.Listen(dbNotificationChannel)
	if err != nil {
		panic(err)
	}
	proc.CloseGracefully(listener.Close)

	RunningHealthChecker = h

	shutDownNotificationListener := make(chan struct{})
	proc.CloseGracefully(func() error {
		close(shutDownNotificationListener)
		return nil
	})
	go h.waitForNotification(listener, shutDownNotificationListener)

	newAvailabilities := []*availability{}
	err = h.stmtSelectAvailabilities.Select(&newAvailabilities)
	if err != nil {
		return err
	}

	h.availabilitiesLock.Lock()
	for _, availability := range newAvailabilities {
		h.availabilities[availability.ID] = availability
	}
	h.availabilitiesLock.Unlock()

	shutDown := make(chan struct{})
	proc.CloseGracefully(func() error {
		close(shutDown)
		return nil
	})

	go h.runCleanUpServices(shutDown)
	h.runHealthChecker(shutDown)
	return nil
}

func (h *healthChecker) runCleanUpServices(shutDown chan struct{}) {
	h.logger.Debug("initial cleaning up stale services")
	servicesRemoved, err := h.cleanUpServices()

	if err != nil {
		h.logger.Error("error cleaning up offline services", zap.Error(err))
	}

	h.logger.Debug("cleanup complete", zap.Int64("services removed", servicesRemoved))

	for {
		select {
		case <-time.After(1 * time.Minute):
			h.logger.Debug("cleaning up offline services")
			servicesRemoved, err := h.cleanUpServices()

			if err != nil {
				h.logger.Error("error cleaning up offline services", zap.Error(err))
				continue
			}

			h.logger.Debug("cleanup complete", zap.Int64("services removed", servicesRemoved))
		case <-shutDown:
			return
		}
	}
}

func (h *healthChecker) cleanUpServices() (int64, error) {
	res, err := h.stmtCleanUpAvailabilities.Exec()
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func (h *healthChecker) waitForNotification(l *pq.Listener, c <-chan struct{}) {
	for {
		select {
		case n := <-l.Notify:
			if n != nil {
				h.onDatabaseNotification(n.Extra)
			}
		case <-time.After(cleanupInterval):
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

func (h *healthChecker) onDatabaseNotification(payload string) {
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

func (h *healthChecker) addAvailability(a *availability) {
	h.availabilitiesLock.Lock()
	h.availabilities[a.ID] = a
	h.availabilitiesLock.Unlock()
}

func (h *healthChecker) removeAvailability(id uint64) {
	h.availabilitiesLock.Lock()
	delete(h.availabilities, id)
	h.availabilitiesLock.Unlock()
}

func (h *healthChecker) runHealthChecker(shutDown chan struct{}) {
	for {
		select {
		case <-time.After(5 * time.Second):
			h.availabilitiesLock.RLock()
			for _, av := range h.availabilities {
				go h.checkInwayStatus(*av)
			}
			h.availabilitiesLock.RUnlock()
		case <-shutDown:
			return
		}
	}
}

func (h *healthChecker) checkInwayStatus(av availability) {
	logger := h.logger.With(zap.String("canonical-service-name", av.OrganizationName+`.`+av.ServiceName), zap.String("inway-address", av.Address))

	resp, err := h.httpClient.Get(`https://` + av.Address + "/.nlx/health/" + av.ServiceName)
	if err != nil {
		logger.Error("failed to check health", zap.Error(err))
		h.updateAvailabilityHealth(av, false)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
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

	h.updateAvailabilityHealth(av, status.Healthy)
	h.updateInwayVersion(av, status.Version)
}

func (h *healthChecker) updateAvailabilityHealth(av availability, newHealth bool) {
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
			h.logger.Info(fmt.Sprintf("inway %s.%s>%s became unhealthy", av.OrganizationName, av.ServiceName, av.Address))
		} else {
			h.logger.Info(fmt.Sprintf("inway %s.%s>%s became healthy", av.OrganizationName, av.ServiceName, av.Address))
		}
	}
}

func (h *healthChecker) updateInwayVersion(av availability, version string) {
	if version == "" {
		h.logger.Info("no inway version recieved")
		return
	}
	res, err := h.stmtUpdateInwayVersion.Exec(av.InwayID, version)
	if err != nil {
		h.logger.Error("failed to update inway version in db", zap.Error(err))
		return
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		h.logger.Error("failed to get rows affected update inwayversion in db", zap.Error(err))
		return
	}
	if rowsAffected == 1 {
		h.logger.Info(fmt.Sprintf("inway %s.%s>%s version %s",
			av.OrganizationName, av.ServiceName, av.Address, version))
	}
}
