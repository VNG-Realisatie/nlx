import React from 'react'
import { Route, Redirect } from 'react-router-dom'

import LoginPage from './pages/LoginPage/index'

const App = () =>
   <div>
     <Route exact path="/">
       <Redirect to="/inloggen" />
     </Route>

     <Route path="/inloggen">
       <LoginPage />
     </Route>
   </div>

export default App;
