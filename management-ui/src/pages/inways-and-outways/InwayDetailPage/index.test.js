// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Routes, MemoryRouter } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import InwayDetailPage from './index'

/* eslint-disable react/prop-types */
jest.mock('./InwayDetailPageView', () => ({ inway }) => (
  <div data-testid="inway-details">{inway.name}</div>
))
/* eslint-enable react/prop-types */

test('display inway details', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListInways = jest.fn().mockResolvedValue({
    inways: [
      {
        name: 'my-inway',
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

  const rootStore = new RootStore({
    managementApiClient,
  })

  await rootStore.inwayStore.fetchInways()

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <MemoryRouter initialEntries={['/my-inway']}>
        <Routes>
          <Route path=":name" element={<InwayDetailPage />} />
        </Routes>
      </MemoryRouter>
    </StoreProvider>,
  )
  expect(screen.getByTestId('inway-details')).toHaveTextContent('my-inway')
})

test('display a non-existing inway', async () => {
  const rootStore = new RootStore({})

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <MemoryRouter initialEntries={['/my-inway']}>
        <Routes>
          <Route path=":name" element={<InwayDetailPage />} />
        </Routes>
      </MemoryRouter>
    </StoreProvider>,
  )

  const message = await screen.findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this inway')

  const closeButton = await screen.findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
