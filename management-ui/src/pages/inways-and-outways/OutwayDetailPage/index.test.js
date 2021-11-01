// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Route, StaticRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../../test-utils'
import OutwayModel from '../../../stores/models/OutwayModel'
import OutwayDetailPage from './index'

/* eslint-disable react/prop-types */
jest.mock('./OutwayDetailPageView', () => ({ outway }) => (
  <div data-testid="outway-details">{outway.name}</div>
))
/* eslint-enable react/prop-types */

test('display outway details', () => {
  const outwayModel = new OutwayModel({
    outwayData: {
      name: 'forty-two',
    },
  })

  const { getByTestId } = renderWithProviders(
    <Router location="/inways-and-outways/forty-two">
      <Route path="/inways-and-outways/:name">
        <OutwayDetailPage outway={outwayModel} />
      </Route>
    </Router>,
  )

  expect(getByTestId('outway-details')).toHaveTextContent('forty-two')
})

test('display a non-existing outway', async () => {
  const { findByTestId } = renderWithProviders(
    <Router location="/inways-and-outways/forty-two">
      <Route path="/inways-and-outways/:name">
        <OutwayDetailPage outway={null} />
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this outway')

  const closeButton = await findByTestId('close-button')
  expect(closeButton).toBeTruthy()
})
