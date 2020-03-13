// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { func, string, arrayOf, instanceOf } from 'prop-types'
import CloseIcon from '../CloseIcon'
import {
  StyledLogDetailPane,
  StyledHeader,
  StyledTitle,
  StyledSubtitle,
  StyledDl,
  StyledCloseButton,
} from './index.styles'

const dateOptions = {
  day: '2-digit',
  month: '2-digit',
  year: 'numeric',
}

const timeOptions = {
  hour: '2-digit',
  minute: '2-digit',
  second: '2-digit',
}

const LogDetailPane = ({
  id,
  subjects,
  requestedBy,
  requestedAt,
  application,
  reason,
  date,
  closeHandler,
}) => (
  <StyledLogDetailPane>
    <StyledHeader>
      <StyledTitle>Log</StyledTitle>
      <StyledCloseButton onClick={() => closeHandler()}>
        <CloseIcon />
      </StyledCloseButton>
    </StyledHeader>

    <StyledSubtitle>Requested</StyledSubtitle>
    <StyledDl>
      <dt>Data</dt>
      <dd>{subjects.join(', ')}</dd>

      <dt>By</dt>
      <dd>{requestedBy}</dd>

      <dt>At</dt>
      <dd>{requestedAt}</dd>

      <dt>Process</dt>
      <dd>{reason}</dd>

      <dt>Application</dt>
      <dd>{application}</dd>
    </StyledDl>

    <StyledSubtitle>Details</StyledSubtitle>
    <StyledDl>
      <dt>ID</dt>
      <dd>{id}</dd>

      <dt>Date</dt>
      <dd>{date.toLocaleDateString('nl-nl', dateOptions)}</dd>

      <dt>Time</dt>
      <dd>{date.toLocaleTimeString('nl-nl', timeOptions)}</dd>
    </StyledDl>
  </StyledLogDetailPane>
)

LogDetailPane.propTypes = {
  closeHandler: func,
  id: string,
  subjects: arrayOf(string).isRequired,
  requestedBy: string.isRequired,
  requestedAt: string,
  application: string,
  reason: string,
  date: instanceOf(Date),
}

LogDetailPane.defaultProps = {
  closeHandler: () => {},
}

export default LogDetailPane
