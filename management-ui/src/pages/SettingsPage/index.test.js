// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { unstable_HistoryRouter as HistoryRouter } from 'react-router-dom'
import { createMemoryHistory } from 'history'
import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import { RootStore, StoreProvider } from '../../stores'
import SettingsPage from './index'

jest.mock('../../components/OrganizationName', () => () => 'organization')
jest.mock(
  '../../components/PageTemplate/OrganizationInwayCheck',
  () => () => null,
)
jest.mock(
  '../../components/PageTemplate/OrganizationEmailAddressCheck',
  () => () => null,
)

jest.mock('./GeneralSettings', () => () => (
  <div data-testid="general-settings" />
))

test('redirects to /settings/general when navigating to /settings', async () => {
  const rootStore = new RootStore({})
  const history = createMemoryHistory()

  renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <HistoryRouter history={history}>
        <UserContextProvider user={{ id: '42' }}>
          <SettingsPage />
        </UserContextProvider>
      </HistoryRouter>
    </StoreProvider>,
  )

  expect(history.location.pathname).toEqual('/general')
})

test('the /general route renders the General settings', () => {
  const history = createMemoryHistory({ initialEntries: ['/general'] })
  const rootStore = new RootStore({})

  const { getByTestId } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <HistoryRouter history={history}>
        <UserContextProvider user={{ id: '42' }}>
          <SettingsPage />
        </UserContextProvider>
      </HistoryRouter>
    </StoreProvider>,
  )
  expect(getByTestId('general-settings')).toBeInTheDocument()
})
