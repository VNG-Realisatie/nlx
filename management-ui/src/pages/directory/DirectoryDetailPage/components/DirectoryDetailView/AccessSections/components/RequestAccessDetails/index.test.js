// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../../../test-utils'
import DirectoryServiceModel from '../../../../../../../../stores/models/DirectoryServiceModel'
import RequestAccessDetails from './index'

test('Request Access Details', () => {
  const service = new DirectoryServiceModel({
    serviceData: {
      serviceName: 'service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'organization',
      },
    },
  })

  const { container } = renderWithProviders(
    <RequestAccessDetails service={service} />,
  )

  expect(screen.queryByText(/One time costs/)).not.toBeInTheDocument()
  expect(screen.queryByText(/Monthly costs/)).not.toBeInTheDocument()
  expect(screen.queryByText(/Cost per request/)).not.toBeInTheDocument()

  service.update({
    serviceData: {
      oneTimeCosts: 500,
      monthlyCosts: 1000,
      requestCosts: 25000,
    },
  })

  expect(container).toHaveTextContent(/One time costs€ 5,00/)
  expect(container).toHaveTextContent(/Monthly costs€ 10,00/)
  expect(container).toHaveTextContent(/Cost per request€ 250,00/)
})
