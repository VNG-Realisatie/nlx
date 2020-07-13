// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../test-utils'
import InwaysPageView from './index'

jest.mock('./InwayRow', () => () => (
  <tr data-testid="mock-row">
    <td>inway</td>
  </tr>
))

test('no inways', () => {
  const { getByText } = renderWithProviders(<InwaysPageView inways={[]} />)
  expect(
    getByText(/^There are no inways registered yet\.$/),
  ).toBeInTheDocument()
})

test('service list', () => {
  const inways = [{ name: 'inway1' }, { name: 'inway2' }]
  const { getByTestId, getAllByTestId } = renderWithProviders(
    <InwaysPageView inways={inways} />,
  )

  expect(getByTestId('inways-list')).toBeInTheDocument()
  expect(getAllByTestId('mock-row')).toHaveLength(2)
})
