import React from 'react'
import { ThemeProvider } from 'styled-components/macro'
import { defaultTheme } from '@commonground/design-system'
import { renderWithProviders } from '../../test-utils'
import { UserContextProvider } from '../../user-context/UserContext'
import LoginPage from './index'

test('renders a welcome message', () => {
  const { getByText } = renderWithProviders(
    <ThemeProvider theme={defaultTheme}>
      <UserContextProvider>
        <LoginPage />
      </UserContextProvider>
    </ThemeProvider>,
  )
  expect(getByText(/^Welcome$/)).toBeInTheDocument()
})
