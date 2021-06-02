// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import VersionLogger from './components/VersionLogger'
import { StyledApp } from './App.styles'
import GlobalStyles from './components/GlobalStyles/GlobalStyles'
import Header from './components/Header'
import ServiceOverviewPage from './pages/ServicesOverviewPage'
import DocumentationPage from './pages/DocumentationPage'

const App = () => (
  <StyledApp>
    <GlobalStyles />
    <Router>
      <>
        <Header />

        <Route exact path="/" component={ServiceOverviewPage} />
        <Route
          path="/documentation/:organizationName/:serviceName"
          component={DocumentationPage}
        />
      </>
    </Router>
    <VersionLogger />
  </StyledApp>
)

export default App
