// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { MemoryRouter as Router } from 'react-router-dom'
import { RootStore, StoreProvider } from '../../stores'

import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context'
import PageTemplate from './index'

jest.mock('../OrganizationName', () => () => <span>test</span>)
jest.mock('./OrganizationInwayCheck', () => () => null)

test('PageTemplate', () => {
  const rootStore = new RootStore({})

  const { getByText } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Router>
        <PageTemplate>
          <p>Page content</p>
        </PageTemplate>
      </Router>
    </StoreProvider>,
  )

  expect(getByText(/^Page content$/)).toBeInTheDocument()
})

test('PageTemplate with Header', () => {
  const rootStore = new RootStore({})

  const { getByText, getByTestId } = renderWithProviders(
    <StoreProvider rootStore={rootStore}>
      <Router>
        <UserContextProvider user={{}}>
          <PageTemplate>
            <PageTemplate.Header
              title="Page title"
              description="Page description"
            />
            <p>Page content</p>
          </PageTemplate>
        </UserContextProvider>
      </Router>
    </StoreProvider>,
  )

  expect(getByText(/^Page title$/)).toBeInTheDocument()
  expect(getByText(/^Page description$/)).toBeInTheDocument()
  expect(getByText(/^Page content$/)).toBeInTheDocument()
  expect(getByTestId('user-navigation')).toBeInTheDocument()
})
