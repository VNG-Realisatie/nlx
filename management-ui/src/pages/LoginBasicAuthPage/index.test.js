// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { act, fireEvent } from '@testing-library/react'
import { createMemoryHistory } from 'history'
import { Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import { RootStore, StoreProvider } from '../../stores'
import LoginBasicAuthPage from './index'

test('user is not authenticated', async () => {
  const fetchUser = () => {
    throw new Error('not authenticated')
  }

  const { findByText, getByLabelText } = renderWithProviders(
    <UserContextProvider fetchAuthenticatedUser={fetchUser}>
      <LoginBasicAuthPage />
    </UserContextProvider>,
  )

  expect(await findByText(/^Welcome$/)).toBeInTheDocument()
  expect(getByLabelText('Email address')).toBeInTheDocument()
  expect(getByLabelText('Password')).toBeInTheDocument()
})

test('user is authenticated', async () => {
  const history = createMemoryHistory({ initialEntries: ['/login'] })
  const logout = jest.fn()

  const { getByText } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <LoginBasicAuthPage logout={logout} />
      </UserContextProvider>
    </Router>,
  )

  const logOutButton = getByText('Log out')

  await act(async () => {
    fireEvent.click(logOutButton)
  })

  expect(logout).toHaveBeenCalledTimes(1)
  expect(window.location.pathname).toEqual('/')
})

test('when authentication fails', async () => {
  const fetchUser = () => {
    throw new Error('not authenticated')
  }

  const loginHandler = jest.fn().mockRejectedValue(new Error('arbitrary error'))

  const rootStore = new RootStore({})

  const { getByLabelText, getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <UserContextProvider fetchAuthenticatedUser={fetchUser}>
        <LoginBasicAuthPage loginHandler={loginHandler} />
      </UserContextProvider>
    </StoreProvider>,
  )

  const emailInput = getByLabelText('Email address')
  const passwordInput = getByLabelText('Password')

  await act(async () => {
    fireEvent.change(emailInput, { target: { value: 'naam@e-mailadres.nl' } })
    fireEvent.change(passwordInput, { target: { value: '1234' } })
  })

  const submitButton = getByText('Log in')

  await act(async () => {
    fireEvent.click(submitButton)
  })

  expect(getByText('Invalid credentials.')).toBeInTheDocument()
})

test('when authentication succeeds', async () => {
  const fetchUser = () => {
    throw new Error('not authenticated')
  }

  const loginHandler = jest.fn().mockResolvedValue(true)
  const storeCredentials = jest.fn()

  const history = createMemoryHistory({ initialEntries: ['/login'] })

  const rootStore = new RootStore({})

  const { getByLabelText, getByText } = renderWithProviders(
    <Router history={history}>
      <StoreProvider rootStore={rootStore}>
        <UserContextProvider fetchAuthenticatedUser={fetchUser}>
          <LoginBasicAuthPage
            loginHandler={loginHandler}
            storeCredentials={storeCredentials}
          />
        </UserContextProvider>
      </StoreProvider>
    </Router>,
  )

  const emailInput = getByLabelText('Email address')
  const passwordInput = getByLabelText('Password')

  await act(async () => {
    fireEvent.change(emailInput, { target: { value: 'naam@e-mailadres.nl' } })
    fireEvent.change(passwordInput, { target: { value: '1234' } })
  })

  const submitButton = getByText('Log in')

  await act(async () => {
    fireEvent.click(submitButton)
  })

  expect(storeCredentials).toHaveBeenCalledWith('naam@e-mailadres.nl', '1234')
  expect(window.location.pathname).toEqual('/')
})
