// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'
import PropTypes from 'prop-types'

const selectedData = [
    'created',
    'logrecord-id',
    'source_organization',
    'service_name',
]

class TableRow extends Component {
    render() {
        const { data, code } = this.props
        const time = new Date(data['created']).toLocaleString()
        return (
            <tr className={code && 'code'}>
                {selectedData.map((property, i) => (
                    <td key={i}>
                        {property === 'created'
                            ? time
                            : data[property].toString()}
                    </td>
                ))}
                <td>
                    {Object.keys(data['data']).map((key) => (
                        <span key={key}>
                            <b>{key}</b> {data['data'][key]}
                        </span>
                    ))}
                </td>
            </tr>
        )
    }
}

TableRow.propTypes = {
    data: PropTypes.shape({
        created: PropTypes.string.isRequired,
        'logrecord-id': PropTypes.string.isRequired,
        source_organization: PropTypes.string.isRequired,
        service_name: PropTypes.string.isRequired,
        data: PropTypes.object.isRequired,
    }),
}

export default TableRow
