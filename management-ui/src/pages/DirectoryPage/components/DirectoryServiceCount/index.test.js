// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import DirectoryServiceCount from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(<DirectoryServiceCount services={[{}]} />),
  ).not.toThrow()
})
