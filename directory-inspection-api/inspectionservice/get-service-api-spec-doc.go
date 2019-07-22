// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

package inspectionservice

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
)

type getServiceAPISpecHandler struct {
	logger *zap.Logger

	httpClient *http.Client

	stmtSelectServiceInway *sqlx.Stmt
}

func newGetServiceAPISpecHandler(
	db *sqlx.DB,
	logger *zap.Logger,
	rootCA *x509.CertPool,
	certKeyPair *tls.Certificate,
) (*getServiceAPISpecHandler, error) {
	h := &getServiceAPISpecHandler{
		logger: logger.With(zap.String("handler", "list-services")),
	}

	h.httpClient = newHTTPClient(rootCA, certKeyPair)

	var err error
	h.stmtSelectServiceInway, err = db.Preparex(`
		SELECT
			inways.address AS inway_address,
			services.api_specification_type
		FROM directory.services
			INNER JOIN directory.organizations
				ON services.organization_id = organizations.id
			INNER JOIN directory.availabilities
				ON services.id = availabilities.service_id AND availabilities.healthy = true
			INNER JOIN directory.inways
				ON availabilities.inway_id = inways.id
        WHERE organizations.name = $1 AND services.name = $2
        LIMIT 1
	`)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare stmtSelectServiceInway")
	}

	return h, nil
}

func (h *getServiceAPISpecHandler) GetServiceAPISpec(ctx context.Context, req *inspectionapi.GetServiceAPISpecRequest) (*inspectionapi.GetServiceAPISpecResponse, error) {
	h.logger.Info("rpc request GetServiceAPISpec()")
	resp := &inspectionapi.GetServiceAPISpecResponse{}

	var inwayAddress string
	err := h.stmtSelectServiceInway.QueryRowx(req.OrganizationName, req.ServiceName).Scan(&inwayAddress, &resp.Type)
	if err != nil {
		h.logger.Error("failed to execute stmtSelectServiceInway", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	resp.Document, err = getInwayAPISpecs(h.httpClient, inwayAddress, req.ServiceName)
	if err != nil {
		h.logger.Info("failed to read api spec doc from remote inway", zap.Error(err))
		return nil, status.New(codes.InvalidArgument, "Invalid inway URL").Err()
	}

	return resp, nil
}
