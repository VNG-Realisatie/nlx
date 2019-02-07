import { combineReducers } from 'redux'

import info from './info'
import logs from './logs'
import irma from './irma'

export default combineReducers({
    info,
    irma,
    logs,
})
