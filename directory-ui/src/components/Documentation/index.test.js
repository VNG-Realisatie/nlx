// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import Documentation from './index'

test('renders without crashing', () => {
  expect(() =>
    render(
      <Documentation
        serviceName="service"
        organizationSerialNumber="00000000000000000001"
      />,
    ),
  ).not.toThrow()
})
