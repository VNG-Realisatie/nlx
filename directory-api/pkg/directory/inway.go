// Copyright © VNG Realisatie 2021
// Licensed under the EUPL

package directory

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/blang/semver/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/common/tls"
	directoryapi "go.nlx.io/nlx/directory-api/api"
	"go.nlx.io/nlx/directory-api/domain"
)

const maxServiceCount = 250

func (h *DirectoryService) RegisterInway(ctx context.Context, req *directoryapi.RegisterInwayRequest) (*directoryapi.RegisterInwayResponse, error) {
	resp := &directoryapi.RegisterInwayResponse{}

	organizationInformation, err := h.getOrganizationInformationFromRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get organization information from request: %v", err)
	}

	componentNLXVersion := nlxversion.NewFromGRPCContext(ctx).Version

	// Created at and Updated at time are the same times when registering new inway
	now := h.clock.Now()

	organizationModel, err := domain.NewOrganization(organizationInformation.Name, organizationInformation.SerialNumber)
	if err != nil {
		msg := fmt.Sprintf("validation failed: %s", err.Error())
		return nil, status.New(codes.InvalidArgument, msg).Err()
	}

	managementAPIProxyAddress, err := calculateManagementAPIProxyAddress(componentNLXVersion, req)
	if err != nil {
		h.logger.Error("cannot compute inway proxy address", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "cannot compute inway proxy address")
	}

	if isVersionInRange(">0.127.0", componentNLXVersion) {
		if !isPortOfAddressValid(req.InwayAddress) {
			return nil, status.Errorf(codes.InvalidArgument, "inway address must use port 443 or 8443")
		}

		if req.IsOrganizationInway {
			if !isPortOfAddressValid(req.ManagementApiProxyAddress) {
				return nil, status.Errorf(codes.InvalidArgument, "management API proxy address must use port 443 or 8443")
			}
		}
	}

	inwayModel, err := domain.NewInway(&domain.NewInwayArgs{
		Name:                      req.InwayName,
		Organization:              organizationModel,
		IsOrganizationInway:       req.IsOrganizationInway,
		Address:                   req.InwayAddress,
		ManagementAPIProxyAddress: managementAPIProxyAddress,
		NlxVersion:                componentNLXVersion,
		CreatedAt:                 now,
		UpdatedAt:                 now,
	},
	)
	if err != nil {
		msg := fmt.Sprintf("validation failed: %s", err.Error())
		return nil, status.New(codes.InvalidArgument, msg).Err()
	}

	err = h.repository.RegisterInway(inwayModel)
	if err != nil {
		h.logger.Error("register inway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
		return nil, status.New(codes.Internal, "failed to register inway").Err()
	}

	err = h.setOrganizationInway(ctx, organizationInformation, inwayModel)
	if err != nil {
		h.logger.Error("failed to configure organization inway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
		return nil, status.New(codes.Internal, "failed to set organization inway ").Err()
	}

	if len(req.Services) > maxServiceCount {
		return nil, status.New(codes.InvalidArgument, fmt.Sprintf("inway registers more services than allowed (max. %d)", maxServiceCount)).Err()
	}

	for _, s := range req.Services {
		err = h.registerService(organizationInformation, req.InwayAddress, s)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

func (h *DirectoryService) setOrganizationInway(ctx context.Context, organizationInformation *tls.OrganizationInformation, inwayModel *domain.Inway) error {
	if inwayModel.IsOrganizationInway() {
		h.logger.Debug("is organization inway", zap.String("inway", inwayModel.ToString()))
		err := h.repository.SetOrganizationInway(ctx, organizationInformation.SerialNumber, inwayModel.Address())

		if err != nil {
			h.logger.Error("set organization inway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
			return status.New(codes.Internal, "failed to set organization inway").Err()
		}
	} else {
		h.logger.Debug("is not organization inway", zap.String("inway", inwayModel.ToString()))
		err := h.repository.ClearIfSetAsOrganizationInway(ctx, organizationInformation.SerialNumber, inwayModel.Address())

		if err != nil {
			h.logger.Error("failed to execute ClearIfSetAsOrganizationInway", zap.String("inway", inwayModel.ToString()), zap.Error(err))
			return status.New(codes.Internal, "failed to update organization inway settings").Err()
		}
	}

	return nil
}

func (h *DirectoryService) registerService(organizationInformation *tls.OrganizationInformation, inwayAddress string, service *directoryapi.RegisterInwayRequest_RegisterService) error {
	serviceSpecificationType := getAPISpecificationTypeForService(
		h.httpClient,
		h.logger,
		service.ApiSpecificationDocumentUrl,
		inwayAddress,
		service.Name,
	)

	organization, err := domain.NewOrganization(organizationInformation.Name, organizationInformation.SerialNumber)
	if err != nil {
		msg := fmt.Sprintf("validation for organization with serial number '%s' failed: %s", organizationInformation.SerialNumber, err.Error())
		return status.New(codes.InvalidArgument, msg).Err()
	}

	serviceModel, err := domain.NewService(
		&domain.NewServiceArgs{
			Name:                 service.Name,
			Organization:         organization,
			Internal:             service.Internal,
			DocumentationURL:     service.DocumentationUrl,
			APISpecificationType: serviceSpecificationType,
			PublicSupportContact: service.PublicSupportContact,
			TechSupportContact:   service.TechSupportContact,
			Costs: &domain.NewServiceCostsArgs{
				OneTime: uint(service.OneTimeCosts),
				Monthly: uint(service.MonthlyCosts),
				Request: uint(service.RequestCosts),
			},
		},
	)
	if err != nil {
		msg := fmt.Sprintf("validation for service named '%s' failed: %s", service.Name, err.Error())
		return status.New(codes.InvalidArgument, msg).Err()
	}

	err = h.repository.RegisterService(serviceModel)
	if err != nil {
		h.logger.Error("failed to register service", zap.Error(err))
		return status.New(codes.Internal, "database error").Err()
	}

	return nil
}

func getAPISpecificationTypeForService(httpClient *http.Client, logger *zap.Logger, apiSpecificationDocumentURL, inwayAddress, serviceName string) domain.SpecificationType {
	if len(apiSpecificationDocumentURL) < 1 {
		return ""
	}

	specificationType, err := getAPISpecsTypeViaInway(httpClient, inwayAddress, serviceName)
	if err != nil {
		logger.Error(
			"invalid api specification document provided by inway",
			zap.String("api specification document url", apiSpecificationDocumentURL),
			zap.Error(err),
		)

		return ""
	}

	return domain.SpecificationType(specificationType)
}

func isPortOfAddressValid(address string) bool {
	_, port, err := net.SplitHostPort(address)
	if err != nil {
		return false
	}

	return port == "443" || port == "8443"
}

func isVersionInRange(wantedVersionRange, componentVersion string) bool {
	v, err := semver.Parse(strings.TrimPrefix(componentVersion, "v"))
	if err != nil {
		return false
	}

	return semver.MustParseRange(wantedVersionRange)(v)
}

func calculateManagementAPIProxyAddress(nlxComponentVersion string, req *directoryapi.RegisterInwayRequest) (string, error) {
	if !req.IsOrganizationInway {
		return "", nil
	}

	if isVersionInRange(">0.127.0", nlxComponentVersion) {
		return req.ManagementApiProxyAddress, nil
	}

	if req.InwayAddress == "" {
		return "", fmt.Errorf("empty inway address provided")
	}

	host, port, err := net.SplitHostPort(req.InwayAddress)
	if err != nil {
		return "", fmt.Errorf("invalid format for inway address: %w", err)
	}

	portNum, err := strconv.Atoi(port)
	if err != nil {
		return "", fmt.Errorf("invalid format for inway address port: %w", err)
	}

	return fmt.Sprintf("%s:%d", host, portNum+1), nil
}
