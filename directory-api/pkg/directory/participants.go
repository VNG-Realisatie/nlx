// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

func (h *DirectoryService) ListParticipants(ctx context.Context, _ *emptypb.Empty) (*directoryapi.ListParticipantsResponse, error) {
	h.logger.Info("rpc request ListParticipants")

	participants, err := h.repository.ListParticipants(ctx)
	if err != nil {
		h.logger.Error("failed to select participants from db", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	return convertFromDatabaseParticipants(participants), nil
}

func convertFromDatabaseParticipants(models []*domain.Participant) *directoryapi.ListParticipantsResponse {
	result := &directoryapi.ListParticipantsResponse{
		Participants: make([]*directoryapi.ListParticipantsResponse_Participant, len(models)),
	}

	for i, p := range models {
		result.Participants[i] = &directoryapi.ListParticipantsResponse_Participant{
			Organization: &directoryapi.Organization{
				Name:         p.Organization().Name(),
				SerialNumber: p.Organization().SerialNumber(),
			},
			Statistics: &directoryapi.ListParticipantsResponse_Participant_Statistics{
				Inways:   uint32(p.Statistics().Inways()),
				Outways:  uint32(p.Statistics().Outways()),
				Services: uint32(p.Statistics().Services()),
			},
			CreatedAt: timestamppb.New(p.CreatedAt()),
		}
	}

	return result
}
