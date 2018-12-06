/**
 * Custom middleware function to load organization list,
 * listens to GET_ORGANIZATIONS action
 * it performs api call to load organizations list
 * on success dispatch GET_ORGANIZATIONS_OK or GET_ORGANIZATIONS_ERR action
 * @param param.getState: fn, received from redux
 * @param param.dispatch: fn, received from redux
 */

import axios from 'axios'
import * as actionType from '../actions'
import { extractError, sortArrayOnProp } from '../../utils/appUtils'

function getOrganizations({ action, dispatch }) {
    axios
        .get(action.payload)
        .then((response) => {
            if (response.data && response.data.organizations) {
                let orgList = response.data.organizations.filter((item) => {
                    return item.hasOwnProperty('insight_irma_endpoint')
                })
                orgList.sort(sortArrayOnProp('name', 'asc'))
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
        next(action)
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
