// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { StaticRouter as Router, Route } from 'react-router-dom'

import { renderWithProviders } from '../../../test-utils'
import InwayDetailPage from './index'

jest.mock('./InwayDetailPageView', () => ({ inway }) => (
  <div data-testid="inway-details">{inway.name}</div>
))

test('display inway details', async () => {
  const getInwayByName = jest.fn().mockResolvedValue({ name: 'forty-two' })

  const { findByTestId } = renderWithProviders(
    <Router location="/inways/forty-two">
      <Route path="/inways/:name">
        <InwayDetailPage getInwayByName={getInwayByName} />
      </Route>
    </Router>,
  )

  expect(await findByTestId('inway-details')).toHaveTextContent('forty-two')
  expect(getInwayByName).toHaveBeenCalledWith('forty-two')
})

test('fetching a non-existing component', async () => {
  const getInwayByName = jest.fn().mockRejectedValue(new Error('not found'))

  const { findByTestId } = renderWithProviders(
    <Router location="/inways/forty-two">
      <Route path="/inways/:name">
        <InwayDetailPage getInwayByName={getInwayByName} />
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this inway.')
})

test('fetching inway details fails for an unknown reason', async () => {
  const getInwayByName = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary reason'))

  const { findByTestId } = renderWithProviders(
    <Router location="/inways/42">
      <Route path="/inways/:name">
        <InwayDetailPage getInwayByName={getInwayByName} />
      </Route>
    </Router>,
  )

  const message = await findByTestId('error-message')
  expect(message).toBeTruthy()
  expect(message.textContent).toBe('Failed to load the details for this inway.')
})
