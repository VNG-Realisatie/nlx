import React from 'react'
import classnames from 'classnames'
import plugIcon from '../assets/icons/plug.svg'

export default class Directory extends React.Component {
    render() {
        const {
            name,
            organizationName,
            inwayAddresses,
            documentationUrl
        } = this.props

        return (
            <tr className={classnames({
                    "status-inactive": !inwayAddresses
                })}
            >
                <td>
                    <svg id="status" viewBox="0 0 10 10" width="10px" height="10px"><circle cx="5" cy="14" r="5" transform="translate(0 -9)" fill="currentColor" fillRule="evenodd"></circle></svg>
                </td>
                <td>
                    {documentationUrl ?
                        <a href={documentationUrl}>{organizationName}</a>
                        :
                        <span>{organizationName}</span>
                    }
                </td>
                <td>
                    <span>{name}</span>
                </td>
                <td>
                    <button type="button" className="btn btn-icon" data-toggle="tooltip" title="Copy API url">
                        <img src={plugIcon} alt="api"/>
                    </button>
                </td>
            </tr>
        )
    }
}
