import React from 'react'
import ReactDOM from 'react-dom'
import ErrorPage from './ErrorPage'

it('renders without crashing', () => {
    expect(() => {
        const div = document.createElement('div')
        ReactDOM.render(<ErrorPage />, div)
    }).not.toThrow()
})
