// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'

import { act } from '@testing-library/react'
import { renderWithProviders } from '../../../../../../test-utils'
import deferredPromise from '../../../../../../utils/deferred-promise'
import ExternalLinkSection from './index'

test('render two links that open in new window', async () => {
  const environment = deferredPromise()
  const getEnvironment = jest.fn().mockResolvedValue(environment)

  const service = {
    documentationURL: 'https://link.to.somewhere', // this will remove the disabled=true from documentationButton
    apiSpecificationType: 'OpenAPI2', // this will remove the disabled=true from specificationButton
    organization: {
      name: 'NLX',
      serialNumber: '01234567890123456789',
    },
  }

  const { container, getByText } = renderWithProviders(
    <ExternalLinkSection service={service} getEnv={getEnvironment} />,
  )

  expect(container).toBeEmptyDOMElement()

  await act(async () => {
    environment.resolve({
      organizationName: 'test',
      organizationSerialNumber: '00000000000000000001',
    })
  })

  const documentationButton = getByText('Documentation')
  const specificationButton = getByText('Specification')

  expect(documentationButton).toHaveTextContent('external-link.svg')
  expect(documentationButton).toHaveAttribute('aria-disabled', 'false')
  expect(documentationButton).toHaveAttribute('target', '_blank')
  expect(specificationButton).toHaveTextContent('external-link.svg')
  expect(specificationButton).toHaveAttribute('aria-disabled', 'false')
  expect(specificationButton).toHaveAttribute('target', '_blank')
})

test('render disabled buttons', async () => {
  const environment = deferredPromise()
  const getEnvironment = jest.fn().mockResolvedValue(environment)

  const service = {
    documentationURL: '',
    specificationURL: '',
    organization: {
      name: 'NLX',
      serialNumber: '01234567890123456789',
    },
  }

  const { container, getByText } = renderWithProviders(
    <ExternalLinkSection service={service} getEnv={getEnvironment} />,
  )

  expect(container).toBeEmptyDOMElement()

  await act(async () => {
    environment.resolve({
      organizationName: 'test',
      organizationSerialNumber: '00000000000000000001',
    })
  })

  const documentationButton = getByText('Documentation')
  const specificationButton = getByText('Specification')

  expect(documentationButton).toHaveAttribute('aria-disabled', 'true')
  expect(specificationButton).toHaveAttribute('aria-disabled', 'true')
})
