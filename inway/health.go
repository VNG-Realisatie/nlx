package inway

import (
	"encoding/json"
	"net/http"
	"strings"

	"go.uber.org/zap"

	"go.nlx.io/nlx/directory-monitor/health"
)

func (i *Inway) handleHealthRequest(w http.ResponseWriter, r *http.Request) {
	i.serviceEndpointsLock.RLock()
	defer i.serviceEndpointsLock.RUnlock()

	serviceName := strings.TrimPrefix(r.URL.Path, "/.nlx/health/")

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
