// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../../../../test-utils'
import ExternalLinkSection from './index'

test('render two links that open in new window', async () => {
  const service = {
    documentationURL: 'https://link.to.somewhere', // this will remove the disabled=true from documentationButton
    organization: {
      name: 'NLX',
      serialNumber: '01234567890123456789',
    },
  }

  const { getByText } = renderWithProviders(
    <ExternalLinkSection service={service} />,
  )

  const documentationButton = getByText('Documentation')

  expect(documentationButton).toHaveTextContent('external-link.svg')
  expect(documentationButton).toHaveAttribute('aria-disabled', 'false')
  expect(documentationButton).toHaveAttribute('target', '_blank')
})

test('render disabled buttons', async () => {
  const service = {
    documentationURL: '',
    organization: {
      name: 'NLX',
      serialNumber: '01234567890123456789',
    },
  }

  const { getByText } = renderWithProviders(
    <ExternalLinkSection service={service} />,
  )

  const documentationButton = getByText('Documentation')

  expect(documentationButton).toHaveAttribute('aria-disabled', 'true')
})
