import React, { Component, Fragment } from 'react'
import { Route } from 'react-router-dom'
import { arrayOf, instanceOf, shape, string, number } from "prop-types";
import {connect} from 'react-redux'

import { fetchOrganizationLogsRequest } from '../../store/actions'
import LogsPage from '../../components/LogsPage'
import LogDetailPaneContainer from '../LogDetailPaneContainer'

const LOGS_PER_PAGE = 20

const getPageFromQueryString = queryString => {
  const page = new URLSearchParams(queryString).get('page')
  return page !== null ? parseInt(page, 10) : undefined
}

export class LogsPageContainer extends Component {
  constructor(props) {
    super(props)

    this.logClickedHandler = this.logClickedHandler.bind(this)
  }

  fetchOrganizationLogs(organization, loginRequestInfo, page = 1) {
    if (!organization || !loginRequestInfo) {
      return
    }

    this.props.fetchOrganizationLogs({
      proofUrl: loginRequestInfo.proofUrl,
      insight_log_endpoint: organization.insight_log_endpoint,
      page
    });
  }

  componentWillReceiveProps(nextProps) {
    const { organization, loginRequestInfo, location: { prevSearch } } = nextProps
    const { organization: prevOrganization, location: { search } } = this.props
    const page = getPageFromQueryString(search)
    const prevPage = getPageFromQueryString(prevSearch)

    if (organization === prevOrganization && page === prevPage) {
      return
    }

    this.fetchOrganizationLogs(organization, loginRequestInfo, page)
  }

  componentDidMount() {
    const { organization, loginRequestInfo, location: { search } } = this.props

    if (!loginRequestInfo) {
      return
    }

    const page = getPageFromQueryString(search)
    this.fetchOrganizationLogs(organization, loginRequestInfo, page)
  }

  logClickedHandler(log) {
    const { history, match: { url } } = this.props
    history.push(`${url}/${log.id}`)
  }

  onPageChangedHandler(page) {
    const { location: { search, pathname }, history } = this.props
    const searchParams = new URLSearchParams(search)
    searchParams.set('page', page)
    history.push(`${pathname}?${searchParams.toString()}`)
  }

  render() {
    const { logs, organization, location: { pathname, search }, match: { url } } = this.props
    const activeLogId = pathname.substr(url.length + 1)
    const currentPage = getPageFromQueryString(search) || 1
    const amountOfPages = logs.pageCount || 1

    return (
      <Fragment>
        <LogsPage 
          logs={logs.records} 
          currentPage={currentPage} 
          amountOfPages={amountOfPages} 
          onPageChangedHandler={page => this.onPageChangedHandler(page)} 
          organizationName={organization.name} 
          activeLogId={activeLogId} 
          logClickedHandler={log => this.logClickedHandler(log)} 
        />
        <Route 
          path={`${url}/:logid/`} 
          render={props => <LogDetailPaneContainer parentURL={url} {...props} />}
        />
      </Fragment>
    )
  }
}

LogsPageContainer.propTypes = {
  match: shape({
    url: string
  }),
  location: shape({
    pathname: string,
    search: string,
  }),
  organization: shape({
    name: string.isRequired,
    insight_log_endpoint: string.isRequired
  }).isRequired,
  loginRequestInfo: shape({
    proofUrl: string
  }),
  logs: shape({
    records: arrayOf(shape({
      id: string,
      subjects: arrayOf(string),
      requestedBy: string,
      requestedAt: string,
      application: string,
      reason: string,
      date: instanceOf(Date)
    })),
    pageCount: number
  })
}

LogsPageContainer.defaultProps = {
  match: {
    url: ''
  },
  location: {
    pathname: ''
  },
  logs: {
    records: [],
    pageCount: 0
  }
}

const mapStateToProps = ({ loginRequestInfo, logs }) => {
  return {
    loginRequestInfo,
    logs
  }
}

const mapDispatchToProps = dispatch => ({
  fetchOrganizationLogs: ({ insight_log_endpoint, proofUrl, page }) =>
    dispatch(fetchOrganizationLogsRequest({ insight_log_endpoint, proofUrl, page, rowsPerPage: LOGS_PER_PAGE }))
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(LogsPageContainer)
