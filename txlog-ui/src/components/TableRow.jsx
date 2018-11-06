import React from 'react'

const selectedData = [
    'created',
    'logrecord-id',
    'source_organization',
    'service_name',
]

export default class TableRow extends React.Component {
    render() {
        const { data, code } = this.props

        const time = new Date(data['created']).toLocaleString()

        return (
            <React.Fragment>
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
            </React.Fragment>
        )
    }
}
