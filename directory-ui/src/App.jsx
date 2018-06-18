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
            sortBy: 'status',
            sortAscending: true,
            services: [],
            filteredServices: []
        }

        this.searchOnChange = this.searchOnChange.bind(this)
        this.switchOnChange = this.switchOnChange.bind(this)
    }

    componentDidMount() {
        axios.get(`/api/directory/list-services`)
            .then(res => {
                const services = res.data.services;
                this.setState({ services });
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
                        <Services serviceList={filteredServices} />
                    </div>
                </section>
            </div>
        )
    }
}

export default App;
