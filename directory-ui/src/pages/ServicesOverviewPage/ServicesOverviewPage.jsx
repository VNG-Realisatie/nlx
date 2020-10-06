// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React, { Component } from 'react'
import debounce from 'debounce'
import { object } from 'prop-types'

import Spinner from '../../components/Spinner'

import ErrorMessage from '../../components/ErrorMessage/ErrorMessage'
import Container from '../../components/Container/Container'
import ServiceDetailPane from '../../components/ServiceDetailPane'
import {
  StyledFilters,
  StyledServicesTableContainer,
} from './ServicesOverviewPage.styles'
import { mapListServicesAPIResponse } from './map-list-services-api-response'

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
      selectedService: null,
    }

    this.handleSearchOnChange = this.handleSearchOnChange.bind(this)
    this.handleSwitchOnChange = this.handleSwitchOnChange.bind(this)
    this.escFunction = this.escFunction.bind(this)
    this.handleOnServiceClicked = this.handleOnServiceClicked.bind(this)
    this.detailPaneCloseHandler = this.detailPaneCloseHandler.bind(this)

    this.searchOnChangeDebouncable = (query) => {
      this.setState({ debouncedQuery: query })
      history.push(`?q=${encodeURIComponent(query)}`)
    }

    this.searchOnChangeDebounced = debounce(this.searchOnChangeDebouncable, 400)
  }

  componentDidMount() {
    document.addEventListener('keydown', this.escFunction, false)

    this.fetchServices()
      .then((response) => mapListServicesAPIResponse(response))
      .then((services) => {
        this.setState({ loading: false, error: false, services })
      })
      .catch(() => {
        this.setState({ loading: false, error: true })
      })
  }

  componentWillUnmount() {
    document.removeEventListener('keydown', this.escFunction, false)
  }

  detailPaneCloseHandler() {
    this.setState({
      selectedService: null,
    })
  }

  fetchServices() {
    return fetch(`/api/directory/list-services`, {
      headers: {
        'Content-Type': 'application/json',
      },
    }).then((response) => response.json())
  }

  escFunction(event) {
    if (event.keyCode === ESCAPE_KEY_CODE) {
      this.setState({ query: '' })
    }
  }

  handleOnServiceClicked(service) {
    this.setState({
      selectedService: service,
    })
  }

  handleSearchOnChange(query) {
    this.setState({ query })
    this.searchOnChangeDebounced(query)
  }

  handleSwitchOnChange(checked) {
    this.setState({ displayOfflineServices: checked })
  }

  render() {
    const {
      displayOfflineServices,
      query,
      debouncedQuery,
      loading,
      error,
      services,
      selectedService,
    } = this.state

    if (loading) {
      return <Spinner />
    }

    if (error) {
      return <ErrorMessage />
    }

    return (
      <Container>
        <StyledFilters
          onQueryChanged={this.handleSearchOnChange}
          onStatusFilterChanged={this.handleSwitchOnChange}
          queryValue={query}
        />

        <StyledServicesTableContainer
          services={services}
          sortBy="organization"
          sortOrder="asc"
          filterQuery={debouncedQuery}
          filterByOnlineServices={!displayOfflineServices}
          onServiceClickedHandler={this.handleOnServiceClicked}
        />

        {selectedService ? (
          <ServiceDetailPane
            organizationName={selectedService.organization}
            contactEmailAddress={selectedService.contactEmailAddress}
            name={selectedService.name}
            closeHandler={this.detailPaneCloseHandler}
          />
        ) : null}
      </Container>
    )
  }
}

ServicesOverviewPage.propTypes = {
  location: object,
  history: object,
}

ServicesOverviewPage.defaultProps = {
  location: window.location,
  history: window.history,
}

export default ServicesOverviewPage
