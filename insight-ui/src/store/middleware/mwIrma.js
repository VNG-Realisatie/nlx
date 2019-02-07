/**
 * Custom middleware function to load QRCode from irma,
 * listens to GET_QRCODE action
 * it performs api call to load organizations list
 * on success dispatch GET_ORGANIZATIONS_OK or GET_ORGANIZATIONS_ERR action
 * @param param.getState: fn, received from redux
 * @param param.dispatch: fn, received from redux
 */

import axios from 'axios'
import * as actionType from '../actions'
import { extractError } from '../../utils/appUtils'
import cfg from '../app.cfg'
/**
 * Start login process for specified organization.
 * Combines 2 API requests and passes info recevied using setState.
 * - getDataSubjects - returns data subject we can choose from
 * - generateJWT - returns JWT based on organization and data subject selected
 * - create qrCode, statusUrl and proofUrl and pass it to state
 */
const initLoginProcess = ({ organization, dispatch }) => {
    if (!organization) {
        dispatch({
            type: actionType.GET_QRCODE_ERR,
            payload: {
                number: 400,
                description: 'Oranization info missing.',
            },
        })
    }
    let dataSubjects
    let urlGetDataSubjects = `${
        organization.insight_log_endpoint
    }/getDataSubjects`

    let urlGenerateJWT = `${organization.insight_log_endpoint}/generateJWT`
    let urlVerification = `${
        organization.insight_irma_endpoint
    }/api/v2/verification/`

    return axios({
        method: 'get',
        url: urlGetDataSubjects,
    })
        .then((response) => {
            dataSubjects = response.data['dataSubjects']

            let keys = Object.keys(dataSubjects)

            return axios({
                method: 'post',
                url: urlGenerateJWT,
                data: {
                    dataSubjects: keys,
                },
            })
        })
        .then((response) => {
            const firstJWT = response.data

            return axios({
                method: 'post',
                url: urlVerification,
                headers: { 'content-type': 'text/plain' },
                data: firstJWT,
            })
        })
        .then((response) => {
            let irmaVerificationRequest = response.data
            const u = irmaVerificationRequest['u']
            /**
             * For local testing with IRMA app
             * IF cfg.localIp is provided in app.cfg.js
             * we overwrite irma endpoint with localIp
             */
            if (cfg.localIp) {
                organization.insight_irma_endpoint = cfg.localIp
            }

            irmaVerificationRequest['u'] = `${
                organization.insight_irma_endpoint
            }/api/v2/verification/${u}`

            const qrCode = JSON.stringify(irmaVerificationRequest)

            const statusUrl = `${
                organization.insight_irma_endpoint
            }/api/v2/verification/${u}/status`

            const proofUrl = `${
                organization.insight_irma_endpoint
            }/api/v2/verification/${u}/getproof`

            dispatch({
                type: actionType.GET_QRCODE_OK,
                payload: {
                    name: organization.name,
                    dataSubjects,
                    qrCode,
                    statusUrl,
                    proofUrl,
                    firstJWT: u,
                    inProgress: false,
                },
            })
        })
        .catch((e) => {
            let error = extractError(e)
            dispatch({
                type: actionType.GET_QRCODE_ERR,
                payload: error,
            })
        })
}

/**
 * Remove interval
 */
let interval = null
function removeInterval() {
    if (interval) {
        clearInterval(interval)
        interval = null
    }
}

/**
 * Starts api call to check if user is authenticated.
 * The status returned from API is passed to handleLoginResponse method
 * The api call is delayed for 1 second.
 */
const getLoginStatus = ({ dispatch, getState }) => {
    let store = getState()
    let { irma } = store.organization
    let { statusUrl } = irma
    let inProgressFlag = false

    // only one interval allowed
    if (interval) {
        removeInterval()
    }

    interval = setInterval(() => {
        if (inProgressFlag === false) {
            let state = getState()
            if (state.organization.irma.inProgress === false) {
                dispatch({
                    type: actionType.IRMA_LOGIN_IN_PROGRESS,
                })
            } else {
                inProgressFlag = true
            }
        }
        axios({
            method: 'get',
            url: statusUrl,
        })
            .then((response) => {
                let { stop, action } = handleLoginResponse(response.data)
                if (action) {
                    dispatch(action)
                }
                if (stop) removeInterval()
            })
            .catch((e) => {
                let error = extractError(e)
                dispatch({
                    type: actionType.IRMA_LOGIN_ERR,
                    payload: error,
                })
                removeInterval()
            })
    }, 1000)
}

const handleLoginResponse = (response) => {
    switch (true) {
        case response.toUpperCase() === 'INITIALIZED' ||
            response.toUpperCase() === 'CONNECTED':
            return {
                stop: false,
                action: null,
            }

        case response.toUpperCase() === 'DONE':
            return {
                stop: true,
                action: {
                    type: actionType.IRMA_GET_PROOF,
                },
            }

        case response.toUpperCase() === 'CANCELLED':
            return {
                stop: true,
                action: {
                    type: actionType.IRMA_LOGIN_ERR,
                    payload: {
                        status: 200,
                        description: 'User refused to disclose attributes',
                    },
                },
            }

        default:
            return {
                stop: true,
                action: {
                    type: actionType.IRMA_LOGIN_ERR,
                    payload: {
                        status: 400,
                        description: `Unexpected response - ${response}`,
                    },
                },
            }
    }
}

const getProof = ({ dispatch, getState }) => {
    let store = getState()
    let { irma } = store.organization
    let { proofUrl } = irma

    axios({
        method: 'get',
        url: proofUrl,
    })
        .then((response) => {
            dispatch({
                type: actionType.IRMA_GET_PROOF_OK,
                payload: response.data,
            })
        })
        .catch((e) => {
            let error = extractError(e)
            dispatch({
                type: actionType.IRMA_GET_PROOF_ERR,
                payload: error,
            })
        })
}

export const mwIrma = ({ getState, dispatch }) => {
    return (next) => (action) => {
        next(action)

        switch (action.type) {
            case actionType.GET_QRCODE:
                return initLoginProcess({
                    organization: {
                        ...action.payload,
                    },
                    dispatch,
                })
            case actionType.IRMA_LOGIN_START:
                getLoginStatus({
                    getState,
                    dispatch,
                })
                break
            case actionType.IRMA_GET_PROOF:
                getProof({
                    getState,
                    dispatch,
                })
                break
            case actionType.IRMA_GET_PROOF_OK:
                let store = getState()
                let api = `${
                    store.organization.info.insight_log_endpoint
                }/fetch`

                let { page, rowsPerPage } = store.organization.logs.pageDef
                let params = {
                    page,
                    rowsPerPage,
                }
                dispatch({
                    type: actionType.GET_ORGANIZATION_LOGS,
                    payload: {
                        api,
                        name: store.organization.irma.name,
                        jwt: store.organization.irma.jwt,
                        params,
                    },
                })
                break
            case actionType.RESET_ORGANIZATION:
                removeInterval()
                break
            default:
        }
    }
}

export default mwIrma
