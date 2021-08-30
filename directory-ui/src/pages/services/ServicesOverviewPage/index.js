// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React, { useEffect, useState } from 'react'
import debounce from 'debounce'
import { object } from 'prop-types'
import { Route, useParams } from 'react-router-dom'
import Spinner from '../../../components/Spinner'
import ErrorMessage from '../../../components/ErrorMessage'
import { Container } from '../../../components/Grid'
import Introduction from '../../../components/Introduction'
import Section from '../../../components/Section'
import News from '../../../components/NewsSection'
import Footer from '../../../components/Footer'
import DirectoryTable from '../../../components/DirectoryTable'
import Filters from '../../../components/Filters'
import DirectoryDetailPage from '../DirectoryDetailPage'
import { mapListServicesAPIResponse } from './map-list-services-api-response'

const ServicesOverviewPage = ({ location, history }) => {
  const urlParams = new URLSearchParams(location.search)
  const serviceName = useParams()?.serviceName

  const [state, setState] = useState({
    loading: true,
    error: null,
    services: [],
    query: urlParams.get('q') || '',
    debouncedQuery: urlParams.get('q') || '',
    displayOfflineServices: true,
    selectedService: null,
  })

  const searchOnChangeDebouncable = (query) => {
    setState({ ...state, debouncedQuery: query })
    history.push(query ? `?q=${encodeURIComponent(query)}` : '')
  }

  const searchOnChangeDebounced = debounce(searchOnChangeDebouncable, 100)

  const handleSearchOnChange = (query) => {
    setState({ ...state, query })
    searchOnChangeDebounced(query)
  }

  const handleSwitchOnChange = (checked) => {
    setState({ ...state, displayOfflineServices: checked })
  }

  useEffect(() => {
    const getServices = async () => {
      try {
        const response = await fetch(`/api/directory/list-services`, {
          headers: {
            'Content-Type': 'application/json',
          },
        })
        const services = await response.json()
        const renamedServices = mapListServicesAPIResponse(services)
        setState({
          ...state,
          loading: false,
          error: false,
          services: renamedServices,
        })
      } catch (e) {
        setState({ ...state, loading: false, error: true })
      }
    }
    getServices()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  const {
    displayOfflineServices,
    query,
    debouncedQuery,
    loading,
    error,
    services,
  } = state

  if (loading) {
    return <Spinner />
  }

  if (error) {
    return <ErrorMessage />
  }

  return (
    <>
      <Introduction />

      <Section>
        <Container>
          <Filters
            onQueryChanged={handleSearchOnChange}
            onStatusFilterChanged={handleSwitchOnChange}
            queryValue={query}
          />

          <DirectoryTable
            services={services}
            selectedServiceName={serviceName}
            filterQuery={debouncedQuery}
            filterByOnlineServices={!displayOfflineServices}
          />

          <Route
            path="/:organizationName/:serviceName"
            render={() => {
              const selectedService = services.find(
                (service) => service.name === serviceName,
              )
              return (
                <DirectoryDetailPage
                  service={selectedService}
                  parentUrl={
                    state.debouncedQuery && `/?q=${state?.debouncedQuery}`
                  }
                />
              )
            }}
          />
        </Container>
      </Section>

      <News />

      <Footer />
    </>
  )
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
