import * as actionType from '../actions'
import cfg from '../app.cfg'
/**
 * Manage the loader states using redux store
 * @param state: object, current redux store state of loader store
 * @param action: object, dispatched redux action
 */

export const orgIrmaReducer = (state = cfg.organization.irma, action) => {
    // just for fun use lowercased action types
    switch (action.type) {
        case actionType.GET_QRCODE_OK:
            // debugger
            return {
                ...state,
                ...action.payload,
            }
        case actionType.IRMA_LOGIN_START:
        case actionType.IRMA_LOGIN_IN_PROGRESS:
            // debugger
            return {
                ...state,
                inProgress: true,
            }
        case actionType.GET_QRCODE_ERR:
        case actionType.IRMA_LOGIN_ERR:
        case actionType.IRMA_GET_PROOF_ERR:
            // debugger
            return {
                ...state,
                qrCode: null,
                jwt: null,
                error: {
                    ...action.payload,
                },
            }
        case actionType.IRMA_GET_PROOF_OK:
            // debugger
            return {
                ...state,
                error: null,
                inProgress: false,
                jwt: action.payload,
            }
        case actionType.RESET_ORGANIZATION:
            // reset irma object of organization
            // debugger
            return {
                ...cfg.organization.irma,
            }
        default:
            return state
    }
}

export default orgIrmaReducer
