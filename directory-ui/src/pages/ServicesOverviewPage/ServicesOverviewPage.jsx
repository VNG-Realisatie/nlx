// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Component } from 'react'
import debounce from 'debounce'

import { Spinner } from '@commonground/design-system'

import ErrorMessage from '../../components/ErrorMessage/ErrorMessage'
import Container from '../../components/Container/Container'
import { StyledFilters, StyledServicesTableContainer } from './ServicesOverviewPage.styles';
import { mapListServicesAPIResponse } from './map-list-services-api-response';
import ServiceDetailPane from '../../components/ServiceDetailPane';

const ESCAPE_KEY_CODE = 27

class ServicesOverviewPage extends Component {
    constructor(props) {
        super(props)

        const { location, history } = this.props

        const urlParams = new URLSearchParams(location.search)

        this.state = {
            loading: true,
            error: null,
            services: [],
            query: urlParams.get('q') || '',
            debouncedQuery: urlParams.get('q') || '',
            displayOfflineServices: true,
            selectedService: null
        }

        this.searchOnChange = this.searchOnChange.bind(this)
        this.switchOnChange = this.switchOnChange.bind(this)
        this.escFunction = this.escFunction.bind(this)
        this.onServiceClickedHandler = this.onServiceClickedHandler.bind(this)
        this.detailPaneCloseHandler = this.detailPaneCloseHandler.bind(this)

        this.searchOnChangeDebouncable = (query) => {
            this.setState({ debouncedQuery: query })
            history.push(`?q=${encodeURIComponent(query)}`)
        }

        this.searchOnChangeDebounced = debounce(this.searchOnChangeDebouncable, 400)
    }

    onServiceClickedHandler(service) {
        this.setState({
            selectedService: service
        });
    }

    detailPaneCloseHandler() {
        this.setState({
            selectedService: null
        });
    }

    fetchServices() {
        return fetch(`/api/directory/list-services`,{
            headers: {
                'Content-Type': 'application/json',
            },
        }).then(response => response.json())
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
        this.searchOnChangeDebounced(query)
    }

    switchOnChange(checked) {
        this.setState({ displayOfflineServices: checked })
    }

    render() {
        const { displayOfflineServices, query, debouncedQuery, loading, error, services, selectedService } = this.state

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
                                                sortBy='organization'
                                                sortOrder='asc'
                                                filterQuery={debouncedQuery}
                                                filterByOnlineServices={!displayOfflineServices}
                                                onServiceClickedHandler={(service) => this.onServiceClickedHandler(service)}
                                                />
                {
                    selectedService ?
                        <ServiceDetailPane organizationName={selectedService.organization}
                                           contactEmail={selectedService.contactEmailAddress}
                                           name={selectedService.name}
                                           closeHandler={this.detailPaneCloseHandler}
                        /> : null
                }
            </Container>
        )
    }
}

ServicesOverviewPage.defaultProps = {
    location: window.location,
    history: window.history
}

export default ServicesOverviewPage
