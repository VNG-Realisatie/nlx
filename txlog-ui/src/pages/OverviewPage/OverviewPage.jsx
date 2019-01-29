import React from 'react'
import Switch from '../../components/Switch/Switch'
import Table from '../../components/Table/Table'
import { TablePagination } from '@material-ui/core'
import axios from 'axios'

import './OverviewPage.scss'
import ErrorPage from '../ErrorPage/ErrorPage'
import { Spinner } from '@commonground/design-system'

import { MuiThemeProvider } from '@material-ui/core/styles'
import muiTheme from '../../styles/muiTheme'
import { InfoOutlined } from '@material-ui/icons'
import { relativeToFullUrl } from '../../helpers/api'

export class Overview extends React.Component {
    state = {
        showLogs: 'out',
        records: [],
        displayOnlyContaining: '',
        sortBy: 'created',
        sortAscending: true,
        loading: true,
        error: false,
        page: 0,
        rowsPerPage: 5,
        rowsPerPageOptions: [5, 10, 20],
        rowCount: 0,
        theads: [
            {
                label: 'Date',
                width: '200px',
            },
            {
                label: 'Id',
            },
            {
                label: 'Organisation',
            },
            {
                label: 'Service',
            },
            {
                label: 'Data',
            },
        ],
    }

    getLogs = ({ showLogs, page, rowsPerPage }) => {
        let options = {
            params: {
                page,
                rowsPerPage,
            },
        }

        let apiPoint = relativeToFullUrl(`/${showLogs}`)

        axios
            .get(apiPoint, options)
            .then((resp) => {
                const { records, page, rowCount, rowsPerPage } = resp.data
                this.setState({
                    showLogs,
                    records,
                    page,
                    rowCount,
                    rowsPerPage,
                    loading: false,
                    error: false,
                })
            })
            .catch((e) => {
                this.setState({
                    records: [],
                    page: 0,
                    rowCount: 0,
                    loading: false,
                    error: true,
                })
            })
    }

    selectLogType = (val) => {
        const { rowsPerPage } = this.state
        this.getLogs({
            showLogs: val,
            page: 0,
            rowsPerPage,
        })
    }

    toggleLogType = (event) => {
        if (this.state.showLogs === 'in') {
            this.selectLogType('out')
        } else {
            this.selectLogType('in')
        }
    }

    getSwitchHtml = () => {
        const inactiveStyle = {
            color: '#ADB5BD',
        }
        const activeStyle = {
            color: '#FEBF24',
        }
        return (
            <div className="d-flex justify-content-center">
                <button
                    className="btn btn-small mr-2"
                    style={
                        this.state.showLogs === 'in'
                            ? activeStyle
                            : inactiveStyle
                    }
                    onClick={() => this.selectLogType('in')}
                >
                    IN
                </button>
                <Switch
                    onChange={this.toggleLogType}
                    checked={this.state.showLogs === 'out'}
                    id="inout"
                    alwaysOn
                />
                <button
                    className="btn btn-small"
                    style={
                        this.state.showLogs === 'out'
                            ? activeStyle
                            : inactiveStyle
                    }
                    onClick={() => this.selectLogType('out')}
                >
                    OUT
                </button>
            </div>
        )
    }

    getTableHtml = () => {
        const { theads, records } = this.state
        let tableHtml = (
            <p className="overview-page-empty">
                <InfoOutlined /> <br />
                No logs to show
            </p>
        )
        if (records.length > 0) {
            tableHtml = (
                <Table
                    heads={theads}
                    rows={records}
                    onSort={null}
                    sortBy="disabled"
                    sortAscending={this.state.sortAscending}
                />
            )
        }
        return tableHtml
    }

    getPaginationHtml = () => {
        const { rowCount, rowsPerPage, rowsPerPageOptions, page } = this.state
        let pagHtml = null
        if (rowCount) {
            pagHtml = (
                <MuiThemeProvider theme={muiTheme}>
                    <TablePagination
                        component="div"
                        count={rowCount}
                        rowsPerPage={rowsPerPage}
                        rowsPerPageOptions={rowsPerPageOptions}
                        page={page}
                        backIconButtonProps={{
                            'aria-label': 'Back',
                        }}
                        nextIconButtonProps={{
                            'aria-label': 'Next',
                        }}
                        labelRowsPerPage="Rows per page:"
                        onChangePage={this.handlePageChange}
                        onChangeRowsPerPage={this.handleChangeRowsPerPage}
                    />
                </MuiThemeProvider>
            )
        }
        return pagHtml
    }

    handlePageChange = (event, page) => {
        const { showLogs, rowsPerPage } = this.state
        this.getLogs({
            showLogs,
            page,
            rowsPerPage,
        })
    }

    handleChangeRowsPerPage = (event) => {
        const { showLogs, page } = this.state
        this.getLogs({
            showLogs,
            page,
            rowsPerPage: event.target.value,
        })
    }

    render() {
        const { loading, error } = this.state

        if (loading) {
            return <Spinner />
        }

        if (error) {
            return <ErrorPage />
        }

        return (
            <React.Fragment>
                <section className="nlx-nav-section">
                    <div className="container nlx-nav-panel">
                        {this.getSwitchHtml()}

                        {this.getPaginationHtml()}
                    </div>
                </section>
                <section className="nlx-content">
                    <div className="container mb-4">{this.getTableHtml()}</div>
                </section>
            </React.Fragment>
        )
    }

    componentDidMount() {
        this.getLogs(this.state)
    }
}

export default Overview
