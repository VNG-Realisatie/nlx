package inspectionservice

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

type listServicesHandler struct {
	logger *zap.Logger

	stmtSelectServices *sqlx.Stmt
}

func newListServicesHandler(db *sqlx.DB, logger *zap.Logger) (*listServicesHandler, error) {
	h := &listServicesHandler{
		logger: logger.With(zap.String("handler", "list-services")),
	}

	var err error
	h.stmtSelectServices, err = db.Preparex(`
		SELECT
			organizations.name AS organization_name,
			services.name AS service_name,
			services.internal as service_internal,
			array_remove(array_agg(inways.address), NULL) AS inway_addresses,
			COALESCE(services.documentation_url, '') AS documentation_url,
			COALESCE(services.api_specification_type, '') AS api_specification_type
		FROM directory.services
			INNER JOIN directory.organizations
				ON services.organization_id = organizations.id
			LEFT JOIN directory.availabilities
				ON services.id = availabilities.service_id AND availabilities.healthy = true
			LEFT JOIN directory.inways
				ON availabilities.inway_id = inways.id
		WHERE
			internal = false OR (internal = true AND organizations.name = $1)
		GROUP BY services.id, organizations.id
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectServices")
	}

	return h, nil
}

func (h *listServicesHandler) ListServices(ctx context.Context, req *inspectionapi.ListServicesRequest) (*inspectionapi.ListServicesResponse, error) {
	h.logger.Info("rpc request ListServices()")
	resp := &inspectionapi.ListServicesResponse{}
	organizationName, err := getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, err
	}
	rows, err := h.stmtSelectServices.Queryx(organizationName)
	if err != nil {
		h.logger.Error("failed to execute stmtSelectServices", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}
	for rows.Next() {
		var respService = &inspectionapi.ListServicesResponse_Service{}
		var inwayAddresses = pq.StringArray{}
		err = rows.Scan(
			&respService.OrganizationName,
			&respService.ServiceName,
			&respService.Internal,
			&inwayAddresses,
			&respService.DocumentationUrl,
			&respService.ApiSpecificationType,
		)
		if err != nil {
			h.logger.Error("failed to scan into struct", zap.Error(err))
			return nil, status.New(codes.Internal, "Database error.").Err()
		}
		respService.InwayAddresses = []string(inwayAddresses)
		resp.Services = append(resp.Services, respService)
	}

	return resp, nil
}
