package monitor

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"go.nlx.io/nlx/common/process"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.nlx.io/nlx/directory-monitor/health"
	"go.uber.org/zap"
)

// HealthChecker checks the inways of a StoredService and modifies it's health state directly in the StoredService struct.
type healthChecker struct {
	wg         *sync.WaitGroup
	logger     *zap.Logger
	httpClient *http.Client

	availabilitiesLock sync.RWMutex
	availabilities     []availability

	stmtSelectAvailabilities *sqlx.Stmt
	stmtUpdateHealth         *sqlx.Stmt
}

type availability struct {
	ID               uint64
	OrganizationName string
	ServiceName      string
	Address          string
}

// RunHealthChecker starts a healthchecker process
func RunHealthChecker(process *process.Process, logger *zap.Logger, db *sqlx.DB, caCertPool *x509.CertPool, certKeyPair tls.Certificate) error {
	h := &healthChecker{
		logger: logger,
		httpClient: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certKeyPair},
			},
		}},
	}

	var err error
	h.stmtSelectAvailabilities, err = db.Preparex(`
		SELECT
			availabilities.id,
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
	h.stmtUpdateHealth, err = db.Preparex(`
		UPDATE directory.availabilities
			SET healthy = $2
			WHERE id = $1 AND healthy != $2
		`)
	if err != nil {
		return errors.Wrap(err, "failed to prepare stmtUpdateHealth")
	}

	return h.run(process)
}

func (h *healthChecker) run(process *process.Process) error {
	chInitialLoad := make(chan struct{})
	go func() {
		for {
			newAvailabilities := []availability{}
			err := h.stmtSelectAvailabilities.Select(&newAvailabilities)
			if err != nil {
				h.logger.Error("failed to fetch availabilities", zap.Error(err))
				time.Sleep(1 * time.Second)
				continue
			}

			// replace availabilities slice in healthChecker
			h.availabilitiesLock.Lock()
			h.availabilities = newAvailabilities
			h.availabilitiesLock.Unlock()

			// signal initial load has completed
			select {
			case chInitialLoad <- struct{}{}:
			default:
			}

			// TODO: #207 use NOTIFY structure on postgres instead of reloading
			// refresh this list every 10 seconds
			time.Sleep(10 * time.Second)
		}
	}()

	<-chInitialLoad // wait for initial load to be done
	shutDown := make(chan struct{})
	process.CloseGracefully(func() error {
		close(shutDown)
		return nil
	})
	for {
		select {
		case <-time.After(5 * time.Second):
			h.availabilitiesLock.RLock()
			for _, av := range h.availabilities {
				go h.checkAvailabilityHealth(av)
			}
			h.availabilitiesLock.RUnlock()
		case <-shutDown:
			return nil
		}
	}
}

func (h *healthChecker) checkAvailabilityHealth(av availability) {
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
}

func (h *healthChecker) updateAvailabilityHealth(av availability, newHealth bool) {
	res, err := h.stmtUpdateHealth.Exec(av.ID, newHealth)
	if err != nil {
		h.logger.Error("failed to update health in db", zap.Error(err))
		return
	}
	affected, err := res.RowsAffected()
	if err != nil {
		h.logger.Error("failed to get affected rows after updating health in db", zap.Error(err))
		return
	}

	if affected == 1 {
		// log health change
		if !newHealth {
			h.logger.Info(fmt.Sprintf("inway %s.%s>%s became unhealthy", av.OrganizationName, av.ServiceName, av.Address))
		}
		if newHealth {
			h.logger.Info(fmt.Sprintf("inway %s.%s>%s became healthy", av.OrganizationName, av.ServiceName, av.Address))
		}
	}
}
