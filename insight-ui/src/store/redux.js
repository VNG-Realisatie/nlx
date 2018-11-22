import { createStore, combineReducers, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'

// REDUCERS
import loaderReducer from './reducers/loaderReducer'
import i18nReducer from './reducers/i18nReducer'
import orgsReducer from './reducers/orgsReducer'
import orgInfoReducer from './reducers/orgInfoReducer'
import orgLogsReducer from './reducers/orgLogsReducer'
import orgIrmaReducer from './reducers/orgIrmaReducer'
// import loactionReducer from './reducers/locationReducers'
// MIDDLEWARE -> mw
// import mwActionLogger from './middleware/mwActionLogger'
import mwOrganizations from './middleware/mwOrganizations'
import mwOrganization from './middleware/mwOrganization'
import mwIrma from './middleware/mwIrma'

const orgReducer = combineReducers({
    info: orgInfoReducer,
    irma: orgIrmaReducer,
    logs: orgLogsReducer,
})

const reducers = combineReducers({
    loader: loaderReducer,
    i18n: i18nReducer,
    organizations: orgsReducer,
    organization: orgReducer,
    // href: loactionReducer,
})

const appStore = createStore(
    reducers,
    composeWithDevTools(
        applyMiddleware(
            // NOTE! the middleware order matters
            // mwActionLogger,
            mwOrganizations,
            mwIrma,
            mwOrganization,
        ),
    ),
)

export default appStore
