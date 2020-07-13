// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { renderWithProviders } from '../../../../../test-utils'
import AuthorizationMode from './index'

test('no authorization mode', () => {
  const { getByText } = renderWithProviders(
    <AuthorizationMode authorizations={[]} mode="none" />,
  )

  expect(getByText('Open')).toBeInTheDocument()
})

test('white list mode', () => {
  const { getByText, getByTestId } = renderWithProviders(
    <AuthorizationMode authorizations={[{}, {}]} mode="whitelist" />,
  )

  expect(getByText('Whitelist')).toBeInTheDocument()
  expect(getByTestId('authorization-mode-count').textContent).toBe('2')
})

test('authorizations not specified', () => {
  const { container, getByTestId } = renderWithProviders(
    <AuthorizationMode mode="whitelist" />,
  )

  expect(container).toHaveTextContent('Whitelist')
  expect(getByTestId('authorization-mode-count')).toHaveTextContent('0')
})
