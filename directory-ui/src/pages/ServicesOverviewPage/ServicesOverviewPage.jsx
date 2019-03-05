import React, { Component } from 'react'

import { Spinner } from '@commonground/design-system'

import ErrorMessage from '../../components/ErrorMessage/ErrorMessage'
import Container from '../../components/Container/Container'
import { StyledFilters, StyledServicesTableContainer } from './ServicesOverviewPage.styles';
import { mapListServicesAPIResponse } from './map-list-services-api-response';

const ESCAPE_KEY_CODE = 27

class ServicesOverviewPage extends Component {
    constructor(props) {
        super(props)

        this.state = {
            loading: true,
            error: null,
            services: [],
            query: '',
            displayOfflineServices: true
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
        if (event.keyCode === ESCAPE_KEY_CODE) {
            this.setState({ query: '' })
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

    searchOnChange(query) {
        this.setState({ query })
    }

    switchOnChange(checked) {
        this.setState({ displayOfflineServices: checked })
    }

    render() {
        const { displayOfflineServices, query, loading, error, services } = this.state

        if (loading) {
            return <Spinner />
        }

        if (error) {
            return <ErrorMessage />
        }

        return (
            <Container>
                <StyledFilters onQueryChanged={this.searchOnChange}
                                onStatusFilterChanged={this.switchOnChange}
                                queryValue={query}
                />

                <StyledServicesTableContainer services={services}
                                                filterQuery={query}
                                                filterByOnlineServices={!displayOfflineServices}
                                                />
            </Container>
        )
    }
}

export default ServicesOverviewPage
