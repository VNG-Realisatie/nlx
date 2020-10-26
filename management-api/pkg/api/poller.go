package api

import (
	"context"
	"sync"
	"time"

	"github.com/gogo/protobuf/types"
	"go.uber.org/zap"

	common_tls "go.nlx.io/nlx/common/tls"
	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/management-api/api"
	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/directory"
	"go.nlx.io/nlx/management-api/pkg/management"
)

type accessProofPoller struct {
	logger                     *zap.Logger
	orgCert                    *common_tls.CertificateBundle
	configDatabase             database.ConfigDatabase
	directoryClient            directory.Client
	createManagementClientFunc func(context.Context, string, *common_tls.CertificateBundle) (management.Client, error)
}

func newAccessProofPoller(logger *zap.Logger, directoryClient directory.Client, configDatabase database.ConfigDatabase, orgCert *common_tls.CertificateBundle) *accessProofPoller {
	return &accessProofPoller{
		logger:                     logger,
		directoryClient:            directoryClient,
		configDatabase:             configDatabase,
		orgCert:                    orgCert,
		createManagementClientFunc: management.NewClient,
	}
}

func (poller *accessProofPoller) Poll(ctx context.Context) {
	wg := &sync.WaitGroup{}
	ticker := time.NewTicker(pollInterval * time.Second)

	defer ticker.Stop()

pollingLoop:
	for {
		select {
		case <-ticker.C:
			wg.Add(1)

			go func() {
				defer wg.Done()

				if err := poller.syncAccessProofs(context.TODO()); err != nil {
					poller.logger.Error("failed to sync access proofs", zap.Error(err))
				}
			}()
		case <-ctx.Done():
			break pollingLoop
		}
	}

	wg.Wait()
}

func (poller *accessProofPoller) getOrganizationManagementClient(ctx context.Context, organizationName string) (management.Client, error) {
	response, err := poller.directoryClient.GetOrganizationInway(ctx, &inspectionapi.GetOrganizationInwayRequest{
		OrganizationName: organizationName,
	})
	if err != nil {
		return nil, err
	}

	address, err := computeInwayProxyAddress(response.Address)
	if err != nil {
		return nil, err
	}

	client, err := poller.createManagementClientFunc(ctx, address, poller.orgCert)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (poller *accessProofPoller) parseAccessProof(accessProof *api.AccessProof) (*database.AccessProof, error) {
	createdAt, err := types.TimestampFromProto(accessProof.CreatedAt)
	if err != nil {
		return nil, err
	}

	revokedAt, err := types.TimestampFromProto(accessProof.RevokedAt)
	if err != nil {
		revokedAt = time.Time{}
	}

	return &database.AccessProof{
		ID:               accessProof.Id,
		CreatedAt:        createdAt,
		RevokedAt:        revokedAt,
		OrganizationName: accessProof.OrganizationName,
		ServiceName:      accessProof.ServiceName,
	}, nil
}

func (poller *accessProofPoller) syncAccessProof(ctx context.Context, request *database.OutgoingAccessRequest) error {
	client, err := poller.getOrganizationManagementClient(ctx, request.OrganizationName)
	if err != nil {
		return err
	}

	response, err := client.GetAccessProof(ctx, &external.GetAccessProofRequest{
		ServiceName: request.ServiceName,
	})
	if err != nil {
		return err
	}

	remoteProof, err := poller.parseAccessProof(response)
	if err != nil {
		return err
	}

	localProof, err := poller.configDatabase.GetLatestAccessProofForService(ctx, request.OrganizationName, request.ServiceName)

	switch err {
	case nil:
	case database.ErrNotFound:
		_, err = poller.configDatabase.CreateAccessProof(ctx, &database.AccessProof{
			OrganizationName: remoteProof.OrganizationName,
			ServiceName:      remoteProof.ServiceName,
			CreatedAt:        remoteProof.CreatedAt,
			RevokedAt:        remoteProof.RevokedAt,
		})

		return err
	default:
		return err
	}

	if !remoteProof.RevokedAt.IsZero() &&
		localProof.RevokedAt != remoteProof.RevokedAt {
		if _, err := poller.configDatabase.RevokeAccessProof(
			ctx,
			localProof.OrganizationName,
			localProof.ServiceName,
			localProof.ID,
			remoteProof.RevokedAt,
		); err != nil {
			return err
		}
	}

	return nil
}

func (poller *accessProofPoller) syncAccessProofs(ctx context.Context) error {
	requests, err := poller.configDatabase.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return err
	}

	for _, request := range requests {
		// only sync approved access requests
		if request.State != database.AccessRequestApproved {
			continue
		}

		if err := poller.syncAccessProof(ctx, request); err != nil {
			poller.logger.Error("failed to sync access proof", zap.Error(err))
		}
	}

	return nil
}
