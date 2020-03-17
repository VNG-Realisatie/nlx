// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import React from 'react'
import { renderWithProviders } from '../../../test-utils'
import AuthorizationMode from './index'

test('no authorization mode', () => {
  const { getByText } = renderWithProviders(
    <AuthorizationMode authorizations={[]} mode="none" />,
  )

  expect(getByText('Open')).toBeInTheDocument()
})

test('white list mode', () => {
  const { getByText } = renderWithProviders(
    <AuthorizationMode authorizations={[{}, {}]} mode="whitelist" />,
  )

  expect(getByText('Whitelist (2)')).toBeInTheDocument()
})
