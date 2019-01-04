import React, { Component } from 'react'
import Navigation from './components/Navigation'
import OverviewPage from './pages/OverviewPage/OverviewPage'

import { BrowserRouter as Router, Route } from 'react-router-dom'

class App extends Component {
    render() {
        return (
            <div className="App">
                <Router>
                    <div>
                        <Navigation />
                        <Route exact path="/" component={OverviewPage} />
                    </div>
                </Router>
            </div>
        )
    }
}

export default App
