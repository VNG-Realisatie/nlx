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
import { Configuration, DirectoryApi } from './api'

const directoryApiServiceConfig = new Configuration({
  basePath: '',
})
const directoryApiService = new DirectoryApi(directoryApiServiceConfig)
const rootStore = new RootStore({
  directoryApiService,
})

ReactDOM.render(
  <Router>
    <StoreProvider rootStore={rootStore}>
      <UserContextProvider>
        <App />
      </UserContextProvider>
    </StoreProvider>
  </Router>,
  document.getElementById('root'),
)
