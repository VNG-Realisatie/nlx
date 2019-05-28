// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import 'react-app-polyfill/ie11'

import React from 'react'
import ReactDOM from 'react-dom'
import App from './components/App'
import store from './store';

import { BrowserRouter } from 'react-router-dom'
import { Provider } from 'react-redux'

ReactDOM.render(
  <Provider store={store}>
    <BrowserRouter>
      <App />
    </BrowserRouter>
  </Provider>,
  document.getElementById('root'),
)
