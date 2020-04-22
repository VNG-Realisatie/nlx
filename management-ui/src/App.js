// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { Suspense } from 'react'
import { ThemeProvider } from 'styled-components/macro'
import { GlobalStyles as DSGlobalStyles } from '@commonground/design-system'

import GlobalStyles from './components/GlobalStyles'
import theme from './theme'
import Routes from './routes'

import { StyledContainer } from './App.styles'

const App = () => (
  <StyledContainer>
    <ThemeProvider theme={theme}>
      <DSGlobalStyles />
      <GlobalStyles />

      {/* Suspense is required for XHR backend i18next */}
      <Suspense fallback={null}>
        <Routes />
      </Suspense>
    </ThemeProvider>
  </StyledContainer>
)

export default App
