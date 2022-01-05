// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter } from 'react-router-dom'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import LoginOIDCPage from './index'

test('renders a welcome message', async () => {
  const { findByText } = renderWithProviders(
    <MemoryRouter>
      <UserContextProvider>
        <LoginOIDCPage />
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(await findByText(/^Welcome$/)).toBeInTheDocument()
})

test('when authentication fails', async () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/login#auth-fail']}>
      <UserContextProvider>
        <LoginOIDCPage />
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(await screen.findByTestId('login-error-message')).toHaveTextContent(
    'Something went wrong when logging in. Please try again.',
  )
})

test('when the authenticating user is missing', async () => {
  renderWithProviders(
    <MemoryRouter initialEntries={['/login#auth-missing-user']}>
      <UserContextProvider>
        <LoginOIDCPage />
      </UserContextProvider>
    </MemoryRouter>,
  )

  expect(await screen.findByTestId('login-error-message')).toHaveTextContent(
    "User doesn't exist in NLX Management.",
  )
})
