import React from 'react'

import GlobalStyles from './components/GlobalStyles/GlobalStyles'
import Navigation from './components/Navigation/Navigation'
import ServiceOverviewPage from './pages/ServicesOverviewPage/ServicesOverviewPage'
import DocumentationPage from './pages/DocumentationPage/DocumentationPage'

import { BrowserRouter as Router, Route } from 'react-router-dom'

const App = () => (
    <div className="App">
        <GlobalStyles/>
        <Router>
            <div>
                <Navigation />

                <Route exact path="/" component={ServiceOverviewPage} />
                <Route
                    path="/documentation/:organization_name/:service_name"
                    component={DocumentationPage}
                />
            </div>
        </Router>
    </div>
)

export default App
