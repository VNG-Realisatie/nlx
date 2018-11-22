import cfg from '../app.cfg'
import * as actionType from '../actions'

/**
 * Manage the loader states using redux store
 * @param state: object, current redux store state of loader store
 * @param action: object, dispatched redux action
 */
export const loaderReducer = (state = cfg.loader, action) => {
    // just for fun use lowercased action types
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
