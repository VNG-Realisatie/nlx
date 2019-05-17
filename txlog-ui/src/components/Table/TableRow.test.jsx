// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import ReactDOM from 'react-dom'
import TableRow from './TableRow'

it('renders without crashing', () => {
    const data = {
        created: '2018-12-12T15:35:52.177777Z',
        'logrecord-id': 'foo',
        source_organization: 'source_organization',
        service_name: 'service_name',
        data: {},
    }

    expect(() => {
        const tbody = document.createElement('tbody')
        ReactDOM.render(<TableRow data={data} />, tbody)
    }).not.toThrow()
})
