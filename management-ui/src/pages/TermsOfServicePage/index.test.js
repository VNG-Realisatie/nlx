// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import React from 'react'
import { screen } from '@testing-library/react'
import { renderWithProviders } from '../../test-utils'
import TermsOfServicePage from './index'

test('TermsOfService page', async () => {
  renderWithProviders(<TermsOfServicePage />)

  expect(await screen.findByText(/^Terms of Service$/)).toBeInTheDocument()
})
