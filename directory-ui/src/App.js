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
import theme from './styling/theme'
import GlobalFonts from './styling/GlobalFonts'

const App = () => (
  <main>
    <Router>
      <ThemeProvider theme={theme}>
        <GlobalStyles />
        <GlobalFonts />

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
          <Route path="/" component={ServiceOverviewPage} />
        </Switch>
      </ThemeProvider>
    </Router>
    <VersionLogger />
  </main>
)

export default App
