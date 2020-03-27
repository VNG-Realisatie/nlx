// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import LoginPage from './index'

test('renders a welcome message', () => {
  const { getByText } = renderWithProviders(
    <UserContextProvider>
      <LoginPage />
    </UserContextProvider>,
  )
  expect(getByText(/^Welcome$/)).toBeInTheDocument()
})
