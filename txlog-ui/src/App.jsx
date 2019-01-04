import React, { Component } from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import Navigation from './components/Navigation/Navigation'
import OverviewPage from './pages/OverviewPage/OverviewPage'

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
