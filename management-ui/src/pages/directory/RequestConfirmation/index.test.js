// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import RequestConfirmation from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <RequestConfirmation
        organizationName="organization"
        serviceName="service"
      />,
    ),
  ).not.toThrow()
})
