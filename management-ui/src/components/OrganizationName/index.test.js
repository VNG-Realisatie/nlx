// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { screen, act, waitFor } from '@testing-library/react'
import deferredPromise from '../../utils/deferred-promise'
import { renderWithProviders } from '../../test-utils'
import OrganizationName from './index'

test('show the organizationName', async () => {
  const environment = deferredPromise()
  const getEnvironment = jest.fn().mockResolvedValue(environment)

  const { container, getByTitle } = renderWithProviders(
    <OrganizationName getEnvironment={getEnvironment} />,
  )

  expect(container).toBeEmptyDOMElement()

  await act(async () => {
    environment.resolve({
      organizationName: 'test',
      organizationSerialNumber: '00000000000000000001',
    })
  })

  expect(container).toHaveTextContent('test')
  expect(() => getByTitle('test')).toThrow()
})

test('adding a title when used in the header', async () => {
  const getEnvironment = jest.fn().mockResolvedValue({
    organizationName: 'test',
    organizationSerialNumber: '00000000000000000001',
  })

  renderWithProviders(
    <OrganizationName getEnvironment={getEnvironment} isHeader />,
  )

  await waitFor(() => screen.findByTitle('test'))

  expect(screen.getByTitle('test')).toHaveTextContent(/test/)
  expect(screen.getByTitle('test')).toHaveTextContent(
    'OIN 00000000000000000001',
  )
})
