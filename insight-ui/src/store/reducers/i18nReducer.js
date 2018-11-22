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
export const i18nReducer = (state = cfg.i18n, action) => {
    /* logGroup({
    title:'filterReducer',
    action: action,
    state: state
    }) */
    // debugger
    switch (action.type) {
        case actionType.SET_LANG_OK:
        case actionType.SET_LANG_ERR:
            // debugger
            return {
                ...state,
                lang: {
                    ...action.payload,
                },
            }
        case actionType.SET_LANG_LIST:
            // debugger
            return {
                ...state,
                options: [...action.payload],
            }
        // always return state
        // to continue 'event' chain
        default:
            return state
    }
}

export default i18nReducer
