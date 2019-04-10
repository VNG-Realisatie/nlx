// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { call, put } from 'redux-saga/effects'
import * as TYPES from './types'

const api = url => fetch(url).then(response => response.json())

export const fetchOrganizationsRequest = () => ({
  type: TYPES.FETCH_ORGANIZATIONS_REQUEST
})

export function* fetchOrganizations() {
  try {
    const organizations = yield call(api, '/api/directory/list-organizations')
    yield put({ type: TYPES.FETCH_ORGANIZATIONS_SUCCESS, data: organizations.organizations })
  } catch (err) {
    console.log(err);
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
