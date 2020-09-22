// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'

import { renderWithProviders, fireEvent } from '../../../../../test-utils'
import InwaysSection from './index'

beforeEach(() => {
  jest.useFakeTimers()
})

test('should show inways', async () => {
  const { getByTestId, queryByTestId } = renderWithProviders(
    <Router>
      <InwaysSection inways={['inway1', 'inway2']} />
    </Router>,
  )

  expect(getByTestId('service-inways')).toHaveTextContent(
    'inway.svg' + 'Inways' + '2', // eslint-disable-line no-useless-concat
  )
  expect(queryByTestId('service-inways-list')).toBeNull()

  fireEvent.click(getByTestId('service-inways'))
  jest.runAllTimers()

  expect(getByTestId('service-inways-list')).toBeTruthy()
  expect(getByTestId('service-inway-1')).toHaveTextContent('inway2')
})

test('alternate rendering when not having inways', () => {
  const { getByTestId } = renderWithProviders(
    <Router>
      <InwaysSection inways={[]} />
    </Router>,
  )

  expect(getByTestId('service-inways')).toHaveTextContent(
    'inway.svg' + 'Inways' + '0', // eslint-disable-line no-useless-concat
  )

  fireEvent.click(getByTestId('service-inways'))
  jest.runAllTimers()

  expect(getByTestId('service-no-inways')).toBeTruthy()
})
