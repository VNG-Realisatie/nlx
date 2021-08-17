// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React from 'react'
import { Router } from 'react-router-dom'
import { fireEvent, screen } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../../../../../test-utils'
import Services from './index'

const createComponent = (services) => {
  const history = createMemoryHistory({
    initialEntries: ['/'],
  })

  renderWithProviders(
    <Router history={history}>
      <Services services={services} />
    </Router>,
  )

  return {
    history,
  }
}

test('no services available', async () => {
  createComponent([])

  fireEvent.click(screen.getByText('Requestable services'))

  expect(
    await screen.findByText('No services have been connected'),
  ).toBeInTheDocument()
})

test('listing the services', async () => {
  const { history } = createComponent([
    {
      service: 'My Service',
      organization: 'My Organization',
    },
  ])

  fireEvent.click(screen.getByText('Requestable services'))

  const service = await screen.findByText('My Service')
  expect(service).toBeInTheDocument()
  expect(await screen.findByText('My Organization')).toBeInTheDocument()

  fireEvent.click(service)

  expect(history.location.pathname).toEqual(
    '/directory/My Organization/My Service',
  )
})
