import React from 'react'
import { Route, Redirect } from 'react-router-dom'
import { ThemeProvider } from 'styled-components'

import theme from './theme'

import LoginPage from './pages/LoginPage/index'
import GlobalStyles from './components/GlobalStyles'
import { StyledContainer } from './App.styles'

const App = () => (
  <StyledContainer>
    <ThemeProvider theme={theme}>
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
