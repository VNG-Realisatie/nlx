import cfg from '../app.cfg'
import * as actionType from '../actions'

/**
 * Manage the loader states using redux store
 * @param state: object, current redux store state of loader store
 * @param action: object, dispatched redux action
 */
export const locationReducer = (state = cfg.location.href, action) => {
    // just for fun use lowercased action types
    switch (action.type) {
        case actionType.LOCATION_CHANGE:
            // debugger
            return action.payload
        default:
            return state
    }
}

export default locationReducer
