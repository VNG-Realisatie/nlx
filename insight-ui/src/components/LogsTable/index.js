// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { string, arrayOf, shape, instanceOf, func } from 'prop-types'
import Table from '../Table'
import LogTableRow from './LogTableRow'

const LogsTable = ({ logs, activeLogId, logClickedHandler, ...props }) =>
  <Table {...props}>
    <Table.Head>
      <Table.Row>
        <Table.HeadCell>Requested</Table.HeadCell>
        <Table.HeadCell>Requested by</Table.HeadCell>
        <Table.HeadCell>Requested at</Table.HeadCell>
        <Table.HeadCell>Reason</Table.HeadCell>
        <Table.HeadCell>Datum</Table.HeadCell>
      </Table.Row>
    </Table.Head>

    <Table.Body>
      {
        logs
          .map((log, i) =>
            <LogTableRow key={i}
                         subjects={log.subjects}
                         requestedBy={log.requestedBy}
                         requestedAt={log.requestedAt}
                         reason={log.reason}
                         date={log.date}
                         active={log.id === activeLogId}
                         onClick={() => logClickedHandler(log)}
            />
          )
      }
    </Table.Body>
  </Table>

LogsTable.propTypes = {
  activeLogId: string,
  logClickedHandler: func,
  logs: arrayOf(shape({
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    reason: string,
    date: instanceOf(Date)
  }))
}

LogsTable.defaultProps = {
  logs: [],
  logClickedHandler: () => {}
}

export default LogsTable
