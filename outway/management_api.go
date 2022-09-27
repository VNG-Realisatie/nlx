// Copyright © VNG Realisatie 2022
// Licensed under the EUPL

package outway

import (
	"context"
	"time"

	"github.com/jpillora/backoff"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	common_version "go.nlx.io/nlx/common/version"
	managementapi "go.nlx.io/nlx/management-api/api"
)

const retryFactorManagementAPI = 10
const maxRetryDurationManagementAPI = 20 * time.Second
const minRetryDurationManagementAPI = 100 * time.Millisecond
const announceToManagementAPIInterval = 10 * time.Second

//nolint:dupl // looks the same as announceToDirectory but is not
func (o *Outway) announceToManagementAPI(ctx context.Context) {
	expBackOff := &backoff.Backoff{
		Min:    minRetryDurationManagementAPI,
		Factor: retryFactorManagementAPI,
		Max:    maxRetryDurationManagementAPI,
	}

	sleepDuration := announceToManagementAPIInterval

	for {
		select {
		case <-ctx.Done():
			o.logger.Info("stopping management api announcement")
			return
		case <-time.After(sleepDuration):
			err := o.registerToManagementAPI(ctx)
			if err != nil {
				if errStatus, ok := status.FromError(err); ok && errStatus.Code() == codes.Unavailable {
					o.logger.Info("waiting for management api...", zap.Error(err))

					sleepDuration = expBackOff.Duration()

					continue
				}

				o.logger.Error("failed to register to management api", zap.Error(err))
			}

			o.logger.Info("management api registration successful")

			sleepDuration = announceToManagementAPIInterval

			expBackOff.Reset()
		}
	}
}

func (o *Outway) registerToManagementAPI(ctx context.Context) error {
	publicKeyCert, err := o.orgCert.PublicKeyPEM()
	if err != nil {
		return err
	}

	registerOutwayRequest := &managementapi.RegisterOutwayRequest{
		Name:           o.name,
		SelfAddressApi: o.addressAPI,
		Version:        common_version.BuildVersion,
		PublicKeyPem:   publicKeyCert,
	}

	_, err = o.managementAPIClient.RegisterOutway(ctx, registerOutwayRequest)
	if err != nil {
		return err
	}

	return nil
}
