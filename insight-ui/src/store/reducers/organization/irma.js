// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import * as actionType from '../../actions'
import cfg from '../../app.cfg'

export const irma = (state = cfg.organization.irma, action) => {
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

export default irma
