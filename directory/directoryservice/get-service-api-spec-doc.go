package directoryservice

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/directory/directoryapi"
)

type getServiceAPISpecHandler struct {
	logger *zap.Logger

	httpClient *http.Client

	stmtSelectServiceInway *sqlx.Stmt
}

func newGetServiceAPISpecHandler(db *sqlx.DB, logger *zap.Logger, rootCA *x509.CertPool, certKeyPair tls.Certificate) (*getServiceAPISpecHandler, error) {
	h := &getServiceAPISpecHandler{
		logger: logger.With(zap.String("handler", "list-services")),
	}

	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		// This can happen when the internals of net/http change.
		// Afaik an interface implementation isn't under the Go1 compatibility promise.
		// TODO: #209 consider setting up a custom http.Transport to use as the proxies RoundTripper.
		panic("http.DefaultTransport must be of type *http.Transport")
	}
	// load client certificate
	transport.TLSClientConfig = &tls.Config{
		RootCAs:      rootCA,
		Certificates: []tls.Certificate{certKeyPair},
	}
	h.httpClient = &http.Client{
		Transport: transport,
	}

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

func (h *getServiceAPISpecHandler) GetServiceAPISpec(ctx context.Context, req *directoryapi.GetServiceAPISpecRequest) (*directoryapi.GetServiceAPISpecResponse, error) {
	h.logger.Info("rpc request GetServiceAPISpec()")
	resp := &directoryapi.GetServiceAPISpecResponse{}

	var inwayAddress string
	err := h.stmtSelectServiceInway.QueryRowx(req.OrganizationName, req.ServiceName).Scan(&inwayAddress, &resp.Type)
	if err != nil {
		h.logger.Error("failed to execute stmtSelectServiceInway", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	inwayURL := url.URL{
		Scheme: "https",
		Host:   inwayAddress,
		Path:   path.Join("/.nlx/api-spec-doc/", req.ServiceName),
	}

	res, err := h.httpClient.Get(inwayURL.String())
	if err != nil {
		h.logger.Info("failed to fetch api spec doc from remote inway", zap.String("inwayURL", inwayURL.String()), zap.Error(err))
		return nil, status.New(codes.InvalidArgument, "Invalid inway URL").Err()
	}
	defer res.Body.Close()

	resp.Document, err = ioutil.ReadAll(res.Body)
	if err != nil {
		h.logger.Info("failed to read api spec doc from remote inway", zap.String("inwayURL", inwayURL.String()), zap.Error(err))
		return nil, status.New(codes.InvalidArgument, "Invalid inway URL").Err()
	}

	return resp, nil
}
