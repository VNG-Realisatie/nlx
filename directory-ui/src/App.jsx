// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Fragment } from 'react'

import GlobalStyles from './components/GlobalStyles/GlobalStyles'
import Header from './components/Header/Header'
import ServiceOverviewPage from './pages/ServicesOverviewPage/ServicesOverviewPage'
import DocumentationPage from './pages/DocumentationPage/DocumentationPage'

import { BrowserRouter as Router, Route } from 'react-router-dom'

const App = () => (
    <div className="App">
        <GlobalStyles/>
        <Router>
            <Fragment>
                <Header />

                <Route exact path="/" component={ServiceOverviewPage} />
                <Route
                    path="/documentation/:organization_name/:service_name"
                    component={DocumentationPage}
                />
            </Fragment>
        </Router>
    </div>
)

export default App
