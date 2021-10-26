// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { fireEvent, screen } from '@testing-library/react'
import { renderWithProviders, waitFor } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import Routes from '../../../routes'

jest.mock('../../../components/PageTemplate')
jest.mock('./Inways', () => () => <p data-testid="inways-list">mock inways</p>)

function renderPage(rootStore) {
  const history = createMemoryHistory({
    initialEntries: ['/inways-and-outways'],
  })

  return renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{}}>
        <StoreProvider rootStore={rootStore}>
          <Routes />
        </StoreProvider>
      </UserContextProvider>
    </Router>,
  )
}

test('the InwaysAndOutwaysPage page', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [
      {
        name: 'name',
        version: 'version',
        hostname: 'hostname',
        selfAddress: 'self-address',
        services: [
          {
            name: 'service-1',
          },
        ],
      },
    ],
  })

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'Outway Name',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        version: '0.0.42',
      },
    ],
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  renderPage(rootStore)

  const showInwaysButton = screen.getByLabelText('Show Inways')
  expect(showInwaysButton.getAttribute('href')).toBe(
    '/inways-and-outways/inways',
  )

  const showOutwaysButton = screen.getByLabelText('Show Outways')
  expect(showOutwaysButton.getAttribute('href')).toBe(
    '/inways-and-outways/outways',
  )

  expect(screen.getByRole('progressbar')).toBeInTheDocument()
  expect(() => screen.getByTestId('inways-list')).toThrow()

  await waitFor(() =>
    expect(screen.getByTestId('inways-list')).toHaveTextContent('mock inways'),
  )

  fireEvent.click(showOutwaysButton)

  await waitFor(() =>
    expect(screen.getByTestId('outways-list')).toBeInTheDocument(),
  )
})

test('failed to load inways', async () => {
  const managementApiClient = new ManagementApi()
  managementApiClient.managementListInways = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'name',
        ipAddress: '127.0.0.1',
        publicKeyPEM: 'public-key-pem',
        version: 'version',
      },
    ],
  })

  const rootStore = new RootStore({
    managementApiClient,
  })

  renderPage(rootStore)

  expect(() => screen.getByTestId('inways-list')).toThrow()
  expect(
    await screen.findByText(/^Failed to load the inways$/),
  ).toBeInTheDocument()
})
