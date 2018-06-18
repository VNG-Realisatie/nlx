import React from 'react'
import plugIcon from '../assets/icons/plug.svg'

export default class Directory extends React.Component {
    render() {
        const {
            name,
            service,
            api,
            offline
        } = this.props

        return (
            <tr className={offline && "status-inactive"}>
                <td>
                    <svg id="status" viewBox="0 0 10 10" width="10px" height="10px"><circle cx="5" cy="14" r="5" transform="translate(0 -9)" fill="currentColor" fillRule="evenodd"></circle></svg>
                </td>
                <td>
                    <span>{name}</span>
                </td>
                <td>
                    <span>{service}</span>
                </td>
                <td>
                    {api &&
                        <button type="button" className="btn btn-icon" data-toggle="tooltip" title="Copy API url">
                            <img src={plugIcon} alt="api"/>
                        </button>
                    }
                </td>
            </tr>
        )
    }
}
