import React from 'react'
import { arrayOf, shape, string, instanceOf } from 'prop-types'
import LogsTable from '../LogsTable'
import ErrorMessage from '../ErrorMessage'
import { StyledLogsPage } from './index.styles'

const LogsPage = ({ logs, organizationName, onSearchQueryChanged }) =>
  logs && logs.length ?
    <StyledLogsPage>
      <LogsTable logs={logs} />
    </StyledLogsPage> :
    <ErrorMessage title="No logs found">
      <p><strong>{organizationName}</strong> has no logs to show, unfortunately.</p>
    </ErrorMessage>

LogsPage.propTypes = {
  logs: arrayOf(shape({
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    reason: string,
    date: instanceOf(Date)
  })),
  organizationName: string.isRequired,
}

LogsPage.defaultProps = {
  logs: [],
}

export default LogsPage
