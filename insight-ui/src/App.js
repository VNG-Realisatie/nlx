import React, { Component } from 'react'
import { BrowserRouter } from 'react-router-dom'

import { withStyles, CssBaseline } from '@material-ui/core'
import { globalStyles } from './styles/muiTheme'

import Drawer from './components/Drawer'

class App extends Component {
    render() {
        return (
            <BrowserRouter>
                <div className="App">
                    <CssBaseline />
                    <Drawer appTitle="NLX Insights" />
                </div>
            </BrowserRouter>
        )
    }
}

export default withStyles(globalStyles)(App)
