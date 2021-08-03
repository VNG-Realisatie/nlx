// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { ThemeProvider } from 'styled-components'
import { DomainNavigation } from '@commonground/design-system'
import VersionLogger from './components/VersionLogger'
import GlobalStyles from './styling/GlobalStyles'
import Header from './components/Header'
import ServiceOverviewPage from './pages/ServicesOverviewPage'
import DocumentationPage from './pages/DocumentationPage'
import theme from './styling/theme'

const App = () => (
  <main>
    <GlobalStyles />
    <Router>
      <ThemeProvider theme={theme}>
        <DomainNavigation
          activeDomain="NLX"
          gitLabLink="https://gitlab.com/commonground/nlx/nlx"
        />

        <Header />

        <Route exact path="/" component={ServiceOverviewPage} />
        <Route
          path="/documentation/:organizationName/:serviceName"
          component={DocumentationPage}
        />
      </ThemeProvider>
    </Router>
    <VersionLogger />
  </main>
)

export default App
