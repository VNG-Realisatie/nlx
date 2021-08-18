// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../../../../test-utils'
import ContactSection from './index'

test('renders without crashing', () => {
  const service = {
    contactEmailAddress: 'mail@service.io',
  }

  const { getByText, rerender } = renderWithProviders(
    <ContactSection service={{}} />,
  )

  expect(getByText('Geen contactgegevens beschikbaar')).toBeInTheDocument()

  rerender(<ContactSection service={service} />)

  expect(getByText('mail@service.io')).toBeInTheDocument()
})
