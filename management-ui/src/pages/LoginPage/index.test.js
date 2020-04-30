// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders, waitFor } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import LoginPage from './index'

test('renders a welcome message', async () => {
  const { getByText } = renderWithProviders(
    <UserContextProvider>
      <LoginPage />
    </UserContextProvider>,
  )
  waitFor(() => expect(getByText(/^Welcome$/)).toBeInTheDocument())
})
