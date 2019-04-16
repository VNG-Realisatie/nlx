import React from 'react'
import { arrayOf, shape, string, instanceOf, func } from 'prop-types'
import LogsTable from '../LogsTable'
import ErrorMessage from '../ErrorMessage'
import { StyledLogsPage, StyledSearch } from './index.styles'

const LogsPage = ({ logs, organizationName, onSearchQueryChanged }) =>
  logs && logs.length ?
    <StyledLogsPage>
      <StyledSearch placeholder="Filter logs…" onQueryChanged={onSearchQueryChanged} />
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
  onSearchQueryChanged: func
}

LogsPage.defaultProps = {
  logs: [],
  onSearchQueryChanged: () => {}
}

export default LogsPage
