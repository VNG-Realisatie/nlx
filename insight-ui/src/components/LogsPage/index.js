// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { arrayOf, shape, string, instanceOf, func } from 'prop-types'
import LogsTable from '../LogsTable'
import ErrorMessage from '../ErrorMessage'
import { StyledLogsPage, StyledPagination } from './index.styles'

const LogsPage = ({ logs, organizationName, currentPage, amountOfPages, onPageChangedHandler, activeLogId, logClickedHandler }) =>
  logs && logs.length ?
    <StyledLogsPage>
      <LogsTable logs={logs} activeLogId={activeLogId} logClickedHandler={logClickedHandler} />
      <StyledPagination currentPage={currentPage} amountOfPages={amountOfPages} onPageChangedHandler={onPageChangedHandler} />
    </StyledLogsPage> :
    <ErrorMessage title="No logs found">
      <p><strong>{organizationName}</strong> has no logs to show, unfortunately.</p>
    </ErrorMessage>

LogsPage.propTypes = {
  activeLogId: string,
  logClickedHandler: func,
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
  logClickedHandler: () => {}
}

export default LogsPage
