// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { createStore, applyMiddleware } from 'redux'
import { composeWithDevTools } from 'redux-devtools-extension'
import reducers from './reducers'

// MIDDLEWARE -> mw
import mwOrganizations from './middleware/mwOrganizations'
import mwOrganization from './middleware/mwOrganization'
import mwIrma from './middleware/mwIrma'

const appStore = createStore(
    reducers,
    composeWithDevTools(
        applyMiddleware(mwOrganizations, mwIrma, mwOrganization),
    ),
)

export default appStore
