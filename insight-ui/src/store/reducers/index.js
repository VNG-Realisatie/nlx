// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { combineReducers } from 'redux'
import * as TYPES from '../types'
import organizations from './organizations'

export const loginStatus = (state = null, action) => {
    switch (action.type) {
        case TYPES.IRMA_LOGIN_REQUEST_SUCCESS:
        case TYPES.IRMA_LOGIN_REQUEST_FAILED:
            return action.data
        default:
            return state
    }
}

export const logs = (state = [], action) => {
    switch (action.type) {
        case TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS:
            return action.data.records
        default:
            return state
    }
}

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
    loginStatus: loginStatus,
    logs: logs
})
