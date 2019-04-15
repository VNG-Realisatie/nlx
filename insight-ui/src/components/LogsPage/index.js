import React from 'react'
import { arrayOf, shape, string, instanceOf } from 'prop-types'
import LogsTable from "../LogsTable";

const LogsPage = ({ logs }) =>
  <div>
    <LogsTable logs={logs} />
  </div>

LogsPage.propTypes = {
  logs: arrayOf(shape({
    subjects: arrayOf(string),
    requestedBy: string,
    requestedAt: string,
    reason: string,
    date: instanceOf(Date)
  }))
}

LogsPage.defaultProps = {
  logs: []
}

export default LogsPage
