// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import StatusIndicator from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(<StatusIndicator status="up" />),
  ).not.toThrow()
})

test('renders an icon', () => {
  const { container, rerender } = renderWithProviders(
    <StatusIndicator status="up" />,
  )
  expect(container).toHaveTextContent('status-up.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Up')

  rerender(<StatusIndicator status="down" />)
  expect(container).toHaveTextContent('status-down.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Down')

  rerender(<StatusIndicator status="degraded" />)
  expect(container).toHaveTextContent('status-degraded.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Degraded')

  rerender(<StatusIndicator status="unknown" />)
  expect(container).toHaveTextContent('status-unknown.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Unknown')
})

test('does not render when status is invalid', () => {
  // Suppress console output during test
  global.console.warn = jest.fn()
  global.console.error = jest.fn()

  const { container, rerender } = renderWithProviders(
    <StatusIndicator status={null} />,
  )
  expect(container).toBeEmptyDOMElement()

  rerender(<StatusIndicator status="invalid" />)
  expect(container).toBeEmptyDOMElement()
})

describe('status text', () => {
  it('is hidden by default', () => {
    const { queryByText } = renderWithProviders(<StatusIndicator status="up" />)
    expect(queryByText('Up')).toBeNull()
  })

  it('shown with bool prop `showText`', () => {
    const { getByText } = renderWithProviders(
      <StatusIndicator status="up" showText />,
    )
    expect(getByText('Up')).toBeInTheDocument()
  })
})
