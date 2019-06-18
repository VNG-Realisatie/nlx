// Copyright © VNG Realisatie 2019
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'

describe('App', () => {
    it('renders without crashing', () => {
        expect(() => {
            const div = document.createElement('div')
            ReactDOM.render(<App />, div)
        }).not.toThrow()
    })
})
