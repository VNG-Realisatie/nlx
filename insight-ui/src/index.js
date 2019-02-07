import 'react-app-polyfill/ie11'

import React from 'react'
import ReactDOM from 'react-dom'
import App from './App'

import * as serviceWorker from './serviceWorker'

import { MuiThemeProvider } from '@material-ui/core/styles'
import muiTheme from './styles/muiTheme'

import { Provider } from 'react-redux'
import appStore from './store/redux'
import cfg from './store/app.cfg'
import * as actionType from './store/actions'

// get organizations
appStore.dispatch({
    type: actionType.GET_IRMA_ORGANIZATIONS,
    payload: cfg.organizations.api,
})

ReactDOM.render(
    <Provider store={appStore}>
        <MuiThemeProvider theme={muiTheme}>
            <App />
        </MuiThemeProvider>
    </Provider>,
    document.getElementById('root'),
)

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister()
