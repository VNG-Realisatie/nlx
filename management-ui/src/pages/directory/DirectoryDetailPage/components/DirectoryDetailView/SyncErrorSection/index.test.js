// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import OutgoingAccessRequestSyncErrorModel, {
  SYNC_ERROR_INTERNAL,
} from '../../../../../../stores/models/OutgoingAccessRequestSyncErrorModel'
import SyncErrorSection from './index'

test('render the error message', async () => {
  const syncError = new OutgoingAccessRequestSyncErrorModel({
    syncErrorData: {
      organizationSerialNumber: '00000000000000000001',
      error: SYNC_ERROR_INTERNAL,
    },
  })

  renderWithProviders(<SyncErrorSection syncError={syncError} />)

  const errorMessage = screen.getByText(
    'Internal error while trying to retrieve the current state of your access requests. Please consult your system administrator.',
  )

  expect(errorMessage).toBeInTheDocument()
})
