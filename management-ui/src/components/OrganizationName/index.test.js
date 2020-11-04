// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { screen, act, waitFor } from '@testing-library/react'
import deferredPromise from '../../test-utils/deferred-promise'
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
    environment.resolve({ organizationName: 'test' })
  })

  expect(container).toHaveTextContent('test')
  expect(() => getByTitle('test')).toThrow()
})

test('adding a title when used in the header', async () => {
  const getEnvironment = jest
    .fn()
    .mockResolvedValue({ organizationName: 'test' })

  renderWithProviders(
    <OrganizationName getEnvironment={getEnvironment} isHeader />,
  )

  await waitFor(() => screen.findByTitle('test'))

  expect(screen.getByTitle('test')).toHaveTextContent('test')
})
