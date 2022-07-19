// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, fireEvent, screen } from '@testing-library/react'
import {
  Route,
  Routes,
  unstable_HistoryRouter as HistoryRouter,
  MemoryRouter,
} from 'react-router-dom'
import { createMemoryHistory } from 'history'
import UserContext from '../../../user-context'
import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import EditServicePage from './index'

jest.mock(
  '../../../components/PageTemplate/OrganizationInwayCheck',
  () => () => null,
)

jest.mock(
  '../../../components/PageTemplate/OrganizationEmailAddressCheck',
  () => () => null,
)

// eslint-disable-next-line react/prop-types
jest.mock('../../../components/ServiceForm', () => ({ onSubmitHandler }) => (
  <form
    onSubmit={() => onSubmitHandler({ name: 'mock-service' })}
    data-testid="form"
  >
    <button type="submit" />
  </form>
))

describe('the EditServicePage', () => {
  beforeAll(() => {
    global.scrollTo = jest.fn()
  })
  afterEach(() => {
    jest.resetModules()
  })

  it('before the service has been loaded', async () => {
    jest.useFakeTimers()

    const rootStore = new RootStore({
      managementApiClient: new ManagementApi(),
    })

    const userContext = { user: { id: '42' } }
    renderWithProviders(
      <MemoryRouter initialEntries={['/mock-service/edit-service']}>
        <UserContext.Provider value={userContext}>
          <StoreProvider rootStore={rootStore}>
            <Routes>
              <Route path=":name/edit-service" element={<EditServicePage />} />
            </Routes>
          </StoreProvider>
        </UserContext.Provider>
      </MemoryRouter>,
    )

    expect(await screen.findByRole('progressbar')).toBeTruthy()
    const linkBack = screen.getByLabelText(/Back/)
    expect(linkBack.getAttribute('href')).toBe('/services/mock-service')
  })

  it('when fetching the services fails', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementCreateService = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))

    const rootStore = new RootStore({
      managementApiClient,
    })

    const userContext = { user: { id: '42' } }
    renderWithProviders(
      <MemoryRouter initialEntries={['/mock-service/edit-service']}>
        <UserContext.Provider value={userContext}>
          <StoreProvider rootStore={rootStore}>
            <Routes>
              <Route path=":name/edit-service" element={<EditServicePage />} />
            </Routes>
          </StoreProvider>
        </UserContext.Provider>
      </MemoryRouter>,
    )

    expect(await screen.findByRole('alert')).toBeTruthy()
    expect(screen.queryByRole('alert').textContent).toBe(
      'Failed to load the service',
    )
  })

  it('after the service has been fetched', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementListServices = jest.fn().mockResolvedValue({
      services: [{ name: 'mock-service' }],
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    await rootStore.servicesStore.fetchAll()

    const userContext = { user: { id: '42' } }
    renderWithProviders(
      <MemoryRouter initialEntries={['/mock-service/edit-service']}>
        <UserContext.Provider value={userContext}>
          <StoreProvider rootStore={rootStore}>
            <Routes>
              <Route path=":name/edit-service" element={<EditServicePage />} />
            </Routes>
          </StoreProvider>
        </UserContext.Provider>
      </MemoryRouter>,
    )

    expect(await screen.findByTestId('form')).toBeTruthy()
  })

  it('successfully submitting the form', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementUpdateService = jest.fn().mockResolvedValue({
      name: 'mock-service',
    })
    managementApiClient.managementListServices = jest.fn().mockResolvedValue({
      services: [{ name: 'mock-service' }],
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    await rootStore.servicesStore.fetchAll()

    const history = createMemoryHistory({
      initialEntries: ['/mock-service/edit-service'],
    })

    renderWithProviders(
      <HistoryRouter history={history}>
        <StoreProvider rootStore={rootStore}>
          <Routes>
            <Route path="/:name/edit-service" element={<EditServicePage />} />
            <Route path="*" element={null} />
          </Routes>
        </StoreProvider>
      </HistoryRouter>,
    )

    const editServiceForm = await screen.findByTestId('form')
    await act(async () => {
      fireEvent.submit(editServiceForm)
    })

    expect(managementApiClient.managementUpdateService).toHaveBeenCalled()
    expect(history.location.pathname).toEqual('/services/mock-service')
    expect(history.location.search).toEqual('?lastAction=edited')
  })

  it('re-submitting the form when the previous submission went wrong', async () => {
    const managementApiClient = new ManagementApi()

    managementApiClient.managementUpdateService = jest
      .fn()
      .mockResolvedValue({
        name: 'mock-service',
      })
      .mockRejectedValueOnce(new Error('arbitrary error'))

    managementApiClient.managementListServices = jest.fn().mockResolvedValue({
      services: [{ name: 'mock-service' }],
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    await rootStore.servicesStore.fetchAll()

    const history = createMemoryHistory({
      initialEntries: ['/mock-service/edit-service'],
    })

    renderWithProviders(
      <HistoryRouter history={history}>
        <StoreProvider rootStore={rootStore}>
          <Routes>
            <Route path=":name/edit-service" element={<EditServicePage />} />
            <Route path="*" element={null} />
          </Routes>
        </StoreProvider>
      </HistoryRouter>,
    )

    const editServiceForm = await screen.findByTestId('form')

    await act(async () => {
      await fireEvent.submit(editServiceForm)
    })

    expect(managementApiClient.managementUpdateService).toHaveBeenCalledTimes(1)
    expect(screen.queryByRole('alert')).toBeTruthy()
    expect(screen.queryByRole('alert')).toHaveTextContent(
      'Failed to update the service',
    )

    await act(async () => {
      fireEvent.submit(editServiceForm)
    })

    expect(managementApiClient.managementUpdateService).toHaveBeenCalledTimes(2)

    expect(history.location.pathname).toEqual('/services/mock-service')
    expect(history.location.search).toEqual('?lastAction=edited')
  })

  it('submitting when the HTTP response is not ok', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementUpdateService = jest.fn().mockRejectedValue({
      response: {
        status: 403,
      },
    })

    managementApiClient.managementListServices = jest.fn().mockResolvedValue({
      services: [{ name: 'mock-service' }],
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    await rootStore.servicesStore.fetchAll()

    renderWithProviders(
      <MemoryRouter initialEntries={['/mock-service/edit-service']}>
        <StoreProvider rootStore={rootStore}>
          <Routes>
            <Route path=":name/edit-service" element={<EditServicePage />} />
          </Routes>
        </StoreProvider>
      </MemoryRouter>,
    )

    const editServiceForm = await screen.findByTestId('form')

    await act(async () => {
      fireEvent.submit(editServiceForm)
    })

    expect(screen.queryByRole('alert')).toHaveTextContent(
      "Failed to update the serviceYou don't have the required permission.",
    )
  })
})
