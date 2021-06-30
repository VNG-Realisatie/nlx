// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { RootStore, StoreProvider } from '../../../stores'

import { renderWithProviders } from '../../../test-utils'
import { UserContextProvider } from '../../../user-context'
import Header from './index'

jest.mock('../../OrganizationName', () => () => <span>test</span>)

test('Header with page elements', () => {
  const rootStore = new RootStore({})

  const { getByText, getByTestId } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Router>
        <UserContextProvider user={{}}>
          <Header title="Page title" description="Page description" />
        </UserContextProvider>
      </Router>
    </StoreProvider>,
  )

  expect(getByText(/^Page title$/)).toBeInTheDocument()
  expect(getByText(/^Page description$/)).toBeInTheDocument()
  expect(getByTestId('user-navigation')).toBeInTheDocument()
})
