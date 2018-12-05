import * as actionType from '../actions'
import cfg from '../app.cfg'
/**
 * Manage the loader states using redux store
 * @param state: object, current redux store state of loader store
 * @param action: object, dispatched redux action
 */
export const orgIrmaReducer = (state = cfg.organization.irma, action) => {
    switch (action.type) {
        case actionType.GET_QRCODE_OK:
            return {
                ...state,
                ...action.payload,
            }
        case actionType.IRMA_LOGIN_START:
        case actionType.IRMA_LOGIN_IN_PROGRESS:
            return {
                ...state,
                inProgress: true,
            }
        case actionType.GET_QRCODE_ERR:
        case actionType.IRMA_LOGIN_ERR:
        case actionType.IRMA_GET_PROOF_ERR:
            return {
                ...state,
                qrCode: null,
                jwt: null,
                error: {
                    ...action.payload,
                },
            }
        case actionType.IRMA_GET_PROOF_OK:
            return {
                ...state,
                error: null,
                inProgress: false,
                jwt: action.payload,
            }
        case actionType.RESET_ORGANIZATION:
            return {
                ...cfg.organization.irma,
            }
        default:
            return state
    }
}

export default orgIrmaReducer
