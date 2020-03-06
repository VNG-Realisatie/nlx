import React from 'react'
import { BrowserRouter as Router, Route, Redirect } from 'react-router-dom'

import LoginPage from './pages/LoginPage/index'

const App = () =>
  <div>
    <Router>
      <Route exact path="/">
        <Redirect to="/login" />
      </Route>

      <Route path="/login">
        <LoginPage>s</LoginPage>
      </Route>
    </Router>
  </div>

export default App;
