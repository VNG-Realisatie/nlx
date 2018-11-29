import cfg from '../app.cfg'
import * as actionType from '../actions'

/**
 * Language reducer containing lang info and translations
 * @param {object} state previous state object
 * @param {string} action.type redux action type, string constant
 * @param {object} action.payload redux action type, string constant
 */
export const orgsReducer = (state = cfg.organizations, action) => {
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

export default orgsReducer
