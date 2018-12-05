import cfg from '../app.cfg'
import * as actionType from '../actions'
/**
 * Language reducer containing lang info and translations
 * @param {object} state previous state object
 * @param {object} action redux action with type and payload
 * @param {string} action.type redux action type, string constant
 * @param {object} action.payload redux action type, string constant
 */
export const i18nReducer = (state = cfg.i18n, action) => {
    switch (action.type) {
        case actionType.SET_LANG_OK:
        case actionType.SET_LANG_ERR:
            return {
                ...state,
                lang: {
                    ...action.payload,
                },
            }
        case actionType.SET_LANG_LIST:
            return {
                ...state,
                options: [...action.payload],
            }
        default:
            return state
    }
}

export default i18nReducer
