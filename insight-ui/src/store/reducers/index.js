// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { combineReducers } from 'redux'
import * as TYPES from '../types'

const filterOutInvalidOrganizations = organizations =>
  organizations
    .filter(organization =>
      organization.insight_irma_endpoint && organization.insight_log_endpoint
    )

export const organizations = (state = [], action) => {
    switch (action.type) {
        case TYPES.FETCH_ORGANIZATIONS_SUCCESS:
            return filterOutInvalidOrganizations(action.data)
        default:
            return state
    }
}

export default combineReducers({
    organizations: organizations,
    loginInformation: (state = {}, action) => {
        switch (action.type) {
            case TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS:
                return action.data
            default:
                return state
        }
    }
})
