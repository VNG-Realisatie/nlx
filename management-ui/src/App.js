// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { Suspense } from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'
import { ThemeProvider } from 'styled-components/macro'
import { GlobalStyles as DSGlobalStyles } from '@commonground/design-system'
import { shape, string } from 'prop-types'
import GlobalStyles from './components/GlobalStyles'
import theme from './theme'

import LoginPage from './pages/LoginPage/index'
import ServicesPage from './pages/ServicesPage'
import AddServicePage from './pages/AddServicePage'
import NotFoundPage from './pages/NotFoundPage'

import { StyledContainer } from './App.styles'
import { UserContextProvider } from './user-context'

const App = ({ user }) => (
  <UserContextProvider user={user}>
    <StyledContainer>
      <ThemeProvider theme={theme}>
        <DSGlobalStyles />
        <GlobalStyles />

        {/* Suspense is required for XHR backend i18next */}
        <Suspense fallback={null}>
          <Switch>
            <Redirect exact path="/" to="/login" />
            <Route path="/login" component={LoginPage} />
            <Route path="/services/add-service" component={AddServicePage} />
            <Route path="/services" component={ServicesPage} />

            <Route path="*" component={NotFoundPage} />
          </Switch>
        </Suspense>
      </ThemeProvider>
    </StyledContainer>
  </UserContextProvider>
)

App.propTypes = {
  user: shape({
    id: string,
    fullName: string,
    email: string,
    pictureUrl: string,
  }),
}

export default App
