// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../test-utils'
import ServicesPageView from './index'

jest.mock('./ServiceRow', () => () => (
  <tr data-testid="mockRow">
    <td>service</td>
  </tr>
))

test('no services', () => {
  const { getByText } = renderWithProviders(<ServicesPageView services={[]} />)
  expect(getByText(/^There are no services yet\.$/)).toBeInTheDocument()
})

test('service list', () => {
  const services = [{ name: 'service 1' }, { name: 'service 2' }]
  const { getByTestId, getAllByTestId } = renderWithProviders(
    <ServicesPageView services={services} />,
  )

  expect(getByTestId('services-list')).toBeInTheDocument()
  expect(getAllByTestId('mockRow')).toHaveLength(2)
})
