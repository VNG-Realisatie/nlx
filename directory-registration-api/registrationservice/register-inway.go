package registrationservice

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

type registerInwayHandler struct {
	logger *zap.Logger

	stmtInsertAvailability *sqlx.Stmt

	httpClient *http.Client

	regexpName *regexp.Regexp
}

func NewRegisterInwayHandler(db *sqlx.DB, logger *zap.Logger, rootCA *x509.CertPool, certKeyPair tls.Certificate) (*registerInwayHandler, error) {
	h := &registerInwayHandler{
		logger: logger.With(zap.String("handler", "register-inway")),
	}

	var err error

	h.httpClient = newHTTPClient(rootCA, certKeyPair)

	h.regexpName, err = regexp.Compile(`^[a-zA-Z0-9-]{1,100}$`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile regexpName")
	}

	// NOTE: We do not have an endpoint yet to create services separately, therefore insert on demand.
	h.stmtInsertAvailability, err = db.Preparex(`
		WITH org AS (
			INSERT INTO directory.organizations (name, insight_log_endpoint, insight_irma_endpoint)
				VALUES ($1, $7, $8)
				ON CONFLICT ON CONSTRAINT organizations_uq_name
					DO UPDATE SET
						insight_log_endpoint = EXCLUDED.insight_log_endpoint,
						insight_irma_endpoint = EXCLUDED.insight_irma_endpoint
				RETURNING id
		), service AS (
			INSERT INTO directory.services (organization_id, name, internal, documentation_url, api_specification_type)
				SELECT org.id, $2, $3, NULLIF($4, ''), NULLIF($5, '')
					FROM org
				ON CONFLICT ON CONSTRAINT services_uq_name
					DO UPDATE SET documentation_url = EXCLUDED.documentation_url -- (possibly) no-op update to return id
				RETURNING id
		), inway AS (
			INSERT INTO directory.inways (organization_id, address)
				SELECT org.id, $6
					FROM org
				ON CONFLICT ON CONSTRAINT inways_uq_address
					DO UPDATE SET address = EXCLUDED.address -- no-op update to return id
				RETURNING id
		)
		INSERT INTO directory.availabilities (inway_id, service_id)
			SELECT inway.id, service.id
				FROM inway, service
			ON CONFLICT ON CONSTRAINT availabilities_uq_inway_service DO NOTHING
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtAssertService")
	}

	return h, nil
}

func (h *registerInwayHandler) RegisterInway(ctx context.Context, req *registrationapi.RegisterInwayRequest) (*registrationapi.RegisterInwayResponse, error) {
	h.logger.Info("rpc request RegisterInway", zap.String("inway address", req.InwayAddress))
	resp := &registrationapi.RegisterInwayResponse{}
	organizationName, err := getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, err
	}

	if !h.regexpName.MatchString(organizationName) {
		h.logger.Info("invalid organisation name in registerinwayrequest", zap.String("organization name", organizationName))
		return nil, status.New(codes.InvalidArgument, "Invalid organisation name").Err()
	}

	for _, service := range req.Services {
		if !h.regexpName.MatchString(service.Name) {
			h.logger.Info("invalid service name in registerinwayrequest", zap.String("service name", service.Name))
			return nil, status.New(codes.InvalidArgument, "Invalid servicename").Err()
		}
		h.logger.Info("service documentation url", zap.String("documentation url", service.ApiSpecificationDocumentUrl))
		var inwayAPISpecificationType string
		if len(service.ApiSpecificationDocumentUrl) > 0 {
			inwayAPISpecificationType, err = getInwayAPISpecsType(h.httpClient, req.InwayAddress, service.Name)
			if err != nil {
				h.logger.Info("invalid documentation url provided by inway", zap.String("documentation url", service.ApiSpecificationDocumentUrl), zap.Error(err))
				return nil, status.New(codes.InvalidArgument, "Invalid documentation URL provided").Err()
			}

			h.logger.Info("detected api spec", zap.String("apispectype", inwayAPISpecificationType))
		}

		_, err = h.stmtInsertAvailability.Exec(
			organizationName,
			service.Name,
			service.Internal,
			service.DocumentationUrl,
			inwayAPISpecificationType,
			req.InwayAddress,
			service.InsightApiUrl,
			service.IrmaApiUrl,
		)

		if err != nil {
			userFriendlyErrorText := "Database error."
			statusCode := codes.Internal
			pqErr, ok := err.(*pq.Error)
			if ok {
				switch pqErr.Constraint {
				case "services_check_typespec":
					userFriendlyErrorText = fmt.Sprintf("invalid api-specification-type '%s' configured for service '%s'", service.ApiSpecificationType, service.Name)
					statusCode = codes.InvalidArgument
				}
			}
			h.logger.Error("failed to execute stmtInsertAvailability", zap.Error(err))
			return nil, status.New(statusCode, userFriendlyErrorText).Err()
		}
	}

	return resp, nil
}

func newHTTPClient(rootCA *x509.CertPool, certKeyPair tls.Certificate) *http.Client {
	transport := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig: &tls.Config{
			RootCAs:      rootCA,
			Certificates: []tls.Certificate{certKeyPair},
		},
	}
	return &http.Client{
		Transport: transport,
	}
}

func getOrganisationNameFromRequest(ctx context.Context) (string, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return "", errors.New("failed to obtain peer from context")
	}
	tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
	return tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0], nil
}
