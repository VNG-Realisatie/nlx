import React from 'react'
import ReactDOM from 'react-dom'
import axios from 'axios'
import MockAdapter from 'axios-mock-adapter'
import OverviewPage from './OverviewPage'

describe('OverviewPage', () => {
    beforeAll(() => {
        let mock = new MockAdapter(axios)

        mock.onGet('/api/in').reply(200, {
            records: [],
        })

        mock.onGet('/api/out').reply(200, {
            records: [],
        })
    })

    it('renders without crashing', () => {
        expect(() => {
            const div = document.createElement('div')
            ReactDOM.render(<OverviewPage />, div)
        }).not.toThrow()
    })
})
