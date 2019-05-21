// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import Table from './Table'

it('renders without crashing', () => {
    expect(() => {
        const div = document.createElement('div')
        ReactDOM.render(<Table heads={[]} rows={[]} />, div)
    }).not.toThrow()
})
