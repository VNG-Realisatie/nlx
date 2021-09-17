// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Drawer } from '@commonground/design-system'
import { renderWithProviders } from '../../../../../test-utils'
import DrawerHeader from './index'

const service = {
  serviceName: 'service',
  organizationName: 'organisation',
  state: 'up',
  apiSpecificationType: 'OpenAPI',
  serialNumber: '00000000000000000000',
}

const closeHandler = jest.fn()

test('renders without crashing', () => {
  expect(() =>
    renderWithProviders(
      <Drawer noMask closeHandler={closeHandler}>
        <DrawerHeader service={service} />
      </Drawer>,
    ),
  ).not.toThrow()
})

test('apiSpecificationType is not required', () => {
  const serviceWithouthApiSpec = { ...service }
  delete serviceWithouthApiSpec.apiSpecificationType

  const { queryByText } = renderWithProviders(
    <Drawer noMask closeHandler={closeHandler}>
      <DrawerHeader service={serviceWithouthApiSpec} />
    </Drawer>,
  )

  expect(queryByText('OpenAPI')).toBeNull()
})

test('serial number', () => {
  const { queryByText } = renderWithProviders(
    <Drawer noMask closeHandler={closeHandler}>
      <DrawerHeader service={service} />
    </Drawer>,
  )

  expect(queryByText('Serial Number serialNumber')).toBeInTheDocument()
})
