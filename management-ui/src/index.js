// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import 'react-app-polyfill/stable'

import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter as Router } from 'react-router-dom'
import './i18n'
import { RootStore, StoreProvider } from './stores'
import { UserContextProvider } from './user-context'
import App from './App'
import { Configuration, DirectoryApi, ManagementApi } from './api'
import Routes from './routes'
import UserRepositoryOIDC from './domain/user-repository-oidc'
import UserRepositoryBasicAuth from './domain/user-repository-basic-auth'
import LoginBasicAuthPage from './pages/LoginBasicAuthPage'
import LoginOIDCPage from './pages/LoginOIDCPage'
import { AUTH_BASIC_AUTH, AUTH_OIDC } from './stores/ApplicationStore'

const directoryApiClient = new DirectoryApi(
  new Configuration({
    basePath: '',
  }),
)

const managementApiClient = new ManagementApi(
  new Configuration({
    basePath: '',
  }),
)

const rootStore = new RootStore({
  directoryApiClient,
  managementApiClient,
})

const getAuthStrategy = async () => {
  const defaultStrategy = AUTH_OIDC

  try {
    const response = await fetch('/basic-auth/')
    return response.status === 204 ? AUTH_BASIC_AUTH : AUTH_OIDC
  } catch (error) {
    console.warn(
      `failed to determine auth strategy. falling back to ${defaultStrategy}`,
      error.message,
    )
    return defaultStrategy
  }
}

;(async () => {
  const authStrategy = await getAuthStrategy()
  const isBasicAuth = authStrategy === AUTH_BASIC_AUTH

  if (isBasicAuth) {
    rootStore.applicationStore.setBasicAuthStrategy()

    const credentials = UserRepositoryBasicAuth.getCredentials()
    if (credentials) {
      managementApiClient.configuration.configuration.headers = {
        Authorization: `Basic ${credentials}`,
      }
      directoryApiClient.configuration.configuration.headers = {
        Authorization: `Basic ${credentials}`,
      }
    }
  }

  const fetchUser = isBasicAuth
    ? UserRepositoryBasicAuth.getAuthenticatedUser
    : UserRepositoryOIDC.getAuthenticatedUser

  const logout = isBasicAuth
    ? UserRepositoryBasicAuth.logout
    : UserRepositoryOIDC.logout

  const authPage = isBasicAuth ? LoginBasicAuthPage : LoginOIDCPage

  ReactDOM.render(
    <Router>
      <StoreProvider rootStore={rootStore}>
        <UserContextProvider fetchAuthenticatedUser={fetchUser} logout={logout}>
          <App>
            <Routes authorizationPage={authPage} />
          </App>
        </UserContextProvider>
      </StoreProvider>
    </Router>,
    document.getElementById('root'),
  )
})()
