// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, MemoryRouter, Routes } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../../test-utils'
import { ManagementApi } from '../../../api'
import { RootStore, StoreProvider } from '../../../stores'
import OutwayDetailPage from './index'

/* eslint-disable react/prop-types */
jest.mock('./OutwayDetailPageView', () => ({ outway }) => (
  <div data-testid="outway-details">{outway.name}</div>
))
/* eslint-enable react/prop-types */

test('display outway details', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'my-outway',
      },
    ],
  })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.outwayStore.fetchAll()

  renderWithProviders(
    <MemoryRouter initialEntries={['/my-outway']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":name" element={<OutwayDetailPage />} />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )

  expect(await screen.findByTestId('outway-details')).toHaveTextContent(
    'my-outway',
  )
})

test('display a non-existing outway', async () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  renderWithProviders(
    <MemoryRouter initialEntries={['/my-outway']}>
      <StoreProvider rootStore={rootStore}>
        <Routes>
          <Route path=":name" element={<OutwayDetailPage />} />
        </Routes>
      </StoreProvider>
    </MemoryRouter>,
  )

  const message = await screen.findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this outway')

  const closeButton = await screen.findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
