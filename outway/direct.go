package outway

import (
	"encoding/binary"
	"fmt"
	"hash/crc64"
	"net/http"
	"strconv"
	"strings"

	"go.nlx.io/nlx/common/transactionlog"

	"go.uber.org/zap"
)

type destination struct {
	Organization string
	Service      string
	Path         string
}

func parseURLPath(urlPath string) (*destination, error) {
	pathParts := strings.SplitN(strings.TrimPrefix(urlPath, "/"), "/", 3)
	if len(pathParts) != 3 {
		return nil, fmt.Errorf("invalid path in url")
	}

	return &destination{
		Organization: pathParts[0],
		Service:      pathParts[1],
		Path:         pathParts[2],
	}, nil
}

func (o *Outway) handleDirect(w http.ResponseWriter, r *http.Request) {
	logger := o.logger.With(
		zap.String("connection-method", "direct"),
		zap.String("request-path", r.URL.Path),
		zap.String("request-remote-address", r.RemoteAddr),
	)

	destination, err := parseURLPath(r.URL.Path)
	if err != nil {
		logger.Error("error parsing URL", zap.Error(err))
		http.Error(w, "nlx outway: invalid path in url", http.StatusBadRequest)
		return
	}

	// Authorize request with plugged authorization service if authorization settings are set.
	if o.authorizationSettings != nil {
		authResponse, err := o.authorizeRequest(r.Header, destination)
		if err != nil {
			logger.Error("error authorizing request", zap.Error(err))
			http.Error(w, "nlx outway: error authorizing request", http.StatusInternalServerError)
			return
		}

		o.logger.Info("authorization result", zap.Bool("authorized", authResponse.Authorized), zap.String("reason", authResponse.Reason))
		if !authResponse.Authorized {
			http.Error(w, fmt.Sprintf("nlx outway: authorization failed. reason: %s", authResponse.Reason), http.StatusUnauthorized)
			return
		}
	}

	r.URL.Path = destination.Path

	recordData := createRecordData(r.Header, destination.Path)
	service := o.getService(destination.Organization, destination.Service)
	if service == nil {
		http.Error(w, "nlx outway: unknown service", http.StatusBadRequest)
		logger.Warn("received request for unknown service")
		return
	}

	logrecordIDFlake, err := o.requestFlake.NextID()
	if err != nil {
		logger.Error("could not get new request ID", zap.Error(err))
		http.Error(w, "nlx outway: internal server error", http.StatusInternalServerError)
		return
	}

	logrecordIDFlakeBytes := make([]byte, binary.MaxVarintLen64)
	binary.PutUvarint(logrecordIDFlakeBytes, logrecordIDFlake)
	logrecordIDNum := crc64.Checksum(logrecordIDFlakeBytes, o.ecmaTable)
	logrecordID := strconv.FormatUint(logrecordIDNum, 32)
	r.Header.Set("X-NLX-Logrecord-Id", logrecordID)

	dataSubjects, err := parseDataSubjects(r)
	if err != nil {
		http.Error(w, "nlx outway: invalid data subject header", http.StatusBadRequest)
		o.logger.Warn("invalid data subject header", zap.Error(err))
		return
	}

	o.stripHeaders(r, destination.Organization)

	err = o.txlogger.AddRecord(&transactionlog.Record{
		SrcOrganization:  o.organizationName,
		DestOrganization: destination.Organization,
		ServiceName:      destination.Service,
		LogrecordID:      logrecordID,
		Data:             recordData,
		DataSubjects:     dataSubjects,
	})
	if err != nil {
		http.Error(w, "nlx outway: server error", http.StatusInternalServerError)
		o.logger.Error("failed to store transactionlog record", zap.Error(err))
		return
	}

	o.logger.Info("forwarding API request", zap.String("destination-organization", destination.Organization), zap.String("service", destination.Service), zap.String("logrecord-id", logrecordID))

	service.ProxyHTTPRequest(w, r)
}

func createRecordData(h http.Header, p string) map[string]interface{} {
	recordData := make(map[string]interface{})
	recordData["request-path"] = p
	if processID := h.Get("X-NLX-Request-Process-Id"); processID != "" {
		recordData["doelbinding-process-id"] = processID
	}
	if dataElements := h.Get("X-NLX-Request-Data-Elements"); dataElements != "" {
		recordData["doelbinding-data-elements"] = dataElements
	}

	if userData := h.Get("X-NLX-Requester-User"); userData != "" {
		recordData["doelbinding-user"] = userData
	}

	if claims := h.Get("X-NLX-Requester-Claims"); claims != "" {
		recordData["doelbinding-claims"] = claims
	}

	if userID := h.Get("X-NLX-Request-User-Id"); userID != "" {
		recordData["doelbinding-user-id"] = userID
	}

	if applicationID := h.Get("X-NLX-Request-Application-Id"); applicationID != "" {
		recordData["doelbinding-application-id"] = applicationID
	}

	if subjectIdentifier := h.Get("X-NLX-Request-Subject-Identifier"); subjectIdentifier != "" {
		recordData["doelbinding-subject-identifier"] = subjectIdentifier
	}

	return recordData
}

func parseDataSubjects(r *http.Request) (map[string]string, error) {
	if dataSubjectsHeader := r.Header.Get("X-NLX-Request-Data-Subject"); dataSubjectsHeader != "" {
		return transactionlog.ParseDataSubjectHeader(dataSubjectsHeader)
	}

	return map[string]string{}, nil
}
