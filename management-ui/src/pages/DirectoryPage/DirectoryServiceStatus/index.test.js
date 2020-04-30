// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import DirectoryServiceStatus from './index'

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(<DirectoryServiceStatus status="up" />),
  ).not.toThrow()
})

test('renders an icon', () => {
  const { container, rerender } = renderWithProviders(
    <DirectoryServiceStatus status="up" />,
  )
  expect(container).toHaveTextContent('status-up.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Up')

  rerender(<DirectoryServiceStatus status="down" />)
  expect(container).toHaveTextContent('status-down.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Down')

  rerender(<DirectoryServiceStatus status="degraded" />)
  expect(container).toHaveTextContent('status-degraded.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Degraded')

  rerender(<DirectoryServiceStatus status="unknown" />)
  expect(container).toHaveTextContent('status-unknown.svg')
  expect(container.querySelector('svg')).toHaveAttribute('title', 'Unknown')
})

test('renders the unknown icon when status is not valid', () => {
  jest.spyOn(console, 'warn').mockImplementation(() => undefined)
  const { container, rerender } = renderWithProviders(
    <DirectoryServiceStatus status={null} />,
  )
  expect(container).toBeEmpty()
  expect(console.warn).toHaveBeenCalledTimes(1)
  console.warn.mockRestore()

  jest.spyOn(console, 'warn').mockImplementation(() => undefined)
  jest.spyOn(console, 'error').mockImplementation(() => undefined)
  rerender(<DirectoryServiceStatus status="invalid" />)
  expect(container).toBeEmpty()

  expect(console.warn).toHaveBeenCalledTimes(1)
  expect(console.error).toHaveBeenCalledTimes(1)
  console.warn.mockRestore()
  console.error.mockRestore()
})
