import cfg from '../app.cfg'
import * as actionType from '../actions'
// eslint-disable-next-line
import { logGroup } from '../../utils/logGroup';

/**
 * Language reducer containing lang info and translations
 * @param {object} state previous state object
 * @param {object} action redux action with type and payload
 * @param {string} action.type redux action type, string constant
 * @param {object} action.payload redux action type, string constant
 */

export const orgsReducer = (state = cfg.organizations, action) => {
    // debugger
    switch (action.type) {
        case actionType.GET_IRMA_ORGANIZATIONS_OK:
            // debugger
            return {
                ...state,
                error: null,
                list: [...action.payload],
            }
        case actionType.GET_IRMA_ORGANIZATIONS_ERR:
            // debugger
            return {
                ...state,
                error: {
                    ...action.payload,
                },
                list: [],
            }
        // always return state
        // to continue 'event' chain
        default:
            return state
    }
}

export default orgsReducer
