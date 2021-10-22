// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import InwayDetailPage from './index'

/* eslint-disable react/prop-types */
jest.mock('./InwayDetailPageView', () => ({ inway }) => (
  <div data-testid="inway-details">{inway.name}</div>
))
/* eslint-enable react/prop-types */

test('display inway details', () => {
  const { getByTestId } = renderWithProviders(
    <Router location="/inways/forty-two">
      <Route path="/inways/:name">
        <InwayDetailPage inway={{ name: 'forty-two' }} />
      </Route>
    </Router>,
  )

  expect(getByTestId('inway-details')).toHaveTextContent('forty-two')
})

test('display a non-existing inway', async () => {
  const { findByTestId } = renderWithProviders(
    <Router location="/inways/forty-two">
      <Route path="/inways/:name">
        <InwayDetailPage inway={null} />
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this inway')

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
