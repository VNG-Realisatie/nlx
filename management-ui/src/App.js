// Copyright © VNG Realisatie 2019
// Licensed under the EUPL
import React, { Fragment } from 'react'
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom'
import { ThemeProvider } from 'styled-components'

import theme from './theme'
import GlobalStyles from './components/GlobalStyles'

import Home from './pages/Home'
import InwayList from './pages/InwayList'
import ServiceList from './pages/ServiceList'
import ServiceCreate from './pages/ServiceCreate'
import ServiceUpdate from './pages/ServiceUpdate'
import NotFound from './pages/NotFound'

const App = () => (
    <ThemeProvider theme={theme}>
        <Fragment>
            <GlobalStyles />
            <Router>
                <Switch>
                    <Route exact path="/" component={Home} />
                    <Route exact path="/inways" component={InwayList} />
                    <Route exact path="/services" component={ServiceList} />
                    <Route
                        exact
                        path="/services/create"
                        component={ServiceCreate}
                    />
                    <Route
                        exact
                        path="/services/update/:name"
                        component={ServiceUpdate}
                    />
                    <Route component={NotFound} />
                </Switch>
            </Router>
        </Fragment>
    </ThemeProvider>
)

export default App
