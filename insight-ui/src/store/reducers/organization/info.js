// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import cfg from '../../app.cfg'
import * as actionType from '../../actions'

export const info = (state = cfg.organization.info, action) => {
    switch (action.type) {
        case actionType.SELECT_ORGANIZATION:
            return {
                ...state,
                ...action.payload,
            }
        case actionType.RESET_ORGANIZATION:
        case actionType.RESET_ORGANIZATION_LOGS:
            return {
                ...cfg.organization.info,
            }

        default:
            return state
    }
}

export default info
