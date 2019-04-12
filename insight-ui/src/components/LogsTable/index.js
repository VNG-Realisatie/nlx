import React from 'react'
import { string, arrayOf, shape, instanceOf } from 'prop-types'
import { Table } from '@commonground/design-system'
import LogTableRow from './LogTableRow'

const LogsTable = ({ logs, ...props }) =>
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
            />
          )
      }
    </Table.Body>
  </Table>

LogsTable.propTypes = {
  logs: arrayOf(shape({
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    reason: string,
    date: instanceOf(Date)
  }))
}

LogsTable.defaultProps = {
  logs: []
}

export default LogsTable
