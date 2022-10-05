// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

func (h *DirectoryService) ListInOutwayStatistics(ctx context.Context, _ *directoryapi.ListInOutwayStatisticsRequest) (*directoryapi.ListInOutwayStatisticsResponse, error) {
	h.logger.Info("rpc request ListOrganizations")

	versionStatistics, err := h.repository.ListVersionStatistics(ctx)
	if err != nil {
		h.logger.Error("failed to select version statistics from storage", zap.Error(err))
		return nil, status.New(codes.Internal, "Storage error.").Err()
	}

	return convertModelToResponse(versionStatistics), nil
}

func convertModelToResponse(model []*domain.VersionStatistics) *directoryapi.ListInOutwayStatisticsResponse {
	result := &directoryapi.ListInOutwayStatisticsResponse{}

	for _, statistics := range model {
		modelType, err := modelTypeToResponseType(statistics.GatewayType())
		if err != nil {
			continue
		}

		result.Versions = append(result.Versions, &directoryapi.ListInOutwayStatisticsResponse_Statistics{
			Type:    modelType,
			Version: statistics.Version(),
			Amount:  statistics.Amount(),
		})
	}

	return result
}

func modelTypeToResponseType(statisticsType domain.VersionStatisticsType) (directoryapi.ListInOutwayStatisticsResponse_Statistics_Type, error) {
	if statisticsType == domain.TypeInway {
		return directoryapi.ListInOutwayStatisticsResponse_Statistics_TYPE_INWAY, nil
	} else if statisticsType == domain.TypeOutway {
		return directoryapi.ListInOutwayStatisticsResponse_Statistics_TYPE_OUTWAY, nil
	}

	return directoryapi.ListInOutwayStatisticsResponse_Statistics_TYPE_UNSPECIFIED, fmt.Errorf("unknown type '%s', expected '%v' or '%v'", statisticsType, domain.TypeInway, domain.TypeOutway)
}
