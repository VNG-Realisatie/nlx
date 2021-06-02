// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { render } from '@testing-library/react'
import DocumentationPage from './index'

test('renders without crashing', () => {
  expect(() =>
    render(
      <DocumentationPage
        match={{
          params: {
            serviceName: 'test-service',
            organizationName: 'test-organization',
          },
        }}
      />,
    ),
  ).not.toThrow()
})
