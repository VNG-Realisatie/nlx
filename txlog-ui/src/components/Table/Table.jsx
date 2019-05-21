// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'
import PropTypes from 'prop-types'
import TableRow from './TableRow'
import SortArrow from './SortArrow'

class Table extends Component {
    render() {
        const { heads, rows, sortBy, sortAscending } = this.props

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
                                        {sortBy === col.sortBy ? (
                                            <SortArrow
                                                sortAscending={sortAscending}
                                            />
                                        ) : null}
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
