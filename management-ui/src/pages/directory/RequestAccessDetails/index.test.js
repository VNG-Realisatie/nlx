// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import RequestAccessDetails from './index'

test('Request Access Details', () => {
  const { container, rerender } = renderWithProviders(
    <RequestAccessDetails
      organizationName="organization"
      serviceName="service"
    />,
  )

  expect(container).not.toHaveTextContent(/One time costs/)
  expect(container).not.toHaveTextContent(/Monthly costs/)
  expect(container).not.toHaveTextContent(/Cost per request/)

  rerender(
    <RequestAccessDetails
      organizationName="organization"
      serviceName="service"
      oneTimeCosts={5}
      monthlyCosts={10}
      requestCosts={250}
    />,
  )

  expect(container).toHaveTextContent(/One time costs€ 5,00/)
  expect(container).toHaveTextContent(/Monthly costs€ 10,00/)
  expect(container).toHaveTextContent(/Cost per request€ 250,00/)
})
