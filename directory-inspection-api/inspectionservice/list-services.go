// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

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
			o.name AS organization_name,
			s.name AS service_name,
			s.internal as service_internal,
			array_remove(array_agg(i.address), NULL) AS inway_addresses,
			COALESCE(s.documentation_url, '') AS documentation_url,
			COALESCE(s.api_specification_type, '') AS api_specification_type,
			COALESCE(s.public_support_contact, '') AS public_support_contact,
			array_remove(array_agg(a.healthy), NULL) as healthy_statuses
		FROM directory.services s
			INNER JOIN directory.organizations o
				ON s.organization_id = o.id
			LEFT JOIN directory.availabilities a
				ON s.id = a.service_id
			LEFT JOIN directory.inways i
				ON a.inway_id = i.id
		WHERE
			(internal = false OR (internal = true AND o.name = $1))
		GROUP BY s.id, o.id
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectServices")
	}

	return h, nil
}

func (h *listServicesHandler) ListServices(
	ctx context.Context,
	req *inspectionapi.ListServicesRequest,
) (*inspectionapi.ListServicesResponse, error) {
	h.logger.Info("rpc request ListServices()")
	resp := &inspectionapi.ListServicesResponse{}
	organizationName, err := getOrganisationNameFromRequest(ctx)
	if err != nil {
		return nil, err
	}
	h.logger.Debug("querying services", zap.String("organizationName", organizationName))
	rows, err := h.stmtSelectServices.Queryx(organizationName)
	if err != nil {
		h.logger.Error("failed to execute stmtSelectServices", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}
	for rows.Next() {
		var respService = &inspectionapi.ListServicesResponse_Service{}
		var inwayAddresses = pq.StringArray{}
		var healthyStatuses = pq.BoolArray{}
		err = rows.Scan(
			&respService.OrganizationName,
			&respService.ServiceName,
			&respService.Internal,
			&inwayAddresses,
			&respService.DocumentationUrl,
			&respService.ApiSpecificationType,
			&respService.PublicSupportContact,
			&healthyStatuses,
		)
		if err != nil {
			h.logger.Error("failed to scan into struct", zap.Error(err))
			return nil, status.New(codes.Internal, "Database error.").Err()
		}

		if len(inwayAddresses) != len(healthyStatuses) {
			h.logger.Error("length inwayadresses do not match healthchecks")
		} else {
			respService.InwayAddresses = inwayAddresses
			respService.HealthyStates = healthyStatuses
		}
		resp.Services = append(resp.Services, respService)
	}

	return resp, nil
}
