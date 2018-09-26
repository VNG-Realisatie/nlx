import React, { Component } from 'react';
import { BrowserRouter } from 'react-router-dom';

import { withStyles, CssBaseline } from '@material-ui/core';
import { globalStyles} from './styles/muiTheme';

import ResponsiveDrawer from './components/Drawer';

class App extends Component {
    state={
        appTitle:"NLx Insights"
    }
    componentDidMount() {
        //this.props.getSessionData();
        //console.log(this.props)
    }

    render() {
        return (            
            <BrowserRouter>                
                <div className="App">
                    <CssBaseline />
                    <ResponsiveDrawer
                        appTitle={this.state.appTitle}/>
                </div>
            </BrowserRouter>            
        );
    }
}

export default withStyles(globalStyles)(App);
