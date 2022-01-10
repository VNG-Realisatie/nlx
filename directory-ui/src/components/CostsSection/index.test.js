// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { fireEvent, renderWithProviders } from '../../test-utils'
import CostsSection from './index'

test('Costs section', async () => {
  const { container, rerender, getByText } = renderWithProviders(
    <MemoryRouter>
      <CostsSection oneTimeCosts={0} monthlyCosts={0} requestCosts={0} />
    </MemoryRouter>,
  )

  expect(container).toHaveTextContent(/Geen/)

  rerender(
    <MemoryRouter>
      <CostsSection />
    </MemoryRouter>,
  )
  expect(container).toHaveTextContent(/Geen/)

  rerender(
    <MemoryRouter>
      <CostsSection oneTimeCosts={5} monthlyCosts={0} requestCosts={10} />
    </MemoryRouter>,
  )

  fireEvent.click(getByText(/kosten/i))

  expect(container).toHaveTextContent(/eenmalige en per aanvraag/)

  expect(container).toHaveTextContent(/Eenmalige kosten€ 5,00/)
  expect(container).toHaveTextContent(/Kosten per aanvraag€ 10,00/)

  rerender(
    <MemoryRouter>
      <CostsSection oneTimeCosts={5} monthlyCosts={10} requestCosts={15} />
    </MemoryRouter>,
  )

  expect(container).toHaveTextContent(/eenmalige, maandelijkse en per aanvraag/)
})
