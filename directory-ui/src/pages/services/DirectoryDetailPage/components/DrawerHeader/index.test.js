// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Drawer } from '@commonground/design-system'
import { renderWithProviders } from '../../../../../test-utils'
import { SERVICE_STATE_UP } from '../../../../../components/StateIndicator'
import DrawerHeader from './index'

const service = {
  name: 'service',
  organization: 'organisation',
  status: SERVICE_STATE_UP,
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
