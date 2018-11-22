/**
 * Custom middleware function to load QRCode from irma,
 * listens to GET_QRCODE action
 * it performs api call to load organizations list
 * on success dispatch GET_ORGANIZATIONS_OK or GET_ORGANIZATIONS_ERR action
 * @param param.getState: fn, received from redux
 * @param param.dispatch: fn, received from redux
 */
// import { logGroup } from '../../utils/logGroup'
import axios from 'axios'
import * as actionType from '../actions'
import { extractError } from '../../utils/appUtils'

/**
 * Start login process for specified organization.
 * Combines 2 API requests and passes info recevied using setState.
 * - getDataSubjects - returns data subject we can choose from
 * - generateJWT - returns JWT based on organization and data subject selected
 * - create qrCode, statusUrl and proofUrl and pass it to state
 */
const initLoginProcess = ({ organization, dispatch }) => {
    // debugger
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

    axios({
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
            // debugger
            // prepend IrmaVerifiationRequest with URL
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
    // debugger
    let store = getState()
    let { irma } = store.organization
    let { statusUrl } = irma
    let inProgressFlag = false

    // only one interval allowed
    if (interval) {
        // debugger
        // eslint-disable-next-line
        console.warn(
            `getLoginStatus...
            attempt to init second interval...
            clearing previous interval first`,
        )
        removeInterval()
    }

    interval = setInterval(() => {
        // debugger
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
                    // debugger
                    dispatch(action)
                }
                if (stop) removeInterval()
            })
            .catch((e) => {
                // debugger
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
    // decide on action
    switch (true) {
        // we wait QRcode to be scanned
        case response.toUpperCase() === 'INITIALIZED' ||
            // code scanned wait for approval
            response.toUpperCase() === 'CONNECTED':
            return {
                stop: false,
                action: null,
            }
        // QRCode auth completed
        case response.toUpperCase() === 'DONE':
            return {
                stop: true,
                action: {
                    type: actionType.IRMA_GET_PROOF,
                },
            }
        // QRCode auth cancelled
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
        // unexpected value returned
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
    // debugger
    let store = getState()
    let { irma } = store.organization
    let { proofUrl } = irma

    axios({
        method: 'get',
        url: proofUrl,
    })
        .then((response) => {
            // debugger
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
        // pass current action
        // note! that action will reach reducers
        // before switch statement on the next line
        // is executed
        next(action)
        // decide additional actions
        switch (action.type) {
            case actionType.GET_QRCODE:
                // debugger
                initLoginProcess({
                    organization: {
                        ...action.payload,
                    },
                    dispatch,
                })
                break
            case actionType.IRMA_LOGIN_START:
                // debugger
                getLoginStatus({
                    getState,
                    dispatch,
                })
                break
            case actionType.IRMA_GET_PROOF:
                // debugger
                getProof({
                    getState,
                    dispatch,
                })
                break
            case actionType.IRMA_GET_PROOF_OK:
                // get current state from redux store
                // because next() is executed before switch
                // we can pull infro from store including
                // the values provided in the current action
                // debugger
                let store = getState()
                let api = `${
                    store.organization.info.insight_log_endpoint
                }/fetch`
                dispatch({
                    type: actionType.GET_ORGANIZATION_LOGS,
                    payload: {
                        api,
                        name: store.organization.irma.name,
                        jwt: store.organization.irma.jwt,
                    },
                })
                break
            case actionType.RESET_ORGANIZATION:
                // debugger
                removeInterval()
                break
            default:
            // do nothing
        }
    }
}

export default mwIrma
