// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { takeLatest } from 'redux-saga/effects'

import * as TYPES from './types'
import {
  fetchOrganizations,
  fetchOrganizationLogs,
  fetchIrmaLoginInformation,
  getIrmaLoginStatus,
  fetchProof,
} from './actions'

function* fetchProofOnLoginSuccess(action) {
  yield takeLatest(TYPES.IRMA_LOGIN_REQUEST_SUCCESS, () =>
    fetchProof(action.data),
  )
}

export default function* () {
  yield takeLatest(TYPES.FETCH_ORGANIZATIONS_REQUEST, fetchOrganizations)
  yield takeLatest(TYPES.FETCH_ORGANIZATION_LOGS_REQUEST, (action) =>
    fetchOrganizationLogs(action.data),
  )
  yield takeLatest(TYPES.FETCH_IRMA_LOGIN_INFORMATION_REQUEST, (action) =>
    fetchIrmaLoginInformation(action.data),
  )
  yield takeLatest(TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS, (action) =>
    getIrmaLoginStatus(action.data),
  )
  yield takeLatest(
    TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS,
    fetchProofOnLoginSuccess,
  )
}
