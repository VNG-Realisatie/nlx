// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../../test-utils'
import StatusIcon from './index'

test('renders without crashing', () => {
  expect(() => renderWithProviders(<StatusIcon state="up" />)).not.toThrow()
})

test('renders an icon', () => {
  const { container, rerender } = renderWithProviders(<StatusIcon active />)
  expect(container).toHaveTextContent('state-up.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Active')

  rerender(<StatusIcon />)
  expect(container).toHaveTextContent('state-down.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Inactive')
})
