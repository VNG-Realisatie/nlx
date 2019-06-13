// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { VersionLogger } from '@commonground/design-system'
import Navigation from './components/Navigation/Navigation'
import OverviewPage from './pages/OverviewPage/OverviewPage'

const App = () => (
    <div className="App">
        <Router>
            <div>
                <Navigation />
                <Route exact path="/" component={OverviewPage} />
            </div>
        </Router>
        <VersionLogger />
    </div>
)

export default App
