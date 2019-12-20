// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'

import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import { func } from 'prop-types'

import Login from './pages/Login'

const AuthenticatedApp = ({ login }) => (
    <Router>
        <Switch>
            <Route render={() => <Login login={login} />} />
        </Switch>
    </Router>
)

AuthenticatedApp.propTypes = {
    login: func.isRequired,
}

export default AuthenticatedApp
