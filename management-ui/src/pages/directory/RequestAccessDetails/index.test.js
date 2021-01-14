// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import RequestAccessDetails from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <RequestAccessDetails
        organizationName="organization"
        serviceName="service"
      />,
    ),
  ).not.toThrow()
})
