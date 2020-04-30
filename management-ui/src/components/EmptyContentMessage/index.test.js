// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../test-utils'
import EmptyContentMessage from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <EmptyContentMessage>nothing to see here</EmptyContentMessage>,
    ),
  ).not.toThrow()
})
