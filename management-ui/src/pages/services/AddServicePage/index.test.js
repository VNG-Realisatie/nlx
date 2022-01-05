// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, fireEvent } from '@testing-library/react'
import {
  MemoryRouter,
  Routes,
  Route,
  unstable_HistoryRouter as HistoryRouter,
} from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import AddServicePage from './index'

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
  <form onSubmit={() => onSubmitHandler({ foo: 'bar' })} data-testid="form">
    <button type="submit" />
  </form>
))

describe('the AddServicePage', () => {
  beforeAll(() => {
    global.scrollTo = jest.fn()
  })
  afterEach(() => {
    jest.resetModules()
  })

  it('on initialization', () => {
    const managementApiClient = new ManagementApi()
    const store = new RootStore({ managementApiClient })
    const { getByTestId, queryByRole, getByLabelText } = renderWithProviders(
      <MemoryRouter>
        <StoreProvider rootStore={store}>
          <AddServicePage />
        </StoreProvider>
      </MemoryRouter>,
    )

    const linkBack = getByLabelText(/Back/)
    expect(linkBack.getAttribute('href')).toBe('/services')
    expect(getByTestId('form')).toBeTruthy()
    expect(queryByRole('dialog')).toBeNull()
  })

  it('successfully submitting the form', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementCreateService = jest.fn().mockResolvedValue({
      name: 'my-service',
    })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const history = createMemoryHistory()
    const { findByTestId } = renderWithProviders(
      <HistoryRouter history={history}>
        <StoreProvider rootStore={rootStore}>
          <Routes>
            <Route path="*" element={<AddServicePage />} />
          </Routes>
        </StoreProvider>
      </HistoryRouter>,
    )

    const addComponentForm = await findByTestId('form')
    await act(async () => {
      fireEvent.submit(addComponentForm)
    })

    expect(managementApiClient.managementCreateService).toHaveBeenCalledTimes(1)
    expect(history.location.pathname).toEqual('/services/my-service')
    expect(history.location.search).toEqual('?lastAction=added')
  })

  it('re-submitting the form when the previous submission went wrong', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementCreateService = jest
      .fn()
      .mockResolvedValue({ name: 'my-service' })
      .mockRejectedValueOnce({
        json: () => {
          return { message: 'arbitrary error' }
        },
      })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const history = createMemoryHistory()
    const { findByTestId, queryByRole } = renderWithProviders(
      <HistoryRouter history={history}>
        <StoreProvider rootStore={rootStore}>
          <Routes>
            <Route path="*" element={<AddServicePage />} />
          </Routes>
        </StoreProvider>
      </HistoryRouter>,
    )

    const addComponentForm = await findByTestId('form')

    await act(async () => {
      fireEvent.submit(addComponentForm)
    })

    expect(managementApiClient.managementCreateService).toHaveBeenCalledTimes(1)
    expect(queryByRole('alert').textContent).toBe(
      'Failed adding servicearbitrary error',
    )

    await act(async () => {
      fireEvent.submit(addComponentForm)
    })

    expect(queryByRole('alert')).toBeTruthy()

    expect(managementApiClient.managementCreateService).toHaveBeenCalledTimes(2)
    expect(history.location.pathname).toEqual('/services/my-service')
    expect(history.location.search).toEqual('?lastAction=added')
  })

  it('submitting when the HTTP response is not ok', async () => {
    const managementApiClient = new ManagementApi()
    managementApiClient.managementCreateService = jest
      .fn()
      .mockRejectedValueOnce({
        json: () => {
          return { message: 'arbitrary error' }
        },
      })

    const rootStore = new RootStore({
      managementApiClient,
    })

    const { findByTestId, queryByRole } = renderWithProviders(
      <MemoryRouter>
        <StoreProvider rootStore={rootStore}>
          <AddServicePage />
        </StoreProvider>
      </MemoryRouter>,
    )

    const addComponentForm = await findByTestId('form')

    await act(async () => {
      fireEvent.submit(addComponentForm)
    })

    expect(queryByRole('alert').textContent).toBe(
      'Failed adding servicearbitrary error',
    )
  })
})
