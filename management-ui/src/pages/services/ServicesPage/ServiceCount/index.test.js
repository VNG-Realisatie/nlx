// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import ServiceCount from './index'

test('renders without crashing', () => {
  expect(() => renderWithProviders(<ServiceCount count={1} />)).not.toThrow()
})
