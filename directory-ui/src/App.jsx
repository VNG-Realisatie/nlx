import React from 'react'
import Navigation from './components/Navigation/Navigation'
import Directory from './Directory'
import DocumentationPage from './pages/DocumentationPage/DocumentationPage'

import './static/css/base-addon.css'

import { BrowserRouter as Router, Route } from 'react-router-dom'

const App = () => (
    <div className="App">
        <Router>
            <div>
                <Navigation />

                <Route exact path="/" component={Directory} />
                <Route
                    path="/documentation/:organization_name/:service_name"
                    component={DocumentationPage}
                />
            </div>
        </Router>
    </div>
)

export default App
