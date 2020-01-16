// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React from 'react'

import {
    BrowserRouter as Router,
    Route,
    Switch,
    Redirect,
} from 'react-router-dom'
import { func } from 'prop-types'

import Login from './pages/Login'

const UnauthenticatedApp = ({ login }) => (
    <Router>
        <Switch>
            <Route
                exact
                path="/logout"
                render={() => <Redirect to={{ pathname: '/' }} />}
            />
            <Route render={() => <Login login={login} />} />
        </Switch>
    </Router>
)

UnauthenticatedApp.propTypes = {
    login: func.isRequired,
}

export default UnauthenticatedApp
