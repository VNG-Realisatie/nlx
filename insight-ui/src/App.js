import React, { Component } from 'react'
// import { BrowserRouter } from 'react-router-dom'

import { withStyles, CssBaseline } from '@material-ui/core'
import { globalStyles } from './styles/muiTheme'

import Drawer from './layout/Drawer'

class App extends Component {
    render() {
        return (
            <div className="App">
                <CssBaseline />
                <Drawer />
            </div>
        )
    }
}

export default withStyles(globalStyles)(App)
