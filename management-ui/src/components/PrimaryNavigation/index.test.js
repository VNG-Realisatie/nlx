// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { renderWithProviders } from '../../test-utils'
import PrimaryNavigation from './index'

test('PrimaryNavigation', () => {
  const { getByLabelText } = renderWithProviders(
    <Router>
      <PrimaryNavigation />
    </Router>,
  )

  const linkHome = getByLabelText('Homepage')
  expect(linkHome.getAttribute('href')).toBe('/')

  const linkInwaysAndOutways = getByLabelText('Inways and Outways page')
  expect(linkInwaysAndOutways.getAttribute('href')).toBe('/inways-and-outways')

  const linkServices = getByLabelText('Services page')
  expect(linkServices.getAttribute('href')).toBe('/services')

  const linkDirectory = getByLabelText('Directory page')
  expect(linkDirectory.getAttribute('href')).toBe('/directory')

  const linkFinances = getByLabelText('Finances page')
  expect(linkFinances.getAttribute('href')).toBe('/finances')

  const linkAuditLog = getByLabelText('Audit log page')
  expect(linkAuditLog.getAttribute('href')).toBe('/audit-log')

  const linkSettings = getByLabelText('Settings page')
  expect(linkSettings.getAttribute('href')).toBe('/settings')

  const linkOrders = getByLabelText('Orders page')
  expect(linkOrders.getAttribute('href')).toBe('/orders')
})
