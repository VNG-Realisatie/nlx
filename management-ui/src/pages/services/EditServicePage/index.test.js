// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, fireEvent } from '@testing-library/react'
import { Route, Router, StaticRouter } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import UserContext from '../../../user-context'
import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import EditServicePage from './index'

jest.mock('../../../components/PageTemplate/OrganizationInwayCheck', () => () =>
  null,
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
  afterEach(() => {
    jest.resetModules()
  })

  it('before the service has been loaded', async () => {
    jest.useFakeTimers()

    const rootStore = new RootStore({
      managementApiClient: new ManagementApi(),
    })

    const userContext = { user: { id: '42' } }
    const { findByRole, getByLabelText } = renderWithProviders(
      <StaticRouter location="/services/mock-service/edit-service">
        <Route path="/services/:name/edit-service">
          <UserContext.Provider value={userContext}>
            <StoreProvider rootStore={rootStore}>
              <EditServicePage />
            </StoreProvider>
          </UserContext.Provider>
        </Route>
      </StaticRouter>,
    )

    expect(await findByRole('progressbar')).toBeTruthy()
    const linkBack = getByLabelText(/Back/)
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
    const { findByRole, queryByRole } = renderWithProviders(
      <StaticRouter>
        <UserContext.Provider value={userContext}>
          <StoreProvider rootStore={rootStore}>
            <EditServicePage />
          </StoreProvider>
        </UserContext.Provider>
      </StaticRouter>,
    )

    expect(await findByRole('alert')).toBeTruthy()
    expect(queryByRole('alert').textContent).toBe('Failed to load the service')
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
    const { findByTestId } = renderWithProviders(
      <StaticRouter location="/services/mock-service/edit-service">
        <Route path="/services/:name/edit-service">
          <UserContext.Provider value={userContext}>
            <StoreProvider rootStore={rootStore}>
              <EditServicePage />
            </StoreProvider>
          </UserContext.Provider>
        </Route>
      </StaticRouter>,
    )

    expect(await findByTestId('form')).toBeTruthy()
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
      initialEntries: ['/services/mock-service/edit-service'],
    })

    const { findByTestId } = renderWithProviders(
      <Router history={history}>
        <StoreProvider rootStore={rootStore}>
          <Route path="/services/:name/edit-service">
            <EditServicePage />
          </Route>
        </StoreProvider>
      </Router>,
    )

    const editServiceForm = await findByTestId('form')
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
      initialEntries: ['/services/mock-service/edit-service'],
    })

    const { findByTestId, queryByRole } = renderWithProviders(
      <Router history={history}>
        <StoreProvider rootStore={rootStore}>
          <Route path="/services/:name/edit-service">
            <EditServicePage />
          </Route>
        </StoreProvider>
      </Router>,
    )

    const editServiceForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(editServiceForm)
    })

    expect(managementApiClient.managementUpdateService).toHaveBeenCalledTimes(1)
    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert')).toHaveTextContent(
      'Failed to update the service',
    )

    await act(async () => {
      await fireEvent.submit(editServiceForm)
    })

    expect(managementApiClient.managementUpdateService).toHaveBeenCalledTimes(2)

    expect(history.location.pathname).toEqual('/services/mock-service')
    expect(history.location.search).toEqual('?lastAction=edited')
  })

  it('submitting when the HTTP response is not ok', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementUpdateService = jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error'))

    managementApiClient.managementListServices = jest.fn().mockResolvedValue({
      services: [{ name: 'mock-service' }],
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    await rootStore.servicesStore.fetchAll()

    const history = createMemoryHistory({
      initialEntries: ['/services/mock-service/edit-service'],
    })

    const { findByTestId, queryByRole } = renderWithProviders(
      <Router history={history}>
        <StoreProvider rootStore={rootStore}>
          <Route path="/services/:name/edit-service">
            <EditServicePage />
          </Route>
        </StoreProvider>
      </Router>,
    )

    const editServiceForm = await findByTestId('form')

    await act(async () => {
      await fireEvent.submit(editServiceForm)
    })

    expect(queryByRole('alert')).toBeTruthy()
    expect(queryByRole('alert')).toHaveTextContent(
      'Failed to update the service',
    )
  })
})
