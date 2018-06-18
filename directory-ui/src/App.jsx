import React, { Component } from 'react'
import Navigation from './components/Navigation'
import Search from './components/Search'
import Switch from './components/Switch'
import Services from './components/Services'
import axios from 'axios';

class App extends Component {
    constructor(props) {
        super(props)

        this.state = {
            displayOnlyContaining: '',
            displayOnlyOnline: false,
            sortBy: 'organization_name',
            sortAscending: true,
            services: [
                {
                    name: "d",
                    organization_name: "A",
                    inway_addresses: true,
                    documentation_url: "https://docs.postman-echo.com/"
                },
                {
                    name: "e",
                    organization_name: "B",
                    inway_addresses: false,
                    documentation_url: "https://docs.postman-echo.com/"
                },
                {
                    name: "b",
                    organization_name: "E",
                    inway_addresses: true,
                    documentation_url: "https://docs.postman-echo.com/"
                },
                {
                    name: "f",
                    organization_name: "D",
                    inway_addresses: true,
                    documentation_url: "https://docs.postman-echo.com/"
                },
                {
                    name: "c",
                    organization_name: "C",
                    inway_addresses: false,
                    documentation_url: "https://docs.postman-echo.com/"
                },
                {
                    name: "a",
                    organization_name: "F",
                    inway_addresses: true,
                    documentation_url: "https://docs.postman-echo.com/"
                },
            ],
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
                // this.setState({ services });
            })
            .catch(e => {
                this.errors.push(e)
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

        let sortedServices = [].concat(filteredServices)
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
                }
            })
            .map((item) => {return item}
        )

        if (!sortAscending) {
            sortedServices = sortedServices.reverse()
        }

        return (
            <div className="App">
                <Navigation />
                <section>
                <div className="container">
                    <div className="row">
                        <div className="col-sm-6 col-lg-4 offset-lg-2">
                            <Search onChange={this.searchOnChange} value={this.state.displayOnlyContaining} />
                        </div>
                        <div className="col-sm-6 col-lg-6 d-flex align-items-center">
                            <Switch id="switch1" onChange={this.switchOnChange} checked={this.state.displayOnlyOnline}>Online services</Switch>
                        </div>
                    </div>
                </div>
                </section>
                <section>
                    <div className="container">
                        <Services
                            serviceList={sortedServices}
                            onSort={this.onSort}
                            sortBy={this.state.sortBy}
                            sortAscending={this.state.sortAscending}
                        />
                    </div>
                </section>
            </div>
        )
    }
}

export default App;
