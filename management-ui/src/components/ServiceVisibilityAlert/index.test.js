// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { renderWithProviders } from '../../test-utils'
import {
  ServiceVisibilityAlert,
  showServiceVisibilityAlert,
  ServiceVisibilityMessage,
} from './index'

const internalWithoutInways = { internal: true, inways: [] }
const internalWithInways = { internal: true, inways: ['inway'] }
const publishedWithoutInways = { internal: false, inways: [] }
const publishedWithInways = { internal: false, inways: ['inway'] }

test('only render when the service is  published without inways', () => {
  expect(showServiceVisibilityAlert(internalWithoutInways)).toEqual(false)
  expect(showServiceVisibilityAlert(internalWithInways)).toEqual(false)
  expect(showServiceVisibilityAlert(publishedWithInways)).toEqual(false)
  expect(showServiceVisibilityAlert(publishedWithoutInways)).toEqual(true)
})

test('ServiceVisibilityMessage', () => {
  const { container, getByRole } = renderWithProviders(
    <ServiceVisibilityMessage />,
  )

  expect(() => getByRole('alert')).not.toThrow()
  expect(container).toHaveTextContent(/^Service not yet accessible$/)
})

test('ServiceVisibilityAlert', () => {
  const { getByRole, getByTestId } = renderWithProviders(
    <ServiceVisibilityAlert />,
  )

  expect(() => getByRole('alert')).not.toThrow()
  expect(getByTestId('title')).toHaveTextContent(/^Service not yet accessible$/)
  expect(getByTestId('content')).toHaveTextContent(
    /^There are no inways connected yet. Until then other organizations cannot access this service.$/,
  )
})
