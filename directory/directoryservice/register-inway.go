package directoryservice

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory/directoryapi"
)

type registerInwayHandler struct {
	logger *zap.Logger

	stmtInsertAvailability *sqlx.Stmt

	regexpName *regexp.Regexp
}

func newRegisterInwayHandler(db *sqlx.DB, logger *zap.Logger) (*registerInwayHandler, error) {
	h := &registerInwayHandler{
		logger: logger.With(zap.String("handler", "register-inway")),
	}

	var err error

	h.regexpName, err = regexp.Compile(`^[a-zA-Z0-9-]{1,100}$`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to compile regexpName")
	}

	// NOTE: We do not have an endpoint yet to create services separately, therefore insert on demand.
	h.stmtInsertAvailability, err = db.Preparex(`
		WITH org AS (
			INSERT INTO directory.organizations (name, insight_log_endpoint, insight_irma_endpoint)
				VALUES ($1, $6, $7)
				ON CONFLICT ON CONSTRAINT organizations_uq_name
					DO UPDATE SET
						insight_log_endpoint = EXCLUDED.insight_log_endpoint,
						insight_irma_endpoint = EXCLUDED.insight_irma_endpoint
				RETURNING id
		), service AS (
			INSERT INTO directory.services (organization_id, name, documentation_url, api_specification_type)
				SELECT org.id, $2, NULLIF($3, ''), NULLIF($4, '')
					FROM org
				ON CONFLICT ON CONSTRAINT services_uq_name
					DO UPDATE SET documentation_url = EXCLUDED.documentation_url -- (possibly) no-op update to return id
				RETURNING id
		), inway AS (
			INSERT INTO directory.inways (organization_id, address)
				SELECT org.id, $5
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

func (h *registerInwayHandler) RegisterInway(ctx context.Context, req *directoryapi.RegisterInwayRequest) (*directoryapi.RegisterInwayResponse, error) {
	h.logger.Info("rpc request RegisterInway", zap.String("inway address", req.InwayAddress))
	resp := &directoryapi.RegisterInwayResponse{}
	peer, ok := peer.FromContext(ctx)
	if !ok {

		return nil, errors.New("failed to obtain peer from context")
	}
	tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
	organizationName := tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0]
	// TODO: #206 when administrative (client-tls mandatory) and inspection (client-tls optional) endpoints have been separated,
	// use proper grpc authentication via middleware and context (based on client-tls fields (CN, O) like we do here)

	if !h.regexpName.MatchString(organizationName) {
		h.logger.Info("invalid organisation name in registerinwayrequest", zap.String("organization name", organizationName))
		return nil, status.New(codes.InvalidArgument, "Invalid organisation name").Err()
	}

	for _, service := range req.Services {
		if !h.regexpName.MatchString(service.Name) {
			h.logger.Info("invalid service name in registerinwayrequest", zap.String("service name", service.Name))
			return nil, status.New(codes.InvalidArgument, "Invalid servicename").Err()
		}
		_, err := url.Parse(service.DocumentationUrl)
		if err != nil {
			h.logger.Info("invalid documentation url provided by inway", zap.String("documentation url", service.DocumentationUrl), zap.Error(err))
			return nil, status.New(codes.InvalidArgument, "Invalid documentation URL provided").Err()
		}

		_, err = h.stmtInsertAvailability.Exec(
			organizationName,
			service.Name,
			service.DocumentationUrl,
			service.ApiSpecificationType,
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
