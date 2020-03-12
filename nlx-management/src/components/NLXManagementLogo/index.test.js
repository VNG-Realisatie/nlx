import React from 'react'
import { render } from '@testing-library/react'

import NLXManagementLogo from './index'

test('renders without crashing', () => {
  expect(() => {
    render(<NLXManagementLogo />)
  }).not.toThrow()
})
