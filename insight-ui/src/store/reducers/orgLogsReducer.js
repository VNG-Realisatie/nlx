import cfg from '../app.cfg'
import * as actionType from '../actions'
// eslint-disable-next-line
// import { logGroup } from '../../utils/logGroup';

/**
 * Language reducer containing lang info and translations
 * @param {object} state previous state object
 * @param {string} action.type redux action type, string constant
 * @param {object} action.payload redux action type, string constant
 */
export const orgLogsReducer = (state = cfg.organization.logs, action) => {
    // debugger
    switch (action.type) {
        case actionType.GET_ORGANIZATION_LOGS_OK:
            // debugger
            return {
                ...state,
                error: null,
                ...action.payload,
            }
        case actionType.GET_ORGANIZATION_LOGS_ERR:
            // debugger
            return {
                ...state,
                items: [],
                error: {
                    ...action.payload,
                },
            }
        case actionType.RESET_ORGANIZATION:
        case actionType.RESET_ORGANIZATION_LOGS:
            // debugger
            return {
                ...cfg.organization.logs,
            }
        // always return state
        // to continue 'event' chain
        default:
            return state
    }
}

export default orgLogsReducer
