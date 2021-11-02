package pgadapter

import (
	"context"
	"errors"

	"go.nlx.io/nlx/directory-api/domain/directory/storage"
	"go.uber.org/zap"
)

// ClearIfSetAsOrganizationInway clears the inway for the given organization.
// This method should be called if IsOrganizationInway is false in the request, to ensure the directory has this correctly set as well
func (r *PostgreSQLRepository) ClearIfSetAsOrganizationInway(ctx context.Context, serialNumber, selfAddress string) error {
	organizationSelfAddress, err := r.GetOrganizationInwayAddress(ctx, serialNumber)
	if err != nil {
		if errors.Is(err, storage.ErrOrganizationNotFound) {
			return nil
		}

		return err
	}

	if selfAddress == organizationSelfAddress {
		r.logger.Warn("unexpected state: inway was incorrectly set as organization inway ",
			zap.String("inway self address", selfAddress),
			zap.String("organization inway self address", organizationSelfAddress))

		return r.ClearOrganizationInway(ctx, serialNumber)
	}

	return nil
}
