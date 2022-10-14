// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package syncer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/management-api/api/external"
	"go.nlx.io/nlx/management-api/pkg/database"
	"go.nlx.io/nlx/management-api/pkg/management"
)

var ErrServiceDoesNotExist = status.Error(codes.NotFound, "service does not exist")

type SyncArgs struct {
	Ctx      context.Context
	Logger   *zap.Logger
	DB       database.ConfigDatabase
	Client   management.Client
	Requests []*database.OutgoingAccessRequest
}

func SyncOutgoingAccessRequests(args *SyncArgs) error {
	if len(args.Requests) < 1 {
		return nil
	}

	hasError := false
	waitGroup := sync.WaitGroup{}

	for _, request := range args.Requests {
		requestToSync := request

		waitGroup.Add(1)

		go func() {
			defer waitGroup.Done()

			args.Logger.Debug(
				"syncing access request",
				zap.String("organization", requestToSync.Organization.SerialNumber),
				zap.String("service", requestToSync.ServiceName),
				zap.String("public key fingerprint", requestToSync.PublicKeyFingerprint),
			)

			err := synchronizeOutgoingAccessRequest(context.Background(), args.DB, args.Client, requestToSync)
			if err != nil {
				args.Logger.Error("failed to synchronize outgoing access request", zap.Error(err))

				hasError = true
			}
		}()
	}

	waitGroup.Wait()

	if hasError {
		return errors.New("error occurred while syncing at least one Outgoing Access Request")
	} else {
		return nil
	}
}

//nolint:gocyclo // unable to reduce complexity
func synchronizeOutgoingAccessRequest(ctx context.Context, configDatabase database.ConfigDatabase, client management.Client, request *database.OutgoingAccessRequest) error {
	switch request.State {
	case database.OutgoingAccessRequestReceived:
		response, err := client.GetAccessRequestState(ctx, &external.GetAccessRequestStateRequest{
			ServiceName:          request.ServiceName,
			PublicKeyFingerprint: request.PublicKeyFingerprint,
		})
		if err != nil {
			if errors.Is(err, ErrServiceDoesNotExist) {
				return configDatabase.DeleteOutgoingAccessRequests(ctx, request.Organization.SerialNumber, request.ServiceName)
			}

			return fmt.Errorf("failed to get access request state from organization")
		}

		newState, err := convertAccessRequestState(response.State)
		if err != nil {
			return fmt.Errorf("unable to convert access request state: %v", err)
		}

		if newState == database.OutgoingAccessRequestApproved {
			err := syncAccessProof(ctx, configDatabase, client, request)
			if err != nil {
				return err
			}
		}

		return configDatabase.UpdateOutgoingAccessRequestState(
			ctx,
			request.ID,
			newState,
			uint(0), // '0', because we don't want to update this value
			nil,
		)

	case database.OutgoingAccessRequestApproved:
		err := syncAccessProof(ctx, configDatabase, client, request)
		if err != nil {
			if errors.Is(err, ErrServiceDoesNotExist) {
				return configDatabase.DeleteOutgoingAccessRequests(ctx, request.Organization.SerialNumber, request.ServiceName)
			}

			return fmt.Errorf("failed to sync access proof for an approved outgoing access request: %e", err)
		}

		return configDatabase.UpdateOutgoingAccessRequestState(
			ctx,
			request.ID,
			database.OutgoingAccessRequestApproved,
			uint(0), // '0', because we don't want to update this value
			nil,
		)

	case database.OutgoingAccessRequestFailed, database.OutgoingAccessRequestRejected:
		return nil

	default:
		return fmt.Errorf("invalid state '%s' for pending access request", request.State)
	}
}

func convertAccessRequestState(state external.AccessRequestState) (database.OutgoingAccessRequestState, error) {
	switch state {
	case external.AccessRequestState_ACCESS_REQUEST_STATE_APPROVED:
		return database.OutgoingAccessRequestApproved, nil
	case external.AccessRequestState_ACCESS_REQUEST_STATE_REJECTED:
		return database.OutgoingAccessRequestRejected, nil
	case external.AccessRequestState_ACCESS_REQUEST_STATE_RECEIVED:
		return database.OutgoingAccessRequestReceived, nil
	case external.AccessRequestState_ACCESS_REQUEST_STATE_FAILED:
		return database.OutgoingAccessRequestFailed, nil
		/*
			If the returned state is revoked the outgoing access request state needs to be set to approved because it means the access proof still needs to be synced.
			This can happen when an access request is rejected immediately after being approved.
		*/
	case external.AccessRequestState_ACCESS_REQUEST_STATE_REVOKED:
		return database.OutgoingAccessRequestApproved, nil
	default:
		return "", fmt.Errorf("invalid state for outgoing access request: %s", state)
	}
}

func syncAccessProof(ctx context.Context, configDatabase database.ConfigDatabase, client management.Client, outgoingAccessRequest *database.OutgoingAccessRequest) error {
	response, err := client.GetAccessGrant(ctx, &external.GetAccessGrantRequest{
		ServiceName:          outgoingAccessRequest.ServiceName,
		PublicKeyFingerprint: outgoingAccessRequest.PublicKeyFingerprint,
	})
	if err != nil {
		return err
	}

	remoteProof, err := convertAccessGrantToDatabaseProof(response.AccessGrant)
	if err != nil {
		return fmt.Errorf("parse access proof: %v", err)
	}

	// skip this AccessRequest as it's not the one related to this AccessProof
	if remoteProof.OutgoingAccessRequest.ID != outgoingAccessRequest.ReferenceID {
		return nil
	}

	localProof, err := configDatabase.GetAccessProofForOutgoingAccessRequest(
		ctx,
		outgoingAccessRequest.ID,
	)

	if err != nil {
		if errors.Is(err, database.ErrNotFound) {
			_, errCreate := configDatabase.CreateAccessProof(ctx, outgoingAccessRequest.ID)
			if errCreate != nil {
				return fmt.Errorf("unable to create access proof: %v", errCreate)
			}

			return nil
		}

		return fmt.Errorf("unable to get local access proof: %v", err)
	}

	if remoteProof.RevokedAt.Valid &&
		!localProof.RevokedAt.Valid {
		if _, err := configDatabase.RevokeAccessProof(
			ctx,
			localProof.ID,
			remoteProof.RevokedAt.Time,
		); err != nil {
			return fmt.Errorf("unable to revoke the access proof: %v", err)
		}
	}

	return nil
}

var ErrInvalidTimeStamp = fmt.Errorf("invalid timestamp")

func convertAccessGrantToDatabaseProof(accessGrant *external.AccessGrant) (*database.AccessProof, error) {
	var createdAt time.Time

	if accessGrant.CreatedAt != nil {
		err := accessGrant.CreatedAt.CheckValid()
		if err != nil {
			return nil, ErrInvalidTimeStamp
		}

		createdAt = accessGrant.CreatedAt.AsTime()
	}

	revokedAt := accessGrant.RevokedAt.AsTime()

	err := accessGrant.RevokedAt.CheckValid()
	if err != nil {
		revokedAt = time.Time{}
	}

	dbAccessProof := &database.AccessProof{
		ID:        uint(accessGrant.Id),
		CreatedAt: createdAt,
		RevokedAt: sql.NullTime{
			Time:  revokedAt,
			Valid: !revokedAt.IsZero(),
		},
		OutgoingAccessRequest: &database.OutgoingAccessRequest{
			ID:          uint(accessGrant.AccessRequestId),
			ServiceName: accessGrant.ServiceName,
		},
	}

	if accessGrant.Organization != nil {
		dbAccessProof.OutgoingAccessRequest.Organization = database.Organization{
			SerialNumber: accessGrant.Organization.SerialNumber,
			Name:         accessGrant.Organization.Name,
		}
	}

	return dbAccessProof, nil
}
