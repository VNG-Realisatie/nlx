// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import '@babel/polyfill'

import React from 'react'
import ReactDOM from 'react-dom'
import { BrowserRouter } from 'react-router-dom';
import App from './App'

ReactDOM.render(<BrowserRouter><App /></BrowserRouter>, document.getElementById('root'))
