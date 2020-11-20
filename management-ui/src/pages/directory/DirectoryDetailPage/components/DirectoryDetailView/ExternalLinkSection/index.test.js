// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { renderWithProviders } from '../../../../../../test-utils'
import ExternalLinkSection from './index'

test('render two links that open in new window', () => {
  const service = {
    documentationURL: 'https://link.to.somewhere',
  }

  const { getByText } = renderWithProviders(
    <ExternalLinkSection service={service} />,
  )

  const documentationButton = getByText('Documentation')
  const specificationButton = getByText('Specification')

  expect(documentationButton).toHaveTextContent('external-link.svg')
  expect(documentationButton).toHaveAttribute('aria-disabled', 'false')
  expect(documentationButton).toHaveAttribute('target', '_blank')
  expect(specificationButton).toHaveTextContent('external-link.svg')
  expect(specificationButton).toHaveAttribute('aria-disabled', 'true')
  expect(specificationButton).toHaveAttribute('target', '_blank')
})

test('render disabled buttons', () => {
  const service = {
    documentationURL: '',
    specificationURL: '',
  }

  const { getByText } = renderWithProviders(
    <ExternalLinkSection service={service} />,
  )

  const documentationButton = getByText('Documentation')
  const specificationButton = getByText('Specification')

  expect(documentationButton).toHaveAttribute('aria-disabled', 'true')
  expect(specificationButton).toHaveAttribute('aria-disabled', 'true')
})
