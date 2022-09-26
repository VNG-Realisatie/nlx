// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import StateIndicator from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(<StateIndicator state="STATE_UP" />),
  ).not.toThrow()
})

test('renders an icon', () => {
  const { container, rerender } = renderWithProviders(
    <StateIndicator state="STATE_UP" />,
  )
  expect(container).toHaveTextContent('state-up.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Up')

  rerender(<StateIndicator state="STATE_DOWN" />)
  expect(container).toHaveTextContent('state-down.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Down')

  rerender(<StateIndicator state="STATE_DEGRADED" />)
  expect(container).toHaveTextContent('state-degraded.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Degraded')

  rerender(<StateIndicator state="STATE_UNSPECIFIED" />)
  expect(container).toHaveTextContent('state-unknown.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Unknown')
})

test('does not render when state is invalid', () => {
  global.console.error = jest.fn()

  const { container, rerender } = renderWithProviders(
    <StateIndicator state={null} />,
  )

  expect(container).toHaveTextContent('state-unknown.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Unknown')

  rerender(<StateIndicator state="STATE_UNKNOWN" />)

  expect(container).toHaveTextContent('state-unknown.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Unknown')
})

describe('state text', () => {
  it('is hidden by default', () => {
    const { queryByText } = renderWithProviders(
      <StateIndicator state="STATE_UP" />,
    )
    expect(queryByText('Up')).toBeNull()
  })

  it('shown with bool prop `showText`', () => {
    const { getByText } = renderWithProviders(
      <StateIndicator state="STATE_UP" showText />,
    )
    expect(getByText('Up')).toBeInTheDocument()
  })
})
