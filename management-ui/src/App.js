import React, { Suspense } from 'react'
import { Route, Redirect } from 'react-router-dom'
import { ThemeProvider } from 'styled-components/macro'
import { GlobalStyles as DSGlobalStyles } from '@commonground/design-system'
import GlobalStyles from './components/GlobalStyles'
import theme from './theme'

import LoginPage from './pages/LoginPage/index'
import ServicesPage from './pages/ServicesPage'
import { StyledContainer } from './App.styles'

const App = () => (
  <StyledContainer>
    <ThemeProvider theme={theme}>
      <DSGlobalStyles />
      <GlobalStyles />

      {/* Suspense is required for XHR backend i18next */}
      <Suspense fallback={null}>
        <Route exact path="/">
          <Redirect to="/login" />
        </Route>

        <Route path="/login">
          <LoginPage />
        </Route>

        <Route path="/services">
          <ServicesPage />
        </Route>
      </Suspense>
    </ThemeProvider>
  </StyledContainer>
)

export default App
