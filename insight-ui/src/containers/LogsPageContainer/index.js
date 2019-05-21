import React, { Component, Fragment } from 'react'
import { Route } from 'react-router-dom'
import { arrayOf, instanceOf, shape, string } from "prop-types";
import {connect} from 'react-redux'

import { fetchOrganizationLogsRequest } from '../../store/actions'
import LogsPage from '../../components/LogsPage'
import LogDetailPaneContainer from '../LogDetailPaneContainer'

export class LogsPageContainer extends Component {
  constructor(props) {
    super(props)

    this.logClickedHandler = this.logClickedHandler.bind(this)
  }

  fetchOrganizationLogs(organization, loginRequestInfo) {
    if (!organization || !loginRequestInfo) {
      return
    }

    this.props.fetchOrganizationLogs({
      proofUrl: loginRequestInfo.proofUrl,
      insight_log_endpoint: organization.insight_log_endpoint
    });
  }

  componentWillReceiveProps(nextProps) {
    const { organization, loginRequestInfo } = nextProps
    const { organization: prevOrganization } = this.props

    if (organization === prevOrganization) {
      return
    }

    this.fetchOrganizationLogs(organization, loginRequestInfo)
  }

  componentDidMount() {
    const { organization, loginRequestInfo } = this.props

    if (!loginRequestInfo) {
      return
    }

    this.fetchOrganizationLogs(organization, loginRequestInfo)
  }

  logClickedHandler(log) {
    const { history, match: { url } } = this.props
    history.push(`${url}/${log.id}`)
  }

  render() {
    const { logs, organization } = this.props
    const { match: { url } } = this.props
    const { pathname } = this.props.location
    const activeLogId = pathname.substr(url.length + 1)

    return (
      <Fragment>
        <LogsPage logs={logs} organizationName={organization.name} activeLogId={activeLogId} logClickedHandler={log => this.logClickedHandler(log)} />
        <Route path={`${url}/:logid/`} render={props => <LogDetailPaneContainer parentURL={url} {...props} />} />
      </Fragment>
    )
  }
}

LogsPageContainer.propTypes = {
  match: shape({
    url: string
  }),
  location: shape({
    pathname: string
  }),
  organization: shape({
    name: string.isRequired,
    insight_log_endpoint: string.isRequired
  }).isRequired,
  loginRequestInfo: shape({
    proofUrl: string
  }),
  logs: arrayOf(shape({
    id: string,
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    application: string,
    reason: string,
    date: instanceOf(Date)
  }))
}

LogsPageContainer.defaultProps = {
  match: {
    url: ''
  },
  location: {
    pathname: ''
  }
}

const mapStateToProps = ({ loginRequestInfo, logs }) => {
  return {
    loginRequestInfo,
    logs
  }
}

const mapDispatchToProps = dispatch => ({
  fetchOrganizationLogs: ({ insight_log_endpoint, proofUrl }) =>
    dispatch(fetchOrganizationLogsRequest({ insight_log_endpoint, proofUrl }))
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(LogsPageContainer)
