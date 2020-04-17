// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { call, put, retry } from 'redux-saga/effects'
import * as TYPES from './types'

export const api = (url) => fetch(url).then((response) => response.json())

// NOTE: we should use this method as default `api` method, but because of
// the insight-api not returning the correct content-type, we have to
// use a separate method for now.
// Bug report at https://gitlab.com/commonground/nlx/nlx/issues/543
export const apiWithResponseDetection = (url) =>
  fetch(url).then((response) => {
    // via https://stackoverflow.com/a/37121496/363448
    const contentType = response.headers.get('content-type')
    if (contentType && contentType.indexOf('application/json') !== -1) {
      return response.json()
    } else {
      return response.text()
    }
  })

export const apiPost = (url, data) =>
  fetch(url, {
    method: 'POST',
    body: JSON.stringify(data),
  })

export const apiPostWithTextAndJSONAsOutput = (url, data) =>
  fetch(url, {
    headers: {
      'Content-Type': 'text/plain',
    },
    method: 'POST',
    body: data,
  }).then((response) => response.json())

export const apiPostWithTextResponse = (url, data) => {
  return apiPost(url, data).then((response) => response.text())
}
export const apiPostWithJSONResponse = (url, data) => {
  return apiPost(url, data).then((response) => response.json())
}

export const fetchOrganizationsRequest = () => ({
  type: TYPES.FETCH_ORGANIZATIONS_REQUEST,
})

export const resetLoginInformation = () => ({
  type: TYPES.RESET_LOGIN_INFORMATION,
})

export const fetchOrganizationLogsRequest = ({
  insight_log_endpoint, // eslint-disable-line camelcase
  proof,
  page,
  rowsPerPage,
}) => ({
  type: TYPES.FETCH_ORGANIZATION_LOGS_REQUEST,
  data: {
    insight_log_endpoint, // eslint-disable-line camelcase
    proof,
    page,
    rowsPerPage,
  },
})

export const fetchIrmaLoginInformationRequest = ({
  insight_log_endpoint, // eslint-disable-line camelcase
  insight_irma_endpoint, // eslint-disable-line camelcase
}) => ({
  type: TYPES.FETCH_IRMA_LOGIN_INFORMATION_REQUEST,
  data: { insight_log_endpoint, insight_irma_endpoint }, // eslint-disable-line camelcase
})

export function* fetchOrganizations() {
  try {
    const organizations = yield call(api, '/api/directory/list-organizations')
    yield put({
      type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
      data: organizations.organizations,
    })
  } catch (err) {
    console.log(err) // eslint-disable-line no-console
  }
}

function mapDataSubjectsResponseToArray(dataSubjectsResponse) {
  return Object.keys(dataSubjectsResponse.dataSubjects)
}

export function* fetchIrmaLoginInformation({
  insight_log_endpoint, // eslint-disable-line camelcase
  insight_irma_endpoint, // eslint-disable-line camelcase
}) {
  try {
    const dataSubjects = yield call(
      api,
      `${insight_log_endpoint}/getDataSubjects`, // eslint-disable-line camelcase
    )
    const requestSessionJwt = yield call(
      apiPostWithTextResponse,
      `${insight_log_endpoint}/generateJWT`, // eslint-disable-line camelcase
      { dataSubjects: mapDataSubjectsResponseToArray(dataSubjects) },
    )
    const irmaSession = yield call(
      apiPostWithTextAndJSONAsOutput,
      `${insight_irma_endpoint}/session`, // eslint-disable-line camelcase
      requestSessionJwt,
    )
    const sessionToken = irmaSession.token

    yield put({
      type: TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS,
      data: {
        dataSubjects: mapDataSubjectsResponseToArray(dataSubjects),
        qrCodeValue: JSON.stringify(irmaSession.sessionPtr),
        statusUrl: `${insight_irma_endpoint}/session/${sessionToken}/status`, // eslint-disable-line camelcase
        proofUrl: `${insight_irma_endpoint}/session/${sessionToken}/getproof`, // eslint-disable-line camelcase
        JWT: sessionToken,
      },
    })
  } catch (err) {
    console.log(err) // eslint-disable-line no-console
  }
}

export const IRMA_LOGIN_STATUS_INITIALIZED = 'INITIALIZED'
export const IRMA_LOGIN_STATUS_CONNECTED = 'CONNECTED'
export const IRMA_LOGIN_STATUS_CANCELLED = 'CANCELLED'
export const IRMA_LOGIN_STATUS_DONE = 'DONE'
export const IRMA_LOGIN_STATUS_SESSION_UNKNOWN = 'SESSION_UNKNOWN'

export const apiHandleLoginStatus = (url) =>
  api(url)
    .then((response) => {
      if (typeof response === 'string') {
        return response.toUpperCase()
      }

      if (response.error === IRMA_LOGIN_STATUS_SESSION_UNKNOWN) {
        return IRMA_LOGIN_STATUS_SESSION_UNKNOWN
      }

      return JSON.stringify(response)
    })
    .then((response) => {
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
          console.error(
            `Unexpected response '${response}' while getting the login status.`,
          )
          return response
      }
    })

export function* getIrmaLoginStatus({ statusUrl }) {
  try {
    const SECOND = 1000
    const response = yield retry(
      60,
      1 * SECOND,
      apiHandleLoginStatus,
      statusUrl,
    )
    yield put({
      type: TYPES.IRMA_LOGIN_REQUEST_SUCCESS,
      data: {
        error: response !== IRMA_LOGIN_STATUS_DONE,
        response,
      },
    })
  } catch (error) {
    yield put({
      type: TYPES.IRMA_LOGIN_REQUEST_FAILED,
      data: {
        error: true,
        response: error.message,
      },
    })
  }
}

export function* fetchProof({ proofUrl }) {
  try {
    const proof = yield call(apiWithResponseDetection, proofUrl)
    yield put({
      type: TYPES.FETCH_PROOF_SUCCESS,
      data: proof,
    })
  } catch (error) {
    yield put({
      type: TYPES.FETCH_PROOF_FAILED,
      response: error,
    })
  }
}

export function* fetchOrganizationLogs({
  page,
  rowsPerPage,
  proof,
  insight_log_endpoint, // eslint-disable-line camelcase
}) {
  try {
    const url = `${insight_log_endpoint}/fetch` // eslint-disable-line camelcase
    const searchParams = new URLSearchParams()

    if (typeof page !== 'undefined' && page !== null) {
      searchParams.append('page', page)
    }

    if (typeof rowsPerPage !== 'undefined' && page !== null) {
      searchParams.append('rowsPerPage', rowsPerPage)
    }

    const queryString =
      Array.from(searchParams.entries()).length > 0
        ? `?${searchParams.toString()}`
        : ''

    const logs = yield call(
      apiPostWithTextAndJSONAsOutput,
      `${url}${queryString}`,
      proof,
    )
    yield put({ type: TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS, data: logs })
  } catch (error) {
    yield put({
      type: TYPES.FETCH_ORGANIZATION_LOGS_FAILED,
      data: {
        error: true,
        response: error,
      },
    })
  }
}
