import React, { Component } from 'react'
import { arrayOf, instanceOf, shape, string } from "prop-types";
import {connect} from 'react-redux'

import { fetchOrganizationLogsRequest } from '../../store/actions'
import LogsPage from '../../components/LogsPage'

export class LogsPageContainer extends Component {
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

  render() {
    const { logs, organization } = this.props
    return <LogsPage logs={logs} organizationName={organization.name} />
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
