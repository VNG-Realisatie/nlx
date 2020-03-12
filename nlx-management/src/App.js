import React from 'react'
import { Route, Redirect } from 'react-router-dom'
import { ThemeProvider } from 'styled-components/macro'
import {
  GlobalStyles as DSGlobalStyles,
  darkTheme,
} from '@commonground/design-system'
import GlobalStyles from './components/GlobalStyles'

import LoginPage from './pages/LoginPage/index'
import { StyledContainer } from './App.styles'

const App = () => (
  <StyledContainer>
    <ThemeProvider theme={darkTheme}>
      <DSGlobalStyles />
      <GlobalStyles />

      <Route exact path="/">
        <Redirect to="/inloggen" />
      </Route>

      <Route path="/inloggen">
        <LoginPage />
      </Route>
    </ThemeProvider>
  </StyledContainer>
)

export default App
