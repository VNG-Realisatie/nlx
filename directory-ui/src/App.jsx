// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import React, { Fragment } from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import { VersionLogger } from '@commonground/design-system'

import { StyledApp } from './App.styles'

import GlobalStyles from './components/GlobalStyles/GlobalStyles'
import Header from './components/Header/Header'
import ServiceOverviewPage from './pages/ServicesOverviewPage/ServicesOverviewPage'
import DocumentationPage from './pages/DocumentationPage/DocumentationPage'

const App = () => (
    <StyledApp>
        <GlobalStyles />
        <Router>
            <Fragment>
                <Header />

                <Route exact path="/" component={ServiceOverviewPage} />
                <Route
                    path="/documentation/:organizationName/:serviceName"
                    component={DocumentationPage}
                />
            </Fragment>
        </Router>
        <VersionLogger />
    </StyledApp>
)

export default App
