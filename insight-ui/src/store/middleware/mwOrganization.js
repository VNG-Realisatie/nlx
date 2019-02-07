/**
 * Custom middleware function to load organization logs,
 * listens to GET_ORGANIZATION_LOGS action
 * it performs api call to load organizations list
 * on success dispatch GET_ORGANIZATION_LOGS_OK or
 * GET_ORGANIZATION_LOGS_ERR
 * @param param.action: fn, received from redux
 * @param param.dispatch: fn, received from redux
 */

import axios from 'axios'
import * as actionType from '../actions'
import { extractError } from '../../utils/appUtils'

function getOrganizationLogs({ action, dispatch }) {
    if (!action.payload) {
        return dispatch({
            type: actionType.GET_ORGANIZATION_LOGS_ERR,
            payload: {
                status: 600,
                description: 'url or jwt missing',
            },
        })
    }
    let { api, jwt, name, params } = action.payload

    if (api && jwt) {
        axios({
            method: 'post',
            url: api,
            headers: { 'content-type': 'text/plain' },
            params,
            data: jwt,
        })
            .then((response) => {
                let { records, page, rowCount } = response.data
                // if page not returned by backend
                // use original param value
                if (typeof page === 'undefined') page = params.page

                dispatch({
                    type: actionType.GET_ORGANIZATION_LOGS_OK,
                    payload: {
                        name,
                        api,
                        jwt,
                        items: records,
                        page,
                        rowCount,
                        rowsPerPage: params.rowsPerPage,
                    },
                })
            })
            .catch((e) => {
                let error = extractError(e)
                dispatch({
                    type: actionType.GET_ORGANIZATION_LOGS_ERR,
                    payload: error,
                })
            })
    } else {
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
        next(action)

        switch (action.type) {
            case actionType.SELECT_ORGANIZATION:
                dispatch({
                    type: actionType.GET_QRCODE,
                    payload: action.payload,
                })
                break
            case actionType.GET_ORGANIZATION_LOGS:
                getOrganizationLogs({ action, dispatch })
                break
            default:
            // do nothing
        }
    }
}

export default mwOrganization
