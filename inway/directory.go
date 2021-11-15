package inway

import (
	"context"
	"fmt"
	"time"

	"github.com/jpillora/backoff"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	"go.nlx.io/nlx/directory-registration-api/registrationapi"
)

const retryFactorDirectory = 10
const maxRetryDurationDirectory = 20 * time.Second
const minRetryDurationDirectory = 100 * time.Millisecond
const announceToDirectoryInterval = 10 * time.Second

func (i *Inway) announceToDirectory(ctx context.Context) {
	expBackOff := &backoff.Backoff{
		Min:    maxRetryDurationDirectory,
		Factor: retryFactorDirectory,
		Max:    minRetryDurationDirectory,
	}

	sleepDuration := announceToDirectoryInterval

	for {
		select {
		case <-ctx.Done():
			i.logger.Info("stopping directory announcement")
			return
		case <-time.After(sleepDuration):
			err := i.registerToDirectory(ctx)
			if err != nil {
				if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
					i.logger.Info("waiting for directory...", zap.Error(err))

					sleepDuration = expBackOff.Duration()

					continue
				}

				i.logger.Error("failed to register to directory", zap.Error(err))
			}

			i.logger.Info("directory registration successful")

			sleepDuration = announceToDirectoryInterval

			expBackOff.Reset()
		}
	}
}

func (i *Inway) registerToDirectory(ctx context.Context) error {
	protoServiceDetails := []*registrationapi.RegisterInwayRequest_RegisterService{}

	for _, service := range i.services {
		protoServiceDetails = append(protoServiceDetails, &registrationapi.RegisterInwayRequest_RegisterService{
			Name:                        service.Name,
			Internal:                    service.Internal,
			DocumentationUrl:            service.DocumentationURL,
			ApiSpecificationDocumentUrl: service.APISpecificationDocumentURL,
			PublicSupportContact:        service.PublicSupportContact,
			TechSupportContact:          service.TechSupportContact,
			OneTimeCosts:                service.OneTimeCosts,
			MonthlyCosts:                service.MonthlyCosts,
			RequestCosts:                service.RequestCosts,
		})
	}

	registerInwayRequest := &registrationapi.RegisterInwayRequest{
		InwayName:           i.name,
		InwayAddress:        i.address,
		IsOrganizationInway: i.isOrganizationInway,
		Services:            protoServiceDetails,
	}

	nlxVersion := nlxversion.NewGRPCContext(ctx, "inway")
	i.logger.Debug("registering inway", zap.Any("RegisterInwayRequest", registerInwayRequest), zap.Any("nlxVersion", nlxVersion))

	resp, err := i.directoryRegistrationClient.RegisterInway(nlxVersion, registerInwayRequest)
	if err != nil {
		return err
	}

	if resp.Error != "" {
		i.logger.Error(fmt.Sprintf("failed to register to directory: %s", resp.Error))
		return fmt.Errorf(resp.Error)
	}

	return nil
}
