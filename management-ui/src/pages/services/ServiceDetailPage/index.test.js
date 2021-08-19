// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, Router, StaticRouter } from 'react-router-dom'
import { act, fireEvent } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { configure } from 'mobx'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import { ManagementApi } from '../../../api'
import ServiceDetailPage from './index'

// eslint-disable-next-line react/prop-types
jest.mock('./ServiceDetailView', () => ({ removeHandler }) => (
  <div data-testid="service-details">
    <button type="button" onClick={removeHandler}>
      Remove service
    </button>
  </div>
))

let fetchIncomingAccessRequests
let fetchAccessGrants

beforeEach(() => {
  fetchIncomingAccessRequests = jest.fn()
  fetchAccessGrants = jest.fn()
})

test('display service details', () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })
  const { getByTestId, getByText } = renderWithAllProviders(
    <StaticRouter location="/services/forty-two">
      <Route path="/services/:name">
        <StoreProvider rootStore={rootStore}>
          <ServiceDetailPage
            service={{
              name: 'forty-two',
              fetchIncomingAccessRequests,
              fetchAccessGrants,
            }}
          />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )

  expect(getByTestId('service-details')).toBeInTheDocument()
  expect(getByText('forty-two')).toBeInTheDocument()
})

test('fetching a non-existing component', async () => {
  const managementApiClient = new ManagementApi()
  const rootStore = new RootStore({ managementApiClient })

  const { findByTestId, getByText } = renderWithAllProviders(
    <StaticRouter location="/services/forty-two">
      <Route path="/services/:name">
        <StoreProvider rootStore={rootStore}>
          <ServiceDetailPage />
        </StoreProvider>
      </Route>
    </StaticRouter>,
  )
  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the service')

  expect(getByText('forty-two')).toBeInTheDocument()

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})

test('removing the service', async () => {
  configure({ safeDescriptors: false })
  const managementApiClient = new ManagementApi()
  managementApiClient.managementDeleteService = jest.fn().mockResolvedValue()

  const rootStore = new RootStore({
    managementApiClient,
  })
  jest.spyOn(rootStore.servicesStore, 'removeService')

  const history = createMemoryHistory({
    initialEntries: ['/services/dummy-service'],
  })

  const { findByText } = renderWithAllProviders(
    <Router history={history}>
      <Route path="/services/:name">
        <StoreProvider rootStore={rootStore}>
          <ServiceDetailPage
            service={{
              name: 'dummy-service',
              fetchIncomingAccessRequests,
              fetchAccessGrants,
            }}
          />
        </StoreProvider>
      </Route>
    </Router>,
  )

  const removeButton = await findByText('Remove service')
  fireEvent.click(removeButton)

  expect(rootStore.servicesStore.removeService).toHaveBeenCalledTimes(1)
  await act(async () => {})
  expect(history.location.pathname).toEqual('/services/dummy-service')
  expect(history.location.search).toEqual('?lastAction=removed')
})
