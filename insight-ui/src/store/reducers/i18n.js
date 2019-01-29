import cfg from '../app.cfg'
import * as actionType from '../actions'

export const i18n = (state = cfg.i18n, action) => {
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

export default i18n
