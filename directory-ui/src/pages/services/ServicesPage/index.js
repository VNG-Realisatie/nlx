// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React, { useEffect, useState } from 'react'
import debounce from '@commonground/design-system/dist/utils/debounce'
import { object } from 'prop-types'
import { Route, Routes, useNavigate, useParams } from 'react-router-dom'
import Spinner from '../../../components/Spinner'
import ErrorMessage from '../../../components/ErrorMessage'
import { Container } from '../../../components/Grid'
import Introduction from '../../../components/Introduction'
import Section from '../../../components/Section'
import News from '../../../components/NewsSection'
import Footer from '../../../components/Footer'
import DirectoryTable from '../../../components/DirectoryTable'
import Filters from '../../../components/Filters'
import DirectoryDetailPage from '../ServiceDetailPage'
import getServices from './get-services'

const ServicesPage = ({ location }) => {
  const navigate = useNavigate()
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
    navigate(query ? `?q=${encodeURIComponent(query)}` : '')
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
    const loadServices = async () => {
      try {
        const services = await getServices()
        setState({
          ...state,
          loading: false,
          error: false,
          services: services,
        })
      } catch (e) {
        setState({ ...state, loading: false, error: true })
      }
    }
    loadServices()
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

          <Routes>
            <Route
              path="/:organizationSerialNumber/:serviceName"
              element={<DirectoryDetailPage />}
            />
          </Routes>
        </Container>
      </Section>

      <News />

      <Footer />
    </>
  )
}

ServicesPage.propTypes = {
  location: object,
  history: object,
}

ServicesPage.defaultProps = {
  location: window.location,
  history: window.history,
}

export default ServicesPage
