// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Routes, BrowserRouter, Route } from 'react-router-dom'
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
      <BrowserRouter>
        <DomainNavigation
          activeDomain="NLX"
          gitLabLink="https://gitlab.com/commonground/nlx/nlx"
        />

        <Header />

        <Routes>
          <Route path="/participants" element={<ParticipantsPage />} />
          <Route path="*" element={<ServiceOverviewPage />} />
        </Routes>
      </BrowserRouter>
      <VersionLogger />
    </ThemeProvider>
  </>
)

export default App
