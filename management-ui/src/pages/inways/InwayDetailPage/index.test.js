// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter as Router, Route } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import { mockServicesStore } from '../../../stores/ServicesStore.mock'
import { mockInwaysStore } from '../../../stores/InwaysStore.mock'
import { StoreProvider } from '../../../stores'
import InwayDetailPage from './index'

/* eslint-disable react/prop-types */
jest.mock('./InwayDetailPageView', () => ({ inway }) => (
  <div data-testid="inway-details">{inway.name}</div>
))
/* eslint-enable react/prop-types */

test('display inway details', () => {
  const store = mockInwaysStore({})
  const { getByTestId } = renderWithProviders(
    <Router location="/inways/forty-two">
      <Route path="/inways/:name">
        <StoreProvider store={store}>
          <InwayDetailPage inway={{ name: 'forty-two' }} />
        </StoreProvider>
      </Route>
    </Router>,
  )

  expect(getByTestId('inway-details')).toHaveTextContent('forty-two')
})

test('fetching a non-existing component', async () => {
  const selectInway = jest.fn()
  const store = mockInwaysStore({ selectInway })

  const { findByTestId } = renderWithProviders(
    <Router location="/inways/forty-two">
      <Route path="/inways/:name">
        <StoreProvider store={store}>
          <InwayDetailPage />
        </StoreProvider>
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this inway.')

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})

test('fetching inway details fails for an unknown reason', async () => {
  const store = mockServicesStore({ error: 'arbitrary reason' })

  const { findByTestId } = renderWithProviders(
    <Router location="/inways/42">
      <Route path="/inways/:name">
        <StoreProvider store={store}>
          <InwayDetailPage />
        </StoreProvider>
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this inway.')

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
