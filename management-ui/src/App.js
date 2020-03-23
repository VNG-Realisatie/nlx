import React, { Suspense } from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'
import { ThemeProvider } from 'styled-components/macro'
import { GlobalStyles as DSGlobalStyles } from '@commonground/design-system'
import GlobalStyles from './components/GlobalStyles'
import theme from './theme'

import LoginPage from './pages/LoginPage/index'
import ServicesPage from './pages/ServicesPage'
import NotFoundPage from './pages/NotFoundPage'

import { StyledContainer } from './App.styles'
import { UserContextProvider } from './user-context/UserContext'

const App = () => (
  <UserContextProvider>
    <StyledContainer>
      <ThemeProvider theme={theme}>
        <DSGlobalStyles />
        <GlobalStyles />

        {/* Suspense is required for XHR backend i18next */}
        <Suspense fallback={null}>
          <Switch>
            <Redirect exact path="/" to="/login" />
            <Route path="/login" component={LoginPage} />
            <Route path="/services" component={ServicesPage} />

            <Route path="*" component={NotFoundPage} />
          </Switch>
        </Suspense>
      </ThemeProvider>
    </StyledContainer>
  </UserContextProvider>
)

export default App
