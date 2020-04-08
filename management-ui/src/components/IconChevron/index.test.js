// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { render } from '@testing-library/react'

import IconChevron from './index'

test('renders without crashing', () => {
  expect(() => {
    render(<IconChevron />)
  }).not.toThrow()
})
