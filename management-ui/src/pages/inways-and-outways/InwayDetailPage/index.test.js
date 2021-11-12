// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter as Router } from 'react-router-dom'
import { renderWithAllProviders } from '../../../test-utils'
import { RootStore, StoreProvider } from '../../../stores'
import InwayDetailPage from './index'

/* eslint-disable react/prop-types */
jest.mock('./InwayDetailPageView', () => ({ inway }) => (
  <div data-testid="inway-details">{inway.name}</div>
))
/* eslint-enable react/prop-types */

test('display inway details', () => {
  const rootStore = new RootStore({})

  const { getByTestId } = renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <Router location="/inways-and-outways/inways/forty-two">
        <Route path="/inways-and-outways/inways/:name">
          <InwayDetailPage
            inway={{ name: 'forty-two' }}
            parentUrl="/inways-and-outways"
          />
        </Route>
      </Router>
    </StoreProvider>,
  )
  expect(getByTestId('inway-details')).toHaveTextContent('forty-two')
})

test('display a non-existing inway', async () => {
  const rootStore = new RootStore({})

  const { findByTestId } = renderWithAllProviders(
    <StoreProvider rootStore={rootStore}>
      <Router location="/inways-and-outways/inways/forty-two">
        <Route path="/inways-and-outways/inways/:name">
          <InwayDetailPage inway={null} parentUrl="/inways-and-outways" />
        </Route>
      </Router>
    </StoreProvider>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this inway')

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
