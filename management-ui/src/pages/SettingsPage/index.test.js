// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Router } from 'react-router-dom'
import { createMemoryHistory } from 'history'

import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import SettingsPage from './index'

jest.mock('../../components/OrganizationName', () => () => 'organization')
jest.mock(
  '../../components/PageTemplate/OrganizationInwayCheck',
  () => () => null,
)

jest.mock('./GeneralSettings', () => () => (
  <div data-testid="general-settings" />
))

test('redirects to /settings/general when navigating to /settings', async () => {
  const history = createMemoryHistory({ initialEntries: ['/settings'] })

  renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <SettingsPage />
      </UserContextProvider>
    </Router>,
  )

  expect(history.location.pathname).toEqual('/settings/general')
})

test('the /settings/general route renders the General settings', () => {
  const history = createMemoryHistory({ initialEntries: ['/settings/general'] })

  const { getByTestId } = renderWithProviders(
    <Router history={history}>
      <UserContextProvider user={{ id: '42' }}>
        <SettingsPage />
      </UserContextProvider>
    </Router>,
  )
  expect(getByTestId('general-settings')).toBeInTheDocument()
})
