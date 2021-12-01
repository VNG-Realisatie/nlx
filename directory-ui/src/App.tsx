// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Switch, BrowserRouter as Router, Route } from 'react-router-dom'
import { ThemeProvider } from 'styled-components'
import { GlobalStyles, DomainNavigation } from '@commonground/design-system'
import VersionLogger from './components/VersionLogger'
import Header from './components/Header'
import ServiceOverviewPage from './pages/services/ServicesPage'
import ParticipantsPage from './pages/participants'
import theme from './theme'
import '@fontsource/source-sans-pro/latin.css'

const App: React.FC = () => (
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
          <Route exact path="/participants" component={ParticipantsPage} />
          <Route
            exact
            path="/:organizationSerialNumber?/:serviceName?"
            component={ServiceOverviewPage}
          />
        </Switch>
      </Router>
      <VersionLogger />
    </ThemeProvider>
  </>
)

export default App
