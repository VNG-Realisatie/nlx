// Copyright © VNG Realisatie 2018
// Licensed under the EUPL
//

import React, { useEffect, useState } from 'react'
import debounce from '@commonground/design-system/dist/utils/debounce'
import { object } from 'prop-types'
import Spinner from '../../components/Spinner'
import ErrorMessage from '../../components/ErrorMessage'
import { Container } from '../../components/Grid'
import Introduction from '../../components/IntroductionParticipants'
import Section from '../../components/Section'
import News from '../../components/NewsSection'
import Footer from '../../components/Footer'
import ParticipantsTable from '../../components/ParticipantTable'
import FiltersParticipant from '../../components/FiltersParticipant'
import getParticipants from './get-participants'

const ParticipantsPage = ({ location, history }) => {
  const urlParams = new URLSearchParams(location.search)

  const [state, setState] = useState({
    loading: true,
    error: null,
    participants: [],
    query: urlParams.get('q') || '',
    debouncedQuery: urlParams.get('q') || '',
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

  useEffect(() => {
    const loadParticipants = async () => {
      try {
        const participants = await getParticipants()
        setState({
          ...state,
          loading: false,
          error: false,
          participants: participants,
        })
      } catch (e) {
        setState({ ...state, loading: false, error: true })
      }
    }
    loadParticipants()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  const { query, debouncedQuery, loading, error, participants } = state

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
          <FiltersParticipant
            onQueryChanged={handleSearchOnChange}
            queryValue={query}
          />

          <ParticipantsTable
            participants={participants}
            filterQuery={debouncedQuery}
          />
        </Container>
      </Section>

      <News />

      <Footer />
    </>
  )
}

ParticipantsPage.propTypes = {
  location: object,
  history: object,
}

ParticipantsPage.defaultProps = {
  location: window.location,
  history: window.history,
}

export default ParticipantsPage
