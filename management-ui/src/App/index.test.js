// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { render } from '@testing-library/react'
import React from 'react'
import { I18nextProvider } from 'react-i18next'
import { RootStore, StoreProvider } from '../stores'
import { UserContextProvider } from '../user-context'
import i18n from '../test-utils/i18nTestConfig'
import App from './index'

test('initializing the App when user is authenticated', async () => {
  const rootStore = new RootStore()
  const applicationStore = rootStore.applicationStore

  applicationStore.getGeneralSettings = jest
    .fn()
    .mockResolvedValue({ organizationInway: 'inway' })

  expect(applicationStore.isOrganizationInwaySet).toBeNull()

  const { findByText } = render(
    <I18nextProvider i18n={i18n}>
      <StoreProvider rootStore={rootStore}>
        <UserContextProvider user={{}}>
          <App>My App</App>
          <div id="root" />
        </UserContextProvider>
      </StoreProvider>
    </I18nextProvider>,
  )

  const welcomeMessage = await findByText('My App')
  expect(welcomeMessage).toBeInTheDocument()
  expect(applicationStore.isOrganizationInwaySet).toBe(true)
})
