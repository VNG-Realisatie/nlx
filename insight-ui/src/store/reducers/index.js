// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { combineReducers } from 'redux'
import * as TYPES from '../types'

export default combineReducers({
    organizations: (state = [], action) => {
        switch (action.type) {
            case TYPES.FETCH_ORGANIZATIONS_SUCCESS:
                return action.data
            default:
                return state
        }
    },
    loginInformation: (state = {}, action) => {
        switch (action.type) {
            case TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS:
                return action.data
            default:
                return state
        }
    }
})
