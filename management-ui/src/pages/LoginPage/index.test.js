// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import LoginPage from './index'

test('renders a welcome message', async () => {
  const history = createMemoryHistory()
  const { findByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider>
        <LoginPage />
      </UserContextProvider>
    </Router>,
  )

  expect(await findByText(/^Welcome$/)).toBeInTheDocument()
})

test('when authentication fails', async () => {
  const history = createMemoryHistory({
    initialEntries: ['/login#auth-fail'],
  })
  const { findByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider>
        <LoginPage />
      </UserContextProvider>
    </Router>,
  )

  expect(await findByTestId('login-error-message')).toHaveTextContent(
    'Something went wrong when logging in. Please try again.',
  )
})

test('when the authenticating user is missing', async () => {
  const history = createMemoryHistory({
    initialEntries: ['/login#auth-missing-user'],
  })
  const { findByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider>
        <LoginPage />
      </UserContextProvider>
    </Router>,
  )

  expect(await findByTestId('login-error-message')).toHaveTextContent(
    "User doesn't exist in NLX Management.",
  )
})
