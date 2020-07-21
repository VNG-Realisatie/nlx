// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import StateIndicator from './index'

test('renders without crashing', () => {
  expect(() => renderWithProviders(<StateIndicator state="up" />)).not.toThrow()
})

test('renders an icon', () => {
  const { container, rerender } = renderWithProviders(
    <StateIndicator state="up" />,
  )
  expect(container).toHaveTextContent('state-up.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Up')

  rerender(<StateIndicator state="down" />)
  expect(container).toHaveTextContent('state-down.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Down')

  rerender(<StateIndicator state="degraded" />)
  expect(container).toHaveTextContent('state-degraded.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Degraded')

  rerender(<StateIndicator state="unknown" />)
  expect(container).toHaveTextContent('state-unknown.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Unknown')
})

test('does not render when state is invalid', () => {
  // Suppress console output during test
  global.console.warn = jest.fn()
  global.console.error = jest.fn()

  const { container, rerender } = renderWithProviders(
    <StateIndicator state={null} />,
  )
  expect(container).toBeEmptyDOMElement()

  rerender(<StateIndicator state="invalid" />)
  expect(container).toBeEmptyDOMElement()
})

describe('state text', () => {
  it('is hidden by default', () => {
    const { queryByText } = renderWithProviders(<StateIndicator state="up" />)
    expect(queryByText('Up')).toBeNull()
  })

  it('shown with bool prop `showText`', () => {
    const { getByText } = renderWithProviders(
      <StateIndicator state="up" showText />,
    )
    expect(getByText('Up')).toBeInTheDocument()
  })
})
