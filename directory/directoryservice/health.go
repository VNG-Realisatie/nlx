package directoryservice

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// StoreHealthChecker checks the inways of a StoredService and modifies it's health state directly in the StoredService struct.
type StoreHealthChecker struct {
	logger     *zap.Logger
	store      *Store
	httpClient *http.Client
}

// StartStoreHealthChecker creates and starts a StoreHealthChecker.
// TODO: once we move to actual persistent storage:
// - make StoreHealthChecker unaware of the Store and directyly talk to the database
// - move the healthchecks into a completely seperate process nlx-dirhealth (?)
func StartStoreHealthChecker(logger *zap.Logger, caCertPool *x509.CertPool, certKeyPair tls.Certificate, store *Store) *StoreHealthChecker {
	h := &StoreHealthChecker{
		logger: logger,
		store:  store,
		httpClient: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{certKeyPair},
			},
		}},
	}
	go h.run()
	return h
}

func (h *StoreHealthChecker) run() {
	for {
		h.store.ServicesLock.RLock()
		for _, service := range h.store.Services {
			service.InwayAddressesLock.RLock()
			for address := range service.InwayAddresses {
				go h.checkInwayHealth(service, address)
			}
			service.InwayAddressesLock.RUnlock()
		}
		h.store.ServicesLock.RUnlock()
		time.Sleep(5 * time.Second)
	}
}

func (h *StoreHealthChecker) checkInwayHealth(service *StoredService, address string) {
	logger := h.logger.With(zap.String("canonical-service-name", service.CanonicalServiceName()), zap.String("inway-address", address))

	resp, err := h.httpClient.Get(`https://` + address + "/health")
	if err != nil {
		logger.Error("failed to check health", zap.Error(err))
		h.updateInwayHealth(service, address, false)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		logger.Info(fmt.Sprintf("inway /health endpoint returned non-200 http status: %d", resp.StatusCode))
		h.updateInwayHealth(service, address, false)
		return
	}
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("failed to read response body from health check", zap.Error(err))
		h.updateInwayHealth(service, address, false)
		return
	}
	if !bytes.Equal(responseBody, []byte("ok")) {
		logger.Info("inway /health endpoint did not return `ok`")
		h.updateInwayHealth(service, address, false)
		return
	}
	h.updateInwayHealth(service, address, true)
}

func (h *StoreHealthChecker) updateInwayHealth(service *StoredService, address string, newHealth bool) {
	// check for change - fast path
	service.InwayAddressesLock.RLock()
	prevHealth := service.InwayAddresses[address]
	service.InwayAddressesLock.RUnlock()
	if prevHealth == newHealth {
		return // no change
	}

	// check for change - with full lock
	service.InwayAddressesLock.Lock()
	defer service.InwayAddressesLock.Unlock()
	prevHealth = service.InwayAddresses[address]
	if prevHealth == newHealth {
		return // no change anymore
	}
	service.InwayAddresses[address] = newHealth

	// log health change
	if !newHealth {
		h.logger.Info(fmt.Sprintf("inway %s>%s became unhealthy", service.CanonicalServiceName(), address))
	}
	if newHealth {
		h.logger.Info(fmt.Sprintf("inway %s>%s became healthy", service.CanonicalServiceName(), address))
	}
}
