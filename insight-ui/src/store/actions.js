// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { call, put, retry } from 'redux-saga/effects'
import * as TYPES from './types'

export const api = url => fetch(url).then(response => response.json())
export const apiPost = (url, data) =>
  fetch(url,{
    method: 'POST',
    body: JSON.stringify(data)
  })

export const apiPostWithTextResponse = (url, data) => {
  return apiPost(url, data)
    .then(response => response.text())
}
export const apiPostWithJSONResponse = (url, data) => {
  return apiPost(url, data)
    .then(response => response.json())
}

export const fetchOrganizationsRequest = () => ({
  type: TYPES.FETCH_ORGANIZATIONS_REQUEST
})

export const fetchIrmaLoginInformationRequest = ({ insight_log_endpoint, insight_irma_endpoint }) =>
  ({
    type: TYPES.FETCH_IRMA_LOGIN_INFORMATION_REQUEST,
    data: { insight_log_endpoint, insight_irma_endpoint }
  })

export function* fetchOrganizations() {
  try {
    const organizations = yield call(api, '/api/directory/list-organizations')
    yield put({ type: TYPES.FETCH_ORGANIZATIONS_SUCCESS, data: organizations.organizations })
  } catch (err) {
    console.log(err);
  }
}

function mapDataSubjectsResponseToArray(dataSubjectsResponse) {
  return Object.keys(dataSubjectsResponse.dataSubjects)
}

export function* fetchIrmaLoginInformation({ insight_log_endpoint, insight_irma_endpoint }) {
  try {
    const dataSubjects = yield call(api, `${insight_log_endpoint}/getDataSubjects`)
    const jwtToken = yield call(apiPostWithTextResponse, `${insight_log_endpoint}/generateJWT`, { dataSubjects: mapDataSubjectsResponseToArray(dataSubjects) })
    const JWTVerification = yield call(apiPostWithJSONResponse, `${insight_irma_endpoint}/api/v2/verification/`, jwtToken)
    const u = JWTVerification.u

    const qrCodeContents = {
      u:`${insight_irma_endpoint}/api/v2/verification/${u}`,
      v: JWTVerification.v,
      vmax: JWTVerification.vmax,
      irmaqr: JWTVerification.irmaqr
    }

    yield put({ type: TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS, data: {
        dataSubjects: mapDataSubjectsResponseToArray(dataSubjects),
        qrCodeValue: JSON.stringify(qrCodeContents),
        statusUrl: `${insight_irma_endpoint}/api/v2/verification/${u}/status`,
        proofUrl: `${insight_irma_endpoint}/api/v2/verification/${u}/getproof`,
        JWT: u,
    }})
  } catch (err) {
    console.log(err);
  }
}

export const IRMA_LOGIN_STATUS_INITIALIZED = 'INITIALIZED'
export const IRMA_LOGIN_STATUS_CONNECTED = 'CONNECTED'
export const IRMA_LOGIN_STATUS_CANCELLED = 'CANCELLED'
export const IRMA_LOGIN_STATUS_DONE = 'DONE'
export const IRMA_LOGIN_STATUS_SESSION_UNKNOWN = 'SESSION_UNKNOWN'

export const apiHandleLoginStatus = url =>
  api(url)
    .then(response => {
      if (typeof response === 'string') {
        return response.toUpperCase()
      }

      if (response.error === IRMA_LOGIN_STATUS_SESSION_UNKNOWN) {
        return IRMA_LOGIN_STATUS_SESSION_UNKNOWN
      }

      return JSON.stringify(response)
    })
    .then(response => {
      switch (response) {
        case IRMA_LOGIN_STATUS_INITIALIZED:
          throw new Error('Login is initialized but not yet confirmed.')

        case IRMA_LOGIN_STATUS_CONNECTED:
          throw new Error('User is connected, but has not yet confirmed.')

        case IRMA_LOGIN_STATUS_DONE:
        case IRMA_LOGIN_STATUS_CANCELLED:
        case IRMA_LOGIN_STATUS_SESSION_UNKNOWN:
          return response

        default:
          console.error(`Unexpected response '${response}' while getting the login status.`)
          return response
      }
    })

export function* getIrmaLoginStatus({ statusUrl }, api = apiHandleLoginStatus) {
  try {
    const SECOND = 1000
    const response = yield retry(60, 1 * SECOND, api, statusUrl)
    yield put({ type: TYPES.IRMA_LOGIN_REQUEST_SUCCESS, data: {
      error: response !== IRMA_LOGIN_STATUS_DONE,
      response
    }})
  } catch (error) {
    yield put({ type: TYPES.IRMA_LOGIN_REQUEST_FAILED, data: {
      error: true,
      response: error
    }})
  }
}

// old actions
// loader actions
export const SHOW_LOADER = 'SHOW_LOADER'
export const HIDE_LOADER = 'HIDE_LOADER'

// language
export const GET_LANGUAGE = 'GET_LANGUAGE'
export const SET_LANG_OK = 'SET_LANG_OK'
export const SET_LANG_ERR = 'SET_LANG_ERR'
export const SET_LANG_LIST = 'SET_LANG_LIST'

// organization list
export const GET_IRMA_ORGANIZATIONS = 'GET_IRMA_ORGANIZATIONS'
export const GET_IRMA_ORGANIZATIONS_OK = 'GET_IRMA_ORGANIZATIONS_OK'
export const GET_IRMA_ORGANIZATIONS_ERR = 'GET_IRMA_ORGANIZATIONS_ERR'

// organization (specific)
export const SELECT_ORGANIZATION = 'SELECT_ORGANIZATION'
export const RESET_ORGANIZATION = 'RESET_ORGANIZATION'
// irma
export const GET_QRCODE = 'GET_QRCODE'
export const GET_QRCODE_OK = 'GET_QRCODE_OK'
export const GET_QRCODE_ERR = 'GET_QRCODE_ERR'
export const IRMA_LOGIN_START = 'IRMA_LOGIN_START'
export const IRMA_LOGIN_IN_PROGRESS = 'IRMA_LOGIN_IN_PROGRESS'
export const IRMA_GET_PROOF = 'IRMA_GET_PROOF'
export const IRMA_GET_PROOF_OK = 'IRMA_GET_PROOF_OK'
export const IRMA_GET_PROOF_ERR = 'IRMA_GET_PROOF_ERR'
export const IRMA_LOGIN_ERR = 'IRMA_LOGIN_ERR'
// logs
export const GET_ORGANIZATION_LOGS = 'GET_ORGANIZATION_LOGS'
export const GET_ORGANIZATION_LOGS_OK = 'GET_ORGANIZATION_LOGS_OK'
export const GET_ORGANIZATION_LOGS_ERR = 'GET_ORGANIZATION_LOGS_ERR'
export const RESET_ORGANIZATION_LOGS = 'RESET_ORGANIZATION_LOGS'
