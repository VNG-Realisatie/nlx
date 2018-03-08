package directoryservice

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

type registerInwayHandler struct {
	stmtInsertAvailability *sqlx.Stmt
}

func newRegisterInwayHandler(db *sqlx.DB, logger *zap.Logger) (*registerInwayHandler, error) {
	h := &registerInwayHandler{}

	var err error
	// NOTE: We do not have an endpoint yet to create services seperately, therefore insert on demand.
	h.stmtInsertAvailability, err = db.Preparex(`
		WITH org AS (
			INSERT INTO directory.organizations (name)
				VALUES ($1)
				ON CONFLICT ON CONSTRAINT organizations_uq_name
					DO UPDATE SET name = EXCLUDED.name
				RETURNING id
		), service AS (
			INSERT INTO directory.services (organization_id, name, documentation_url)
				SELECT org.id, $2, NULLIF($3, '')
					FROM org
				ON CONFLICT ON CONSTRAINT services_uq_name
					DO UPDATE SET documentation_url = EXCLUDED.documentation_url -- (possibly) no-op update to return id
				RETURNING id
		), inway AS (
			INSERT INTO directory.inways (organization_id, address)
				SELECT org.id, $4
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
	fmt.Printf("rpc request RegisterInway(%s)\n", req.InwayAddress)
	resp := &directoryapi.RegisterInwayResponse{}

	peer, ok := peer.FromContext(ctx)
	if !ok {
		return nil, errors.New("failed to obtain peer from context")
	}
	tlsInfo := peer.AuthInfo.(credentials.TLSInfo)
	organizationName := tlsInfo.State.VerifiedChains[0][0].Subject.Organization[0]
	// TODO: when administrative (client-tls mandatory) and inspection (client-tls optional) endpoints have been seperated,
	// use proper grpc authentication via middleware and context (based on client-tls fields (CN, O) like we do here)

	for _, service := range req.Services {
		_, err := h.stmtInsertAvailability.Exec(organizationName, service.Name, service.DocumentationUrl, req.InwayAddress)
		if err != nil {
			return nil, errors.Wrap(err, "failed to insert service and inway")
		}
	}

	return resp, nil
}
