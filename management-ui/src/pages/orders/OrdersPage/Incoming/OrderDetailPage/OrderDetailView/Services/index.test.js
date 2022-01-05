// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../../../../../test-utils'
import Services from './index'

const renderServices = (services) => {
  const history = createMemoryHistory({
    initialEntries: ['/'],
  })

  renderWithProviders(
    <HistoryRouter history={history}>
      <Services services={services} />
    </HistoryRouter>,
  )

  return {
    history,
  }
}

test('no services available', async () => {
  renderServices([])

  fireEvent.click(screen.getByText('Requestable services'))

  expect(
    await screen.findByText('No services have been connected'),
  ).toBeInTheDocument()
})

test('listing the services', async () => {
  const { history } = renderServices([
    {
      service: 'My Service',
      organization: {
        serialNumber: '00000000000000000001',
        name: 'My Organization',
      },
    },
  ])

  fireEvent.click(screen.getByText('Requestable services'))

  const service = await screen.findByText('My Service')
  expect(service).toBeInTheDocument()
  expect(
    await screen.findByText('My Organization (00000000000000000001)'),
  ).toBeInTheDocument()

  fireEvent.click(service)

  expect(history.location.pathname).toEqual(
    '/directory/00000000000000000001/My Service',
  )
})
