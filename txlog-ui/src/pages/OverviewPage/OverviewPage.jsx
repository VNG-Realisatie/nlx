import React, { Component } from 'react'
import Switch from '../../components/Switch'
import Search from '../../components/Search'
import Table from '../../components/Table'
import axios from 'axios'

import ErrorPage from '../ErrorPage/ErrorPage'
import Spinner from '../../components/Spinner'

export default class OverviewPage extends Component {
    constructor(props) {
        super(props)

        this.state = {
            showLogs: 'in',
            logsIn: [],
            logsOut: [],
            displayOnlyContaining: '',
            sortBy: 'created',
            sortAscending: true,
            loading: true,
            error: false,
        }

        this.switch = this.switch.bind(this)
        this.filterLogs = this.filterLogs.bind(this)
    }

    getLogs = () => {
        let logsIn, logsOut
        axios
            .get(`/api/in`)
            .then((res) => {
                logsIn = res.data.records
                return axios.get(`/api/out`)
            })
            .then((res) => {
                logsOut = res.data.records
                this.setState({
                    logsIn,
                    logsOut,
                    loading: false,
                    error: false,
                })
            })
            .catch((e) => {
                this.setState({
                    loading: false,
                    error: true,
                })
            })
    }

    componentDidMount() {
        this.getLogs()
    }

    switch(val) {
        if (val === 'in' || val === 'out') {
            this.setState({ showLogs: val })
        } else {
            this.setState({
                showLogs: this.state.showLogs === 'in' ? 'out' : 'in',
            })
        }
    }

    searchOnChange(e) {
        this.setState({ displayOnlyContaining: e.target.value })
    }

    onSort(val) {
        if (this.state.sortBy === val) {
            this.setState({ sortAscending: !this.state.sortAscending })
            return
        }

        this.setState({
            sortBy: val,
            sortAscending: true,
        })
    }

    filterLogs(logs) {
        const { displayOnlyContaining } = this.state

        const filteredLogs = logs.filter((log) => {
            if (displayOnlyContaining) {
                if (
                    !log['logrecord-id']
                        .toLowerCase()
                        .includes(displayOnlyContaining.toLowerCase()) &&
                    !log.source_organization
                        .toLowerCase()
                        .includes(displayOnlyContaining.toLowerCase()) &&
                    !log.service_name
                        .toLowerCase()
                        .includes(displayOnlyContaining.toLowerCase())
                ) {
                    return false
                }
            }

            return true
        })

        return filteredLogs
    }

    render() {
        const { logsIn, logsOut, loading, error } = this.state

        if (loading) {
            return <Spinner />
        }

        if (error) {
            return <ErrorPage />
        }

        const filteredLogsIn = this.filterLogs(logsIn)
        const filteredLogsOut = this.filterLogs(logsOut)

        const theads = [
            {
                label: 'Date',
                // sortBy: 'created',
                width: '200px',
            },
            {
                label: 'Id',
                // sortBy: 'logrecord-id'
            },
            {
                label: 'Organisation',
                // sortBy: 'destination_organization'
            },
            {
                label: 'Service',
                // sortBy: 'service_name'
            },
            {
                label: 'Data',
            },
        ]

        const inactiveStyle = {
            color: '#ADB5BD',
        }

        const activeStyle = {
            color: '#FEBF24',
        }

        return (
            <React.Fragment>
                <section>
                    <div className="container">
                        <div className="d-flex justify-content-center mb-4">
                            <button
                                className="btn btn-small mr-2"
                                style={
                                    this.state.showLogs === 'in'
                                        ? activeStyle
                                        : inactiveStyle
                                }
                                onClick={() => this.switch('in')}
                            >
                                IN
                            </button>
                            <Switch
                                onChange={this.switch}
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
                                onClick={() => this.switch('out')}
                            >
                                OUT
                            </button>
                        </div>
                        <div className="row">
                            <div className="col-sm-6 col-lg-4 offset-lg-4">
                                <Search
                                    onChange={this.searchOnChange.bind(this)}
                                    value={this.state.displayOnlyContaining}
                                    placeholder="Filter logs"
                                    filter
                                />
                            </div>
                        </div>
                    </div>
                </section>
                <section>
                    <div className="container">
                        <Table
                            heads={theads}
                            rows={
                                this.state.showLogs === 'in'
                                    ? filteredLogsIn.reverse()
                                    : filteredLogsOut.reverse()
                            }
                            // rows={filteredLogs}
                            onSort={this.onSort.bind(this)}
                            sortBy={this.state.sortBy}
                            sortAscending={this.state.sortAscending}
                        />
                    </div>
                </section>
            </React.Fragment>
        )
    }
}
