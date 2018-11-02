import React, { Component } from 'react'
import Navigation from './components/Navigation'
import Overview from './Overview'

import { BrowserRouter as Router, Route } from 'react-router-dom'

class App extends Component {
    render() {
        return (
            <div className="App">
                <Router>
                    <div>
                        <Navigation />
                        <Route exact path="/" component={Overview} />
                    </div>
                </Router>
            </div>
        )
    }
}

export default App
