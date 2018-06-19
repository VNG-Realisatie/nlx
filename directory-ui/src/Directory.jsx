import React from 'react'
import Search from './components/Search'
import Switch from './components/Switch'
import Services from './components/Services'
import axios from 'axios';

export default class Directory extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            displayOnlyContaining: '',
            displayOnlyOnline: false,
            sortBy: 'organization_name',
            sortAscending: true,
            services: [],
            filteredServices: []
        }

        this.searchOnChange = this.searchOnChange.bind(this)
        this.switchOnChange = this.switchOnChange.bind(this)
        this.onSort = this.onSort.bind(this)
    }

    componentDidMount() {
        axios.get(`/api/directory/list-services`)
            .then(res => {
                const services = res.data.services;
                if (services) {
                    this.setState({ services })
                }
            })
            .catch(e => {
                console.error(e);
            })
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
            sortAscending: true
        })
    }

    render() {
        const {
            services,
            displayOnlyOnline,
            displayOnlyContaining
        } = this.state

        const filteredServices = services.filter(service => {
            if (displayOnlyOnline) {
                if (!service.inway_addresses) {
                    return false
                }
            }

            if (displayOnlyContaining) {
                if (
                    !service.name.toLowerCase().includes(displayOnlyContaining.toLowerCase()) &&
                    !service.organization_name.toLowerCase().includes(displayOnlyContaining.toLowerCase())
                ) {
                    return false
                }
            }

            return true
        })

        const {
            sortBy,
            sortAscending
        } = this.state

        let filteredAndSortedServices = [].concat(filteredServices)
            .sort((a, b) => {
                switch (sortBy) {
                    case "inway_addresses": {
                        return (a.inway_addresses > b.inway_addresses)
                    }
                    case "organization_name": {
                        return (a.organization_name > b.organization_name)
                    }
                    case "name": {
                        return (a.name > b.name)
                    }
                    default: {
                        return false
                    }
                }
            })
            .map((item) => {return item}
        )

        if (!sortAscending) {
            filteredAndSortedServices = filteredAndSortedServices.reverse()
        }

        return (
            <React.Fragment>
                <section>
                    <div className="container">
                        <div className="row">
                            <div className="col-sm-6 col-lg-4 offset-lg-2">
                                <Search onChange={this.searchOnChange} value={this.state.displayOnlyContaining} placeholder="Filter services" />
                            </div>
                            <div className="col-sm-6 col-lg-6 d-flex align-items-center">
                                <Switch id="switch1" onChange={this.switchOnChange} checked={this.state.displayOnlyOnline}>Only online services</Switch>
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
                            />
                    </div>
                </section>
            </React.Fragment>
        )
    }
}
