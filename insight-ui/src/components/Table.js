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
    TableSortLabel,
    IconButton,
} from '@material-ui/core'
import { InfoOutlined } from '@material-ui/icons'

import styles from '../styles/Table'

function desc(a, b, orderBy) {
    if (b[orderBy] < a[orderBy]) {
        return -1
    }
    if (b[orderBy] > a[orderBy]) {
        return 1
    }
    return 0
}

function stableSort(array, cmp) {
    const stabilizedThis = array.map((el, index) => [el, index])
    stabilizedThis.sort((a, b) => {
        const order = cmp(a[0], b[0])
        if (order !== 0) return order
        return a[1] - b[1]
    })
    return stabilizedThis.map((el) => el[0])
}

function getSorting(order, orderBy) {
    return order === 'desc'
        ? (a, b) => desc(a, b, orderBy)
        : (a, b) => -desc(a, b, orderBy)
}

class EnhancedTable extends React.Component {
    state = {
        order: 'asc',
        orderBy: 'date',
        // cols:[],
        // data: [],
        page: 0,
        rowsPerPage: 10,
        rowsPerPageOptions: [10, 25, 50],
        modalState: true,
        modalContent: '',
    }

    handleRequestSort = (property) => {
        // debugger
        const orderBy = property
        let order = 'asc'

        if (this.state.orderBy === property && this.state.order === 'asc') {
            order = 'desc'
        }

        this.setState({ order, orderBy })
    }

    openModal = (id) => {
        this.props.onDetails(id)
    }

    handleChangePage = (event, page) => {
        this.setState({ page })
    }

    handleChangeRowsPerPage = (event) => {
        this.setState({ rowsPerPage: event.target.value })
    }

    getTableHead = (cols) => {
        const { order, orderBy } = this.state

        const { classes } = this.props

        const colsHtml = cols.map((col) => {
            return (
                <TableCell
                    key={col.id}
                    numeric={col.numeric}
                    padding={col.disablePadding ? 'none' : 'default'}
                    sortDirection={orderBy === col.id ? order : false}
                    style={col.width ? { width: col.width } : {}}
                >
                    <TableSortLabel
                        active={orderBy === col.id}
                        direction={order}
                        onClick={() => this.handleRequestSort(col.id)}
                    >
                        {col.label}
                    </TableSortLabel>
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
        const { order, orderBy, rowsPerPage, page } = this.state

        const tableBody = stableSort(data, getSorting(order, orderBy))
            .slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage)
            .map((row) => {
                return this.getTableRow(cols, row)
            })

        return tableBody
    }

    render() {
        const { classes, cols, data } = this.props
        const { rowsPerPageOptions, rowsPerPage, page } = this.state

        return (
            <React.Fragment>
                <div className={classes.tableWrapper}>
                    <Table>
                        <TableHead>{this.getTableHead(cols)}</TableHead>
                        <TableBody>{this.getTableBody()}</TableBody>
                    </Table>
                </div>
                <TablePagination
                    component="div"
                    count={data.length}
                    rowsPerPage={rowsPerPage}
                    rowsPerPageOptions={rowsPerPageOptions}
                    page={page}
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
