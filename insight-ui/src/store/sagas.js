import { takeLatest } from 'redux-saga/effects'

import * as TYPES from './types'
import { fetchOrganizations } from './actions'

export default function* () {
  yield takeLatest(TYPES.FETCH_ORGANIZATIONS_REQUEST, fetchOrganizations)
}
