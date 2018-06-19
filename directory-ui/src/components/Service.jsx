import React from 'react'
import classnames from 'classnames'
import plugIcon from '../assets/icons/plug.svg'
import $ from 'jquery'
import {Link} from 'react-router-dom'
import copy from 'copy-to-clipboard'

export default class Directory extends React.Component {
    render() {
        const {
            organizationName,
            name,
            inwayAddresses,
            documentationUrl
        } = this.props

        const apiAddress = `http://{your-outway-address}:12018/${organizationName}/${name}`

        $(function () {
            $('[data-toggle="tooltip"]').tooltip({
                trigger: 'hover focus manual',
                placement: 'bottom',
                template: '<div class="tooltip" role="tooltip"><div class="tooltip-inner"></div></div>'
            })
        })

        return (
            <React.Fragment>

                <tr className={classnames({"status-inactive": !inwayAddresses})}>
                    <td>
                        <svg id="status" viewBox="0 0 10 10" width="10px" height="10px"><circle cx="5" cy="14" r="5" transform="translate(0 -9)" fill="currentColor" fillRule="evenodd"></circle></svg>
                    </td>
                    <td>
                        {documentationUrl ?
                            <Link to="/doc">{organizationName}</Link>
                            :
                            <span>{organizationName}</span>
                        }
                    </td>
                    <td>
                        <span>{name}</span>
                    </td>
                    <td>
                        <button
                            type="button" className="btn btn-icon"
                            data-toggle="tooltip" title="Copy API address"
                            style={{marginTop: '-4px'}}
                            onClick={() => copy(apiAddress)}
                        >
                            <img src={plugIcon} alt="api" style={{marginTop: '-2px'}} />
                        </button>
                    </td>
                </tr>
            </React.Fragment>
        )
    }
}
