// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'
import App from './App'

describe('App', () => {
    it('renders without crashing', () => {
        expect(() => {
            const div = document.createElement('div')
            ReactDOM.render(<App />, div)
        }).not.toThrow()
    })
})
