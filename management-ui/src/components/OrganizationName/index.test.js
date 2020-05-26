// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import deferredPromise from '../../test-utils/deferred-promise'
import { act, renderWithProviders } from '../../test-utils'
import OrganizationName from './index'

test('show the organizationName', async () => {
  const environment = deferredPromise()
  const getEnvironment = jest.fn(() => environment)

  const { container, getByTitle } = renderWithProviders(
    <OrganizationName getEnvironment={getEnvironment} />,
  )

  expect(container).toBeEmpty()

  await act(async () => {
    environment.resolve({ organizationName: 'test' })
  })

  expect(container).toHaveTextContent('test')
  expect(() => getByTitle('test')).toThrow()
})

test('adding a title when used in the header', async () => {
  const getEnvironment = jest.fn(() => ({ organizationName: 'test' }))

  const { findByTitle } = renderWithProviders(
    <OrganizationName getEnvironment={getEnvironment} isHeader />,
  )

  expect(await findByTitle('test')).toHaveTextContent('test')
})
