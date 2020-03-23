import React from 'react'
import { render } from '@testing-library/react'
import IconServices from './IconServices'

test('renders without crashing', () => {
  expect(() => {
    render(<IconServices />)
  }).not.toThrow()
})
