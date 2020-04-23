// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { Suspense } from 'react'
import { Switch, Route, Redirect } from 'react-router-dom'
import { ThemeProvider } from 'styled-components/macro'
import { GlobalStyles as DSGlobalStyles } from '@commonground/design-system'

import GlobalStyles from './components/GlobalStyles'
import theme from './theme'

import AddServicePage from './pages/AddServicePage'
import EditServicePage from './pages/EditServicePage'
import InwaysPage from './pages/InwaysPage'
import LoginPage from './pages/LoginPage/index'
import NotFoundPage from './pages/NotFoundPage'
import ServicesPage from './pages/ServicesPage'

import { StyledContainer } from './App.styles'

const App = () => (
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
          <Route
            path="/services/:name/edit-service"
            component={EditServicePage}
          />
          <Route path="/services" component={ServicesPage} />
          <Route path="/inways" component={InwaysPage} />

          <Route path="*" component={NotFoundPage} />
        </Switch>
      </Suspense>
    </ThemeProvider>
  </StyledContainer>
)

export default App
