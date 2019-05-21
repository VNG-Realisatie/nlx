// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import cfg from '../app.cfg'
import * as actionType from '../actions'

export const organizations = (state = cfg.organizations, action) => {
    switch (action.type) {
        case actionType.GET_IRMA_ORGANIZATIONS_OK:
            return {
                ...state,
                error: null,
                list: [...action.payload],
            }
        case actionType.GET_IRMA_ORGANIZATIONS_ERR:
            return {
                ...state,
                error: {
                    ...action.payload,
                },
                list: [],
            }
        default:
            return state
    }
}

export default organizations
