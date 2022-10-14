// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Routes, MemoryRouter } from 'react-router-dom'
import { fireEvent, screen, waitFor, within } from '@testing-library/react'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementServiceApi } from '../../../api'
import InwayDetailPage from './index'

test('display inway details', async () => {
  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListInways = jest
    .fn()
    .mockResolvedValue({
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

  expect(screen.getByTestId('inway-specs')).toBeInTheDocument()
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

test('remove an Inway', async () => {
  configure({ safeDescriptors: false })

  const managementApiClient = new ManagementServiceApi()

  managementApiClient.managementServiceListInways = jest
    .fn()
    .mockResolvedValue({
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

  managementApiClient.managementServiceDeleteInway = jest
    .fn()
    .mockRejectedValueOnce({ response: { status: 403 } })
    .mockResolvedValue({})

  const rootStore = new RootStore({
    managementApiClient,
  })

  jest.spyOn(rootStore.inwayStore, 'removeInway')

  await rootStore.inwayStore.fetchInways()

  renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <MemoryRouter initialEntries={['/my-inway']}>
        <Routes>
          <Route
            path="/inways-and-outways/inways"
            element={<div>inways page</div>}
          />
          <Route path=":name" element={<InwayDetailPage />} />
        </Routes>
      </MemoryRouter>
    </StoreProvider>,
  )

  fireEvent.click(screen.getByTitle('Remove Inway'))

  let confirmModal = screen.getByRole('dialog')
  let okButton = within(confirmModal).getByText('Remove')

  fireEvent.click(okButton)

  expect(await screen.findByRole('alert')).toHaveTextContent(
    "Failed to remove the inwayYou don't have the required permission.",
  )

  await waitFor(() =>
    expect(rootStore.inwayStore.removeInway).toHaveBeenCalledWith('my-inway'),
  )

  fireEvent.click(screen.getByTitle('Remove Inway'))

  confirmModal = screen.getByRole('dialog')
  okButton = within(confirmModal).getByText('Remove')

  fireEvent.click(okButton)

  expect(
    await screen.findByText('The inway has been removed'),
  ).toBeInTheDocument()
})
