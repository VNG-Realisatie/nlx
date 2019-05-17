// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import PropTypes from 'prop-types'

import {
    withStyles,
    Table,
    TableBody,
    TableCell,
    TableHead,
    TablePagination,
    TableRow,
    IconButton,
} from '@material-ui/core'
import { InfoOutlined } from '@material-ui/icons'

import styles from '../styles/Table'

class EnhancedTable extends React.Component {
    state = {
        modalState: true,
        modalContent: '',
    }

    openModal = (id) => {
        this.props.onDetails(id)
    }

    handleChangePage = (event, page) => {
        this.props.onOptionsChange({
            ...this.props.options,
            page,
        })
    }

    handleChangeRowsPerPage = (event) => {
        this.props.onOptionsChange({
            ...this.props.options,
            page: 0,
            rowsPerPage: event.target.value,
        })
    }

    getTableHead = () => {
        const { classes, cols } = this.props

        const colsHtml = cols.map((col) => {
            return (
                <TableCell
                    key={col.id}
                    numeric={col.numeric}
                    padding={col.disablePadding ? 'none' : 'default'}
                    style={col.width ? { width: col.width } : {}}
                >
                    {col.label}
                </TableCell>
            )
        })

        return (
            <TableRow>
                <TableCell padding="none" className={classes.firstTableCell} />
                {colsHtml}
            </TableRow>
        )
    }

    getTableRow = (cols, row) => {
        const cellsHtml = cols.map((col, i) => {
            let data = row[col.id]
            return (
                <TableCell key={col.id} padding={i === 0 ? 'none' : 'default'}>
                    {data}
                </TableCell>
            )
        })

        return (
            <TableRow key={row.id} tabIndex={-1}>
                <TableCell padding="none">
                    <IconButton onClick={() => this.openModal(row.id)}>
                        <InfoOutlined
                            color="secondary"
                            style={{ fontSize: 20 }}
                        />
                    </IconButton>
                </TableCell>
                {cellsHtml}
            </TableRow>
        )
    }

    getTableBody = () => {
        const { cols, data } = this.props

        const tableBody = data.map((row) => {
            return this.getTableRow(cols, row)
        })

        return tableBody
    }

    render() {
        const { classes, options } = this.props

        return (
            <React.Fragment>
                <div className={classes.tableWrapper}>
                    <Table>
                        <TableHead>{this.getTableHead()}</TableHead>
                        <TableBody>{this.getTableBody()}</TableBody>
                    </Table>
                </div>
                <TablePagination
                    component="div"
                    count={options.rowCount}
                    rowsPerPage={options.rowsPerPage}
                    rowsPerPageOptions={options.rowsPerPageOptions}
                    page={options.page}
                    backIconButtonProps={{
                        'aria-label': 'Vorige pagina',
                    }}
                    nextIconButtonProps={{
                        'aria-label': 'Volgende pagina',
                    }}
                    labelRowsPerPage="Aantal logs per pagina:"
                    onChangePage={this.handleChangePage}
                    onChangeRowsPerPage={this.handleChangeRowsPerPage}
                />
            </React.Fragment>
        )
    }
}

EnhancedTable.propTypes = {
    classes: PropTypes.object.isRequired,
}

export default withStyles(styles)(EnhancedTable)
