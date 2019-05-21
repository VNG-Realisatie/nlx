// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import SortArrow from './SortArrow'

it('renders without crashing', () => {
    expect(() => {
        const div = document.createElement('div')
        ReactDOM.render(<SortArrow />, div)
    }).not.toThrow()
})
