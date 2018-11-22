/**
 * Custom middleware function to load organization list,
 * listens to GET_ORGANIZATIONS action
 * it performs api call to load organizations list
 * on success dispatch GET_ORGANIZATIONS_OK or GET_ORGANIZATIONS_ERR action
 * @param param.getState: fn, received from redux
 * @param param.dispatch: fn, received from redux
 */
// import { logGroup } from '../../utils/logGroup'
import axios from 'axios'
import * as actionType from '../actions'
import { extractError } from '../../utils/appUtils'

function getOrganizations({ action, dispatch }) {
    // debugger
    axios
        .get(action.payload)
        .then((response) => {
            if (response.data && response.data.organizations) {
                // filter only organizations with irma api point
                let orgList = response.data.organizations.filter((item) => {
                    return item.hasOwnProperty('insight_irma_endpoint')
                })
                dispatch({
                    type: actionType.GET_IRMA_ORGANIZATIONS_OK,
                    payload: orgList,
                })
            } else {
                dispatch({
                    type: actionType.GET_IRMA_ORGANIZATIONS_OK,
                    payload: [],
                })
            }
        })
        .catch((e) => {
            let error = extractError(e)
            dispatch({
                type: actionType.GET_IRMA_ORGANIZATIONS_ERR,
                payload: error,
            })
        })
}

export const mwOrganizations = ({ getState, dispatch }) => {
    return (next) => (action) => {
        // pass action
        next(action)
        // decide if additional work required
        // debugger
        switch (action.type) {
            case actionType.GET_IRMA_ORGANIZATIONS:
                getOrganizations({ action, dispatch })
                break
            case actionType.GET_IRMA_ORGANIZATIONS_OK:
                dispatch({
                    type: actionType.HIDE_LOADER,
                })
                break
            default:
            // do nothing
        }
    }
}

export default mwOrganizations
