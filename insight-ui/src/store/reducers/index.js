// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { combineReducers } from 'redux'
import * as TYPES from '../types'
import organizations from './organizations'
import logs from './logs'
import loginStatus from './loginStatus'

export default combineReducers({
    organizations,
    loginRequestInfo: (state = {}, action) => {
        switch (action.type) {
            case TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS:
                return action.data
            default:
                return state
        }
    },
    loginStatus,
    logs
})
