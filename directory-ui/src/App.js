// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Switch, BrowserRouter as Router, Route } from 'react-router-dom'
import { ThemeProvider } from 'styled-components'
import { GlobalStyles, DomainNavigation } from '@commonground/design-system'
import VersionLogger from './components/VersionLogger'
import Header from './components/Header'
import ServiceOverviewPage from './pages/ServicesOverviewPage'
import DocumentationPage from './pages/DocumentationPage'
import theme from './theme'
import '@fontsource/source-sans-pro/latin.css'

const App = () => (
  <>
    <ThemeProvider theme={theme}>
      <GlobalStyles />
      <Router>
        <DomainNavigation
          activeDomain="NLX"
          gitLabLink="https://gitlab.com/commonground/nlx/nlx"
        />

        <Header />

        <Switch>
          <Route
            path="/documentation/:organizationName/:serviceName"
            component={DocumentationPage}
          />
          <Route
            exact
            path="/:organizationName?/:serviceName?"
            component={ServiceOverviewPage}
          />
        </Switch>
      </Router>
      <VersionLogger />
    </ThemeProvider>
  </>
)

export default App
