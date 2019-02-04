import React, { Component } from 'react'
import axios from 'axios'

import { Spinner } from '@commonground/design-system'

import ErrorMessage from '../../components/ErrorMessage/ErrorMessage'
import Search from '../../components/Search/Search'
import Switch from '../../components/Switch/Switch'
import Services from '../../components/Services/Services'

export default class ServicesOverviewPage extends Component {
    constructor(props) {
        super(props)

        this.state = {
            displayOnlyContaining: '',
            displayOnlyOnline: false,
            sortBy: 'organization_name',
            sortAscending: true,
            services: [],
            loading: true,
            error: false,
        }

        this.searchOnChange = this.searchOnChange.bind(this)
        this.switchOnChange = this.switchOnChange.bind(this)
        this.onSort = this.onSort.bind(this)
        this.escFunction = this.escFunction.bind(this)
    }

    escFunction(event) {
        if (event.keyCode === 27) {
            this.setState({ displayOnlyContaining: '' })
        }
    }

    componentDidMount() {
        document.addEventListener('keydown', this.escFunction, false)
        axios
            .get(`/api/directory/list-services`)
            .then((res) => {
                const services = res.data.services
                if (services) {
                    this.setState({
                        services,
                        loading: false,
                        error: false,
                    })
                }
            })
            .catch((e) => {
                this.setState({
                    loading: false,
                    error: true,
                })
            })
    }

    componentWillUnmount() {
        document.removeEventListener('keydown', this.escFunction, false)
    }

    searchOnChange(e) {
        this.setState({ displayOnlyContaining: e.target.value })
    }

    switchOnChange() {
        this.setState({ displayOnlyOnline: !this.state.displayOnlyOnline })
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

    render() {
        const {
            services,
            displayOnlyOnline,
            displayOnlyContaining,
        } = this.state

        if (this.state.loading) {
            return <Spinner />
        }

        if (this.state.error) {
            return <ErrorMessage />
        }

        const filteredServices = services.filter((service) => {
            if (displayOnlyOnline) {
                if (!service.inway_addresses) {
                    return false
                }
            }

            if (displayOnlyContaining) {
                if (
                    !service.service_name
                        .toLowerCase()
                        .includes(displayOnlyContaining.toLowerCase()) &&
                    !service.organization_name
                        .toLowerCase()
                        .includes(displayOnlyContaining.toLowerCase())
                ) {
                    return false
                }
            }

            return true
        })

        const { sortBy, sortAscending } = this.state

        let filteredAndSortedServices = []
            .concat(filteredServices)
            .sort((a, b) => {
                switch (sortBy) {
                    case 'inway_addresses':
                        return a.inway_addresses > b.inway_addresses
                    case 'organization_name':
                        return a.organization_name > b.organization_name
                    case 'name':
                        return a.service_name > b.service_name
                    default:
                        return false
                }
            })
            .map((item) => {
                return item
            })

        if (!sortAscending) {
            filteredAndSortedServices = filteredAndSortedServices.reverse()
        }

        return (
            <React.Fragment>
                <section>
                    <div className="container">
                        <div className="row">
                            <div className="col-sm-6 col-lg-4 offset-lg-2">
                                <div className="mb-4 mb-sm-0">
                                    <Search
                                        onChange={this.searchOnChange}
                                        value={this.state.displayOnlyContaining}
                                        placeholder="Filter services"
                                        filter
                                    />
                                </div>
                            </div>
                            <div className="col-sm-6 col-lg-6 d-flex align-items-center justify-content-center justify-content-sm-start">
                                <Switch
                                    id="switch1"
                                    onChange={this.switchOnChange}
                                    checked={this.state.displayOnlyOnline}
                                >
                                    Only online services
                                </Switch>
                            </div>
                        </div>
                    </div>
                </section>
                <section>
                    <div className="container">
                        <Services
                            serviceList={filteredAndSortedServices}
                            onSort={this.onSort}
                            sortBy={this.state.sortBy}
                            sortAscending={this.state.sortAscending}
                            displayOnlyContaining={
                                this.state.displayOnlyContaining
                            }
                        />
                    </div>
                </section>
            </React.Fragment>
        )
    }
}
