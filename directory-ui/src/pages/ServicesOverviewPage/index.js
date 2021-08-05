// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React, { useEffect, useState } from 'react'
import debounce from 'debounce'
import { object } from 'prop-types'
import { Route, useParams } from 'react-router-dom'
import Spinner from '../../components/Spinner'
import ErrorMessage from '../../components/ErrorMessage'
import Container from '../../components/Container/Container'
import Introduction from '../../components/Introduction'
import Section from '../../components/Section'
import { Col, Row } from '../../components/Grid'
import News from '../../components/NewsSection'
import Footer from '../../components/Footer'
import DirectoryTable from '../../components/DirectoryTable'
import { StyledFilters, StyledServicesTableContainer } from './index.styles'
import { mapListServicesAPIResponse } from './map-list-services-api-response'
import DirectoryDetailPage from './DirectoryDetailPage'

const ESCAPE_KEY_CODE = 27

const ServicesOverviewPage = ({ location, history }) => {
  const urlParams = new URLSearchParams(location.search)
  const { name } = useParams()

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
    history.push(`?q=${encodeURIComponent(query)}`)
  }

  const searchOnChangeDebounced = debounce(searchOnChangeDebouncable, 400)

  // const detailPaneCloseHandler = () => {
  //   setState({
  //     ...state,
  //     selectedService: null,
  //   })
  // }
  const fetchServices = async () => {
    const response = await fetch(`/api/directory/list-services`, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    return await response.json()
  }

  const escFunction = (event) => {
    if (event.keyCode === ESCAPE_KEY_CODE) {
      setState({ ...state, query: '' })
    }
  }

  // const handleOnServiceClicked = (service) => {
  //   setState({
  //     ...state,
  //     selectedService: service,
  //   })
  // }

  // const handleSearchOnChange = (query) => {
  //   setState({ ...state, query })
  //   searchOnChangeDebounced(query)
  // }

  // const handleSwitchOnChange = (checked) => {
  //   setState({ ...state, displayOfflineServices: checked })
  // }

  useEffect(() => {
    document.addEventListener('keydown', escFunction, false)

    fetchServices()
      .then((response) => mapListServicesAPIResponse(response))
      .then((services) => {
        setState({ ...state, loading: false, error: false, services })
      })
      .catch(() => {
        setState({ ...state, loading: false, error: true })
      })

    return () => {
      document.removeEventListener('keydown', escFunction, false)
    }
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
          <Row>
            <Col>
              <Container>
                {/* TODO: Implement proper search */}
                {/* <StyledFilters
                  onQueryChanged={handleSearchOnChange}
                  onStatusFilterChanged={handleSwitchOnChange}
                  queryValue={query}
                /> */}

                <DirectoryTable
                  services={services}
                  selectedServiceName="setvicename"
                />

                {/* TODO check if routing is configured properly */}
                <Route
                  path="/directory/:organizationName/:serviceName"
                  render={({ match }) => {
                    return (
                      services.length && (
                        <DirectoryDetailPage service={services[0]} />
                      )
                    )
                  }}
                />
              </Container>
            </Col>
          </Row>
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
