// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../test-utils'
import Switch from './index'

test('renders without crashing', () => {
  expect(() => renderWithProviders(<Switch />)).not.toThrow()
})
