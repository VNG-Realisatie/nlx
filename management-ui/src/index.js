// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import 'react-app-polyfill/ie11'
import 'react-app-polyfill/stable'

import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter as Router } from 'react-router-dom'

import './i18n'
import { StoreProvider } from './stores'
import { UserContextProvider } from './user-context'
import App from './App'

ReactDOM.render(
  <Router>
    <StoreProvider>
      <UserContextProvider>
        <App />
      </UserContextProvider>
    </StoreProvider>
  </Router>,
  document.getElementById('root'),
)
