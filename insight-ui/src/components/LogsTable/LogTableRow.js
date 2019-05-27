// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { arrayOf, string, instanceOf } from "prop-types";

import { Table } from '@commonground/design-system'
import { StyledLogTableRow, StyledSubjectLabel } from './LogTableRow.styles'

const dateOptions = {
  day: '2-digit',
  month: '2-digit',
  year: 'numeric'
}

const LogTableRow = ({ subjects, requestedBy, requestedAt, reason, date, ...props }) =>
    <StyledLogTableRow {...props}>
      <Table.BodyCell>
        {
          subjects.map((subject, i) =>
            <StyledSubjectLabel key={i}>{subject}</StyledSubjectLabel>)
        }
      </Table.BodyCell>
      <Table.BodyCell>{ requestedBy }</Table.BodyCell>
      <Table.BodyCell>{ requestedAt }</Table.BodyCell>
      <Table.BodyCell>{ reason }</Table.BodyCell>
      <Table.BodyCell align="right">{ date.toLocaleDateString('nl-nl', dateOptions) }</Table.BodyCell>
    </StyledLogTableRow>

LogTableRow.propTypes = {
  subjects: arrayOf(string),
  requestedBy: string,
  requestedAt: string,
  reason: string,
  date: instanceOf(Date),
}

export default LogTableRow
