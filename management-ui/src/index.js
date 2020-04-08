// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import 'react-app-polyfill/ie11'

import React from 'react'
import ReactDOM from 'react-dom'
import './i18n'

import { BrowserRouter as Router } from 'react-router-dom'
import App from './App'

ReactDOM.render(
  <Router>
    <App />
  </Router>,
  document.getElementById('root'),
)
