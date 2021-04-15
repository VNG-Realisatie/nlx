// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package inspectionservice

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"go.nlx.io/nlx/directory-inspection-api/inspectionapi"
	"go.nlx.io/nlx/directory-inspection-api/pkg/database"
)

func (h *InspectionService) ListInOutwayStatistics(ctx context.Context, _ *emptypb.Empty) (*inspectionapi.ListInOutwayStatisticsResponse, error) {
	h.logger.Info("rpc request ListOrganizations")

	versionStatistics, err := h.db.ListVersionStatistics(ctx)
	if err != nil {
		h.logger.Error("failed to select version statistics from db", zap.Error(err))
		return nil, status.New(codes.Internal, "Database error.").Err()
	}

	return convertModelToResponse(versionStatistics), nil
}

func convertModelToResponse(model []*database.VersionStatistics) *inspectionapi.ListInOutwayStatisticsResponse {
	result := &inspectionapi.ListInOutwayStatisticsResponse{}

	for _, statistics := range model {
		modelType, err := modelTypeToResponseType(statistics.Type)
		if err != nil {
			continue
		}

		result.Versions = append(result.Versions, &inspectionapi.ListInOutwayStatisticsResponse_Statistics{
			Type:    modelType,
			Version: statistics.Version,
			Amount:  statistics.Amount,
		})
	}

	return result
}

func modelTypeToResponseType(statisticsType database.VersionStatisticsType) (inspectionapi.ListInOutwayStatisticsResponse_Statistics_Type, error) {
	if statisticsType == database.TypeInway {
		return inspectionapi.ListInOutwayStatisticsResponse_Statistics_INWAY, nil
	} else if statisticsType == database.TypeOutway {
		return inspectionapi.ListInOutwayStatisticsResponse_Statistics_OUTWAY, nil
	}

	return inspectionapi.ListInOutwayStatisticsResponse_Statistics_INWAY, fmt.Errorf("unknown type '%s', expected '%v' or '%v'", statisticsType, database.TypeInway, database.TypeOutway)
}
