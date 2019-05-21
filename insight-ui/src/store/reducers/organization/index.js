// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { combineReducers } from 'redux'

import info from './info'
import logs from './logs'
import irma from './irma'

export default combineReducers({
    info,
    irma,
    logs,
})
