// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, MemoryRouter, Routes } from 'react-router-dom'
import { screen, fireEvent, waitFor, within } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../../test-utils'
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

  renderWithAllProviders(
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

  renderWithAllProviders(
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

test('remove an Outway', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementApi()

  managementApiClient.managementListOutways = jest.fn().mockResolvedValue({
    outways: [
      {
        name: 'my-outway',
      },
    ],
  })

  managementApiClient.managementDeleteOutway = jest
    .fn()
    .mockRejectedValueOnce({ response: { status: 403 } })
    .mockResolvedValue({})

  const rootStore = new RootStore({
    managementApiClient,
  })

  jest.spyOn(rootStore.outwayStore, 'removeOutway')

  await rootStore.outwayStore.fetchAll()

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <MemoryRouter initialEntries={['/my-outway']}>
        <Routes>
          <Route
            path="/inways-and-outways/outways"
            element={<div>outways page</div>}
          />
          <Route path=":name" element={<OutwayDetailPage />} />
        </Routes>
      </MemoryRouter>
    </StoreProvider>,
  )

  fireEvent.click(screen.getByTitle('Remove outway'))

  let confirmModal = screen.getByRole('dialog')
  let okButton = within(confirmModal).getByText('Remove')

  fireEvent.click(okButton)

  expect(await screen.findByRole('alert')).toHaveTextContent(
    "Failed to remove the outwayYou don't have the required permission.",
  )

  await waitFor(() =>
    expect(rootStore.outwayStore.removeOutway).toHaveBeenCalledWith(
      'my-outway',
    ),
  )

  fireEvent.click(screen.getByTitle('Remove outway'))

  confirmModal = screen.getByRole('dialog')
  okButton = within(confirmModal).getByText('Remove')

  fireEvent.click(okButton)

  expect(
    await screen.findByText('The outway has been removed'),
  ).toBeInTheDocument()
})
