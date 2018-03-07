package inway

import (
	"encoding/json"
	"net/http"

	"github.com/VNG-Realisatie/nlx/monitor/health"
	"go.uber.org/zap"
)

func (i *Inway) handleHealthRequest(w http.ResponseWriter, r *http.Request, serviceName string) {
	i.serviceEndpointsLock.RLock()
	defer i.serviceEndpointsLock.RUnlock()

	// We currently only verify that the service still exists in this inway.
	// There is no health check to the actual endpoint defined yet.
	status := health.Status{}
	_, status.Healthy = i.serviceEndpoints[serviceName]

	err := json.NewEncoder(w).Encode(status)
	if err != nil {
		i.logger.Error("failed to encode health status json", zap.Error(err))
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	return
}
