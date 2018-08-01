import React from 'react'
import plugIcon from '../assets/icons/plug.svg'
import $ from 'jquery'
import {Link} from 'react-router-dom'
import copy from 'copy-to-clipboard'
import Highlighter from "react-highlight-words";

export default class Directory extends React.Component {
    render() {
        const {
            data,
            highlightWords
        } = this.props

        const apiAddress = `http://{your-outway-address}:12018/${data.organization_name}/${data.service_name}`

        $(function () {
            $('[data-toggle="tooltip"]').tooltip({
                trigger: 'hover focus manual',
                placement: 'bottom',
                template: '<div class="tooltip" role="tooltip"><div class="tooltip-inner"></div></div>'
            })
        })

        return (
            <React.Fragment>
                <tr>
                    <td style={{
                        color: data.inway_addresses ? '#B3E87B' : '#FF8282'
                    }}>
                        <svg style={{margin: '0 auto'}} id="status" viewBox="0 0 10 10" width="10px" height="10px"><circle cx="5" cy="14" r="5" transform="translate(0 -9)" fill="currentColor" fillRule="evenodd"></circle></svg>
                    </td>
                    <td>
                        <span>
                            <Highlighter
                                searchWords={[highlightWords]}
                                autoEscape={true}
                                textToHighlight={data.organization_name}
                            />
                        </span>
                    </td>
                    <td>
                        {data.api_specification_type ? (
                            <Link to={`/documentation/${data.organization_name}/${data.service_name}`}>
                                <strong>
                                    <Highlighter
                                        searchWords={[highlightWords]}
                                        autoEscape={true}
                                        textToHighlight={data.service_name}
                                    />
                                </strong>
                            </Link>
                        ) : (
                            <span>
                                <Highlighter
                                    searchWords={[highlightWords]}
                                    autoEscape={true}
                                    textToHighlight={data.service_name}
                                />
                            </span>
                        )}
                    </td>
                    <td>
                        {data.api_specification_type || '-' }
                    </td>
                    <td style={{textAlign: 'center'}}>
                        <button
                            style={{marginTop: '-4px'}}
                            type="button" className="btn btn-icon"
                            data-toggle="tooltip" title="Copy API address"
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
