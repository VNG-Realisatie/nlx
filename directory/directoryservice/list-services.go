package directoryservice

import (
	"context"
	"fmt"

	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
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
			services.name AS name,
			COALESCE(services.documentation_url, '') AS documentation_url,
			array_remove(array_agg(inways.address), NULL) AS inway_addresses
		FROM directory.services
			INNER JOIN directory.organizations
				ON services.organization_id = organizations.id
			LEFT JOIN directory.availabilities
				ON services.id = availabilities.service_id AND availabilities.healthy = true
			LEFT JOIN directory.inways
				ON availabilities.inway_id = inways.id
		GROUP BY services.id, organizations.id
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectServices")
	}

	return h, nil
}

func (h *listServicesHandler) ListServices(ctx context.Context, req *directoryapi.ListServicesRequest) (*directoryapi.ListServicesResponse, error) {
	fmt.Println("rpc request ListServices()")
	resp := &directoryapi.ListServicesResponse{}

	rows, err := h.stmtSelectServices.Queryx()
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute stmtSelectServices")
	}
	for rows.Next() {
		var respService = &directoryapi.Service{}
		var inwayAddresses = pq.StringArray{}
		err = rows.Scan(
			&respService.OrganizationName,
			&respService.Name,
			&respService.DocumentationUrl,
			&inwayAddresses,
		)
		if err != nil {
			return nil, errors.Wrap(err, "failed to scan into struct")
		}
		respService.InwayAddresses = []string(inwayAddresses)
		resp.Services = append(resp.Services, respService)
	}

	return resp, nil
}
