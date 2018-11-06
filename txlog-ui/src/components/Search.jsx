import React from 'react'
import classnames from 'classnames'

export default class Search extends React.Component {
    render() {
        const { placeholder, filter } = this.props

        return (
            <div
                className={classnames({
                    search: true,
                    'search-filter': filter,
                })}
            >
                <input
                    className="form-control"
                    type="text"
                    placeholder={placeholder}
                    onChange={this.props.onChange}
                    value={this.props.value}
                />
                <button className="search_button" disabled={filter}>
                    <svg
                        width="13px"
                        height="13px"
                        viewBox="0 0 13 13"
                        version="1.1"
                        style={{ marginLeft: '3px' }}
                    >
                        <g
                            id="Page-1"
                            stroke="none"
                            strokeWidth="1"
                            fill="none"
                            fillRule="evenodd"
                        >
                            <g
                                id="Docs"
                                transform="translate(-187.000000, -75.000000)"
                                fill="currentColor"
                            >
                                <path
                                    d="M196.448336,82.8432574 L200,86.3835377 L198.383538,88 L194.831874,84.4483363 C194.04641,84.9492119 193.11296,85.2451839 192.111208,85.2451839 C189.288091,85.2451839 187,82.9570928 187,80.1225919 C187,77.2880911 189.288091,75 192.122592,75 C194.957093,75 197.245184,77.2880911 197.245184,80.1225919 C197.245184,81.1243433 196.949212,82.0464098 196.448336,82.8432574 Z M192.122592,82.9684764 C193.69352,82.9684764 194.968476,81.6935201 194.968476,80.1225919 C194.968476,78.5516637 193.69352,77.2767075 192.122592,77.2767075 C190.551664,77.2767075 189.276708,78.5516637 189.276708,80.1225919 C189.276708,81.6935201 190.551664,82.9684764 192.122592,82.9684764 Z"
                                    id="search"
                                />
                            </g>
                        </g>
                    </svg>
                </button>
            </div>
        )
    }
}
