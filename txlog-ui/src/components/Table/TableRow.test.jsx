import React from 'react'
import ReactDOM from 'react-dom'
import TableRow from './TableRow'

it('renders without crashing', () => {
    const data = {
        created: '2018-12-12T15:35:52.177777Z',
        'logrecord-id': 'foo',
        sourceOrganization: 'source_organization',
        serviceName: 'service_name',
        data: {},
    }

    expect(() => {
        const tbody = document.createElement('tbody')
        ReactDOM.render(<TableRow data={data} />, tbody)
    }).not.toThrow()
})
