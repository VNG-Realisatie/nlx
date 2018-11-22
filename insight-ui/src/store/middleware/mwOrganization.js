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

function getOrganizationLogs({ action, dispatch }) {
    let { api, jwt, name } = action.payload

    // debugger

    if (api && jwt) {
        axios({
            method: 'post',
            url: api,
            headers: { 'content-type': 'text/plain' },
            data: jwt,
        })
            .then((response) => {
                // debugger
                let items = response.data.records
                dispatch({
                    type: actionType.GET_ORGANIZATION_LOGS_OK,
                    payload: {
                        name,
                        api,
                        items,
                        jwt,
                    },
                })
            })
            .catch((e) => {
                // debugger
                let error = extractError(e)
                dispatch({
                    type: actionType.GET_ORGANIZATION_LOGS_ERR,
                    payload: error,
                })
            })
    } else {
        // debugger
        dispatch({
            type: actionType.GET_ORGANIZATION_LOGS_ERR,
            payload: {
                status: 600,
                description: 'url or jwt missing',
            },
        })
    }
}

export const mwOrganization = ({ getState, dispatch }) => {
    return (next) => (action) => {
        // pass action
        next(action)
        // decide if additional work required
        // debugger
        switch (action.type) {
            case actionType.SELECT_ORGANIZATION:
                // now get qrCode
                // debugger
                dispatch({
                    type: actionType.GET_QRCODE,
                    payload: action.payload,
                })
                break
            case actionType.GET_ORGANIZATION_LOGS:
                // debugger
                getOrganizationLogs({ action, dispatch })
                break
            default:
            // do nothing
        }
    }
}

export default mwOrganization
