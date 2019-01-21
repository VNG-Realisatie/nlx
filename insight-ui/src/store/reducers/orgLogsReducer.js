import cfg from '../app.cfg'
import * as actionType from '../actions'

function mergePageDef(state, payload) {
    let pageDef = {
        ...state.pageDef,
    }
    if (payload.page) {
        pageDef.page = payload.page
    }
    if (payload.rowCount) {
        pageDef.rowCount = payload.rowCount
    } else {
        pageDef.rowCount = payload.items.length
    }
    if (payload.rowsPerPage) {
        pageDef.rowsPerPage = payload.rowsPerPage
    }
    return pageDef
}

/**
 * Language reducer containing lang info and translations
 * @param {object} state previous state object
 * @param {string} action.type redux action type, string constant
 * @param {object} action.payload redux action type, string constant
 */
export const orgLogsReducer = (state = cfg.organization.logs, action) => {
    switch (action.type) {
        case actionType.GET_ORGANIZATION_LOGS_OK:
            let pageDef = mergePageDef(state, action.payload)

            return {
                ...state,
                ...action.payload,
                error: null,
                pageDef: {
                    ...pageDef,
                },
            }
        case actionType.GET_ORGANIZATION_LOGS_ERR:
            return {
                ...state,
                items: [],
                error: {
                    ...action.payload,
                },
            }
        case actionType.RESET_ORGANIZATION:
        case actionType.RESET_ORGANIZATION_LOGS:
            return {
                ...cfg.organization.logs,
            }
        default:
            return state
    }
}

export default orgLogsReducer
