/**
 * Custom middleware function that logs when called,
 * it should return next(action) function
 * and call it to continue action 'chain' and reach the reducer
 * note! if next(action) is not called the chain will break and action
 * will never reach the reducer
 * @param param.getState: fn, received from redux
 * @param param.dispatch: fn, received from redux
 */
import { logGroup } from '../../utils/logGroup'
export const actionLogger = ({ getState, dispatch }) => {
    return (next) => (action) => {
        // call next to continue the chain
        next(action)
        // log
        logGroup({
            title: `ACTION: ${action.type}`,
            method: 'actionLogger',
            ...action,
        })
    }
}

export default actionLogger
