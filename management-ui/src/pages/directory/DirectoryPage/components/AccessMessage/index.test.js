// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../../../test-utils'
import {
  SHOW_HAS_ACCESS,
  SHOW_REQUEST_CREATED,
  SHOW_REQUEST_FAILED,
  SHOW_REQUEST_RECEIVED,
  SHOW_REQUEST_REJECTED,
} from '../../../directoryServiceAccessState'
import AccessMessage from './index'

test('Correctly renders the access request states', () => {
  const { getByText, getByTitle, rerender } = renderWithProviders(
    <AccessMessage displayState={SHOW_REQUEST_CREATED} />,
  )
  expect(getByText('Sending request')).toBeInTheDocument()

  rerender(<AccessMessage displayState={SHOW_REQUEST_FAILED} />)
  expect(getByText('Request could not be sent')).toBeInTheDocument()

  rerender(<AccessMessage displayState={SHOW_REQUEST_RECEIVED} />)
  expect(getByText('Requested')).toBeInTheDocument()

  rerender(<AccessMessage displayState={SHOW_REQUEST_REJECTED} />)
  expect(getByText('Rejected')).toBeInTheDocument()

  rerender(<AccessMessage displayState={SHOW_HAS_ACCESS} />)
  expect(getByText('check.svg')).toBeInTheDocument()
  expect(getByTitle('Approved')).toBeInTheDocument()
})
