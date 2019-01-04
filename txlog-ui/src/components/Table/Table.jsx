import React, { Component } from 'react'
import PropTypes from 'prop-types'
import TableRow from './TableRow'

class Table extends Component {
    render() {
        const { heads, rows, sortBy, sortAscending } = this.props

        const sortArrow = (
            <svg width="8" height="12" viewBox="0 0 8 12" name="sortingArrow">
                <g id="arrow-down" fill="none" fillRule="evenodd">
                    <path
                        id="Shape"
                        fill="currentColor"
                        fillRule="nonzero"
                        transform={
                            sortAscending ? 'rotate(90 4 5)' : 'rotate(-90 4 5)'
                        }
                        d="M5 4h-6v2h6v3l4-4-4-4z"
                    />
                </g>
            </svg>
        )

        return (
            <div className="table-responsive mb-5">
                <table className="table table-bordered">
                    <thead>
                        <tr>
                            {heads.map((col, key) => (
                                <th
                                    scope="col"
                                    key={col.label + key}
                                    className={
                                        sortBy === col.sortBy ? 'sorting' : ''
                                    }
                                    style={{
                                        width: col.width,
                                    }}
                                >
                                    <button
                                        onClick={(e) =>
                                            this.props.onSort(col.sortBy)
                                        }
                                        disabled={!col.sortBy}
                                        style={{
                                            textAlign: col.align,
                                        }}
                                    >
                                        {col.label}{' '}
                                        {sortBy === col.sortBy && sortArrow}
                                    </button>
                                </th>
                            ))}
                        </tr>
                    </thead>
                    <tbody>
                        {rows.map((log) => (
                            <TableRow
                                key={log['logrecord-id']}
                                data={log}
                                code
                            />
                        ))}
                    </tbody>
                </table>
            </div>
        )
    }
}

Table.propTypes = {
    heads: PropTypes.array.isRequired,
    rows: PropTypes.array.isRequired,
}

export default Table
