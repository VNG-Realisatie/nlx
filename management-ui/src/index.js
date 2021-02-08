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

ReactDOM.render(
  <Router>
    <StoreProvider rootStore={rootStore}>
      <UserContextProvider>
        <App>
          <Routes />
        </App>
      </UserContextProvider>
    </StoreProvider>
  </Router>,
  document.getElementById('root'),
)
