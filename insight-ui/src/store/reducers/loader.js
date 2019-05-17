// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import cfg from '../app.cfg'
import * as actionType from '../actions'

export const loaderReducer = (state = cfg.loader, action) => {
    switch (action.type) {
        case actionType.SHOW_LOADER:
            return {
                ...state,
                show: true,
            }
        case actionType.HIDE_LOADER:
            return {
                ...state,
                show: false,
            }
        default:
            return state
    }
}

export default loaderReducer
