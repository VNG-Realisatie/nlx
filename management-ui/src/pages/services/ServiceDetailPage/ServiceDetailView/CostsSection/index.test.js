// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { fireEvent, renderWithProviders } from '../../../../../test-utils'
import CostsSection from './index'

test('Costs section', async () => {
  const { container, rerender, getByText } = renderWithProviders(
    <CostsSection oneTimeCosts={0} monthlyCosts={0} requestCosts={0} />,
  )

  expect(container).toHaveTextContent(/Free/)

  rerender(<CostsSection oneTimeCosts={5} monthlyCosts={0} requestCosts={10} />)

  fireEvent.click(getByText(/Costs/i))

  expect(container).toHaveTextContent(/one time and per request/)

  expect(container).toHaveTextContent(/One time costs€ 5,00/)
  expect(container).toHaveTextContent(/Cost per request€ 10,00/)

  rerender(
    <CostsSection oneTimeCosts={5} monthlyCosts={10} requestCosts={15} />,
  )

  expect(container).toHaveTextContent(/one time, monthly and per request/)
})
