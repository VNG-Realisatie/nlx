// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package outway

import (
	"context"
	"fmt"
	"time"

	"github.com/jpillora/backoff"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.nlx.io/nlx/common/nlxversion"
	directoryapi "go.nlx.io/nlx/directory-api/api"
)

const retryFactorDirectory = 10
const maxRetryDurationDirectory = 20 * time.Second
const minRetryDurationDirectory = 100 * time.Millisecond
const announceToDirectoryInterval = 10 * time.Second

func (o *Outway) announceToDirectory(ctx context.Context) {
	expBackOff := &backoff.Backoff{
		Min:    maxRetryDurationDirectory,
		Factor: retryFactorDirectory,
		Max:    minRetryDurationDirectory,
	}

	sleepDuration := announceToDirectoryInterval

	for {
		select {
		case <-ctx.Done():
			o.logger.Info("stopping directory announcement")
			return
		case <-time.After(sleepDuration):
			err := o.registerToDirectory(ctx)
			if err != nil {
				if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
					o.logger.Info("waiting for directory...", zap.Error(err))

					sleepDuration = expBackOff.Duration()

					continue
				}

				o.logger.Error("failed to register to directory", zap.Error(err))
			}

			o.logger.Info("directory registration successful")

			sleepDuration = announceToDirectoryInterval

			expBackOff.Reset()
		}
	}
}

func (o *Outway) registerToDirectory(ctx context.Context) error {
	registerOutwayRequest := &directoryapi.RegisterOutwayRequest{
		Name: o.name,
	}

	nlxVersion := nlxversion.NewGRPCContext(ctx, "outway")
	o.logger.Debug("registering outway", zap.Any("RegisterOutwayRequest", registerOutwayRequest), zap.Any("nlxVersion", nlxVersion))

	resp, err := o.directoryClient.RegisterOutway(nlxVersion, registerOutwayRequest)
	if err != nil {
		return err
	}

	if resp.Error != "" {
		o.logger.Error(fmt.Sprintf("failed to register to directory: %s", resp.Error))
		return fmt.Errorf(resp.Error)
	}

	return nil
}
