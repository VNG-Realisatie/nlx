// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//

import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../../../test-utils'
import DirectoryServiceModel from '../../../../../../../../stores/models/DirectoryServiceModel'
import CancelRequestAccessDetails from './index'

test('Request Access Details', () => {
  const service = new DirectoryServiceModel({
    serviceData: {
      serviceName: 'service',
      organization: {
        name: 'organization',
      },
    },
  })

  renderWithProviders(<CancelRequestAccessDetails service={service} />)

  expect(screen.queryByText('service')).toBeInTheDocument()
  expect(screen.queryByText('organization')).toBeInTheDocument()
})
