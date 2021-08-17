// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { fireEvent, screen } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../../test-utils'
import Services from './index'

const createComponent = (services) => {
  return renderWithProviders(<Services services={services} />)
}

test('no services available', async () => {
  createComponent([])

  fireEvent.click(screen.getByText('Requestable services'))

  expect(
    await screen.findByText('No services have been connected'),
  ).toBeInTheDocument()
})

test('listing the services', async () => {
  createComponent([
    {
      service: 'My Service',
      organization: 'My Organization',
    },
  ])

  fireEvent.click(screen.getByText('Requestable services'))

  expect(await screen.findByText('My Service')).toBeInTheDocument()
  expect(await screen.findByText('My Organization')).toBeInTheDocument()
})
