import React from 'react'
import plugIcon from '../assets/icons/plug.svg'
import $ from 'jquery'
import { Link } from 'react-router-dom'
import copy from 'copy-to-clipboard'
import Highlighter from 'react-highlight-words'

import './Service.css'

export default class Directory extends React.Component {
    // local state object
    state = {
        // currently the user replaces {your-outway-address} manually
        apiAddress: `http://{your-outway-address}:12018/${
            this.props.data.organization_name
        }/${this.props.data.service_name}`,
        // default tooltip message
        tooltip: 'Copy API address to clipboard',
        // messages shown after succeful/failed copy to clipboard
        copyConfirm: 'Copied!',
        copyFailed: 'Failed',
        // reference to DOM element
        copyBtn: React.createRef(),
    }
    /**
     * Runs after component output has been rendered to DOM
     */
    componentDidMount() {
        // bind jQuery/BS tooltip
        this.bindTooltip()
    }
    bindTooltip = () => {
        // debugger
        let el = this.state.copyBtn.current
        if (el) {
            // console.log("Show tooltip on...", el);
            // pass definitions
            $(el).tooltip({
                trigger: 'manual focus hover',
            })
            // after hidding reset content to default message
            $(el).on('hidden.bs.tooltip', () => {
                this.updateTooltip()
            })
        }
    }
    /**
     * Copy apiAdress to clipboard and update tooltip message.
     * apiAddress is constructed using info provided by parent.
     */
    copyToClipboard = () => {
        // copy constructed url to clipboard
        let result = copy(this.state.apiAddress)
        // update tooltip message
        if (result) {
            // console.log("copyToClipboard...success...",result);
            this.updateTooltip(this.state.copyConfirm)
        } else {
            // console.warn("Service.copyToClipboard...failed...",result);
            this.updateTooltip(this.state.copyFailed)
        }
        // show updated tooltip
        this.showTooltip()
    }
    /**
     * Manually show tooltip,
     * this function is used after changing the content of tooltip
     */
    showTooltip = () => {
        let el = this.state.copyBtn.current
        if (el) {
            // manual trigger to show
            $(el).tooltip('show')
        }
    }
    /**
     * Manually hide tooltip,
     * this function is used after manual show is used
     */
    hideTooltip = () => {
        // get current button DOM element
        let el = this.state.copyBtn.current
        if (el) {
            $(el).tooltip('hide')
        }
    }
    /**
     * Dinamically update tooltip content,
     * after updating content call showTooltip
     * to force update on tooltip that is already shown
     */
    updateTooltip = (msg = this.state.tooltip) => {
        // get current button DOM element
        let el = this.state.copyBtn.current
        if (el) {
            $(el).attr('data-original-title', msg)
        }
    }

    render() {
        const { data, highlightWords } = this.props
        return (
            <React.Fragment>
                <tr>
                    <td
                        style={{
                            color: data.inway_addresses ? '#B3E87B' : '#FF8282',
                        }}
                    >
                        <svg
                            id="status"
                            className="service__status"
                            viewBox="0 0 10 10"
                            width="100%"
                            height="1rem"
                        >
                            <circle
                                cx="5"
                                cy="14"
                                r="5"
                                transform="translate(0 -9)"
                                fill="currentColor"
                                fillRule="evenodd"
                            />
                        </svg>
                    </td>
                    <td>
                        <Highlighter
                            searchWords={[highlightWords]}
                            autoEscape={true}
                            textToHighlight={data.organization_name}
                        />
                    </td>
                    <td>
                        {data.api_specification_type ? (
                            <Link
                                to={`/documentation/${data.organization_name}/${
                                    data.service_name
                                }`}
                            >
                                <strong>
                                    <Highlighter
                                        searchWords={[highlightWords]}
                                        autoEscape={true}
                                        textToHighlight={data.service_name}
                                    />
                                </strong>
                            </Link>
                        ) : (
                            <Highlighter
                                searchWords={[highlightWords]}
                                autoEscape={true}
                                textToHighlight={data.service_name}
                            />
                        )}
                    </td>
                    <td>{data.api_specification_type || '-'}</td>
                    <td className="table-cell-center">
                        <button
                            style={{ marginTop: '-4px' }}
                            type="button"
                            className="btn btn-icon"
                            data-toggle="tooltip"
                            title={this.state.tooltip}
                            ref={this.state.copyBtn}
                            onClick={this.copyToClipboard}
                            onMouseEnter={this.showTooltip}
                            onMouseLeave={this.hideTooltip}
                        >
                            <img
                                src={plugIcon}
                                alt="api"
                                style={{ marginTop: '-2px' }}
                            />
                        </button>
                    </td>
                </tr>
            </React.Fragment>
        )
    }
}
