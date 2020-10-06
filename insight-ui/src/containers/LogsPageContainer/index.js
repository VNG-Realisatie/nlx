// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { Component } from 'react'
import { Route } from 'react-router-dom'
import { arrayOf, instanceOf, shape, string, number, func } from 'prop-types'
import { connect } from 'react-redux'

import { fetchOrganizationLogsRequest } from '../../store/actions'
import LogsPage from '../../components/LogsPage'
import LogDetailPaneContainer from '../LogDetailPaneContainer'

const LOGS_PER_PAGE = 10

export const getPageFromQueryString = (queryString) => {
  const page = new URLSearchParams(queryString).get('page')
  return page !== null ? parseInt(page, 10) : undefined
}

export class LogsPageContainer extends Component {
  constructor(props) {
    super(props)

    this.logClickedHandler = this.logClickedHandler.bind(this)
  }

  componentDidMount() {
    const {
      organization,
      location: { search },
      proof,
    } = this.props

    if (!proof) {
      return
    }

    const page = getPageFromQueryString(search)
    this.fetchOrganizationLogs(organization, proof, page)
  }

  componentDidUpdate(prevProps) {
    const {
      organization,
      location: { search },
      proof,
    } = this.props
    const {
      organization: prevOrganization,
      location: { search: prevSearch },
      proof: prevProof,
    } = prevProps
    const page = getPageFromQueryString(search)
    const prevPage = getPageFromQueryString(prevSearch)

    if (!proof) {
      return
    }

    if (
      organization === prevOrganization &&
      proof === prevProof &&
      page === prevPage
    ) {
      return
    }

    this.fetchOrganizationLogs(organization, proof, page)
  }

  fetchOrganizationLogs(organization, proof, page = 1) {
    if (!organization) {
      return
    }

    if (!proof) {
      return
    }

    // eslint-disable-next-line camelcase
    this.props.fetchOrganizationLogs({
      insight_log_endpoint: organization.insight_log_endpoint, // eslint-disable-line camelcase
      proof,
      page: page - 1, // the API's pages are 0-based
    })
  }

  logClickedHandler(log) {
    const {
      history,
      match: { url },
      location: { search },
    } = this.props
    history.push(`${url}/${log.id}${search}`)
  }

  handleOnPageChanged(page) {
    const {
      location: { search, pathname },
      history,
    } = this.props
    const searchParams = new URLSearchParams(search)
    searchParams.set('page', page)
    history.push(`${pathname}?${searchParams.toString()}`)
  }

  render() {
    const {
      logs,
      organization,
      location: { pathname, search },
      match: { url },
    } = this.props
    const activeLogId = pathname.substr(url.length + 1)
    const currentPage = getPageFromQueryString(search) || 1
    const rowCount = logs.rowCount || 0
    const rowsPerPage = logs.rowsPerPage || 0

    return (
      <>
        <LogsPage
          logs={logs.records}
          currentPage={currentPage}
          rowCount={rowCount}
          rowsPerPage={rowsPerPage}
          onPageChangedHandler={(page) => this.handleOnPageChanged(page)}
          organizationName={organization.name}
          activeLogId={activeLogId}
          logClickedHandler={(log) => this.logClickedHandler(log)}
        />
        <Route
          path={`${url}/:logid/`}
          render={(props) => (
            <LogDetailPaneContainer parentURL={url} {...props} />
          )}
        />
      </>
    )
  }
}

LogsPageContainer.propTypes = {
  match: shape({
    url: string,
  }),
  location: shape({
    pathname: string,
    search: string,
  }),
  history: shape({ push: func.isRequired }).isRequired,
  organization: shape({
    name: string.isRequired,
    insight_log_endpoint: string.isRequired, // eslint-disable-line camelcase
  }).isRequired,
  logs: shape({
    records: arrayOf(
      shape({
        id: string,
        subjects: arrayOf(string),
        requestedBy: string,
        requestedAt: string,
        application: string,
        reason: string,
        date: instanceOf(Date),
      }),
    ),
    rowCount: number,
    rowsPerPage: number,
  }),
  proof: string,
  fetchOrganizationLogs: func.isRequired,
}

LogsPageContainer.defaultProps = {
  match: {
    url: '',
  },
  location: {
    pathname: '',
  },
  logs: {
    records: [],
    rowCount: 0,
    rowsPerPage: 0,
  },
}

const mapStateToProps = ({ logs, proof }) => {
  return {
    logs,
    proof: proof.value,
  }
}

const mapDispatchToProps = (dispatch) => ({
  fetchOrganizationLogs: (
    // eslint-disable-next-line camelcase
    { insight_log_endpoint, proof, page },
  ) =>
    dispatch(
      fetchOrganizationLogsRequest({
        insight_log_endpoint, // eslint-disable-line camelcase
        proof,
        page,
        rowsPerPage: LOGS_PER_PAGE,
      }),
    ),
})

export default connect(mapStateToProps, mapDispatchToProps)(LogsPageContainer)
