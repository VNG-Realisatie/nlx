// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { combineReducers } from 'redux'
import organizations from './organizations'
import logs from './logs'
import loginStatus from './loginStatus'
import loginRequestInfo from './loginRequestInfo'
import proof from './proof'

export default combineReducers({
    organizations,
    loginRequestInfo,
    loginStatus,
    logs,
    proof
})
