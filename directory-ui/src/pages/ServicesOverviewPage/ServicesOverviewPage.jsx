import React, { Component } from 'react'

import { Spinner } from '@commonground/design-system'

import ErrorMessage from '../../components/ErrorMessage/ErrorMessage'
import Search from '../../components/Search/Search'
import Switch from '../../components/Switch/Switch'
import ServicesTableContainer from '../../containers/ServicesTableContainer/ServicesTableContainer'

export const mapListServicesAPIResponse = response =>
  response.services.map(service => ({
    organization: service['organization_name'],
    name: service['service_name'],
    status: service['inway_addresses'] ? 'online' : 'offline',
    documentationLink: service['documentation_url'],
    apiType: service['api_specification_type']
  }))

class ServicesOverviewPage extends Component {
    constructor(props) {
        super(props)

        this.state = {
            loading: true,
            error: null,
            services: []
        }

        this.searchOnChange = this.searchOnChange.bind(this)
        this.switchOnChange = this.switchOnChange.bind(this)
        this.escFunction = this.escFunction.bind(this)
    }

    fetchServices() {
      return fetch(`/api/directory/list-services`)
        .then(response => response.json())
    }

    escFunction(event) {
        if (event.keyCode === 27) {
            this.setState({ displayOnlyContaining: '' })
        }
    }

    componentDidMount() {
        document.addEventListener('keydown', this.escFunction, false)

        this
          .fetchServices()
          .then(response => mapListServicesAPIResponse(response))
          .then(services => {
            this.setState({ loading: false, error: false, services })
          })
          .catch(() => {
            this.setState({ loading: false, error: true })
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

    render() {
        const { displayOnlyOnline, displayOnlyContaining, loading, error, services } = this.state

        if (loading) {
            return <Spinner />
        }

        if (error) {
            return <ErrorMessage />
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
                                        value={displayOnlyContaining}
                                        placeholder="Filter services"
                                        filter
                                    />
                                </div>
                            </div>
                            <div className="col-sm-6 col-lg-6 d-flex align-items-center justify-content-center justify-content-sm-start">
                                <Switch
                                    id="switch1"
                                    onChange={this.switchOnChange}
                                    checked={displayOnlyOnline}
                                >
                                    Only online services
                                </Switch>
                            </div>
                        </div>
                    </div>
                </section>
                <section>
                    <div className="container">
                        <ServicesTableContainer services={services}
                                                filterQuery={displayOnlyContaining}
                                                filterByOnlineServices={displayOnlyOnline}
                        />
                    </div>
                </section>
            </React.Fragment>
        )
    }
}

export default ServicesOverviewPage
