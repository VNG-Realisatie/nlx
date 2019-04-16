import React, { Component } from 'react'
import { arrayOf, instanceOf, shape, string } from "prop-types";
import {connect} from 'react-redux'

import { fetchOrganizationLogsRequest } from '../../store/actions'
import LogsPage from '../../components/LogsPage'

const arrayContainsString = (input, query) =>
  input
    .filter(item => item.includes(query.toLowerCase()))
    .length > 0

export class LogsPageContainer extends Component {
  constructor(props) {
    super(props)

    this.state = {
      query: ''
    }

    this.onSearchQueryChanged = this.onSearchQueryChanged.bind(this)
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

  onSearchQueryChanged(query) {
    this.setState({ query })
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

  getFilteredLogsByQuery(logs = [], query = '') {
    return logs
      .filter(log => {
        const { subjects = [], requestedBy = '', requestedAt = '', reason = '' } = log
        return arrayContainsString(subjects, query) ||
          requestedBy.includes(query.toLowerCase()) ||
          requestedAt.includes(query.toLowerCase()) ||
          reason.includes(query.toLowerCase())
      })
  }

  render() {
    const { logs, organization } = this.props
    const { query } = this.state

    return <LogsPage logs={this.getFilteredLogsByQuery(logs, query)}
                     organizationName={organization.name}
                     onSearchQueryChanged={this.onSearchQueryChanged} />
  }
}

LogsPageContainer.propTypes = {
  organization: shape({
    name: string.isRequired,
    insight_log_endpoint: string.isRequired
  }).isRequired,
  loginRequestInfo: shape({
    proofUrl: string
  }),
  logs: arrayOf(shape({
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    reason: string,
    date: instanceOf(Date)
  }))
}

const mapRawLogsToTableFormat = rawLogs =>
  rawLogs.map(rawLog => ({
    subjects: rawLog.data['doelbinding-data-elements'].split(','),
    requestedBy: rawLog['source_organization'],
    requestedAt: rawLog['destination_organization'],
    reason: rawLog.data['doelbinding-process-id'],
    date: new Date(rawLog['created'])
  }))

const mapStateToProps = ({ loginRequestInfo, logs }) => {
  return {
    loginRequestInfo,
    logs: mapRawLogsToTableFormat(logs)
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
