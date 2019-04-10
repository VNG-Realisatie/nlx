import { call, put } from 'redux-saga/effects'

import * as TYPES from './types'
import { fetchOrganizations, api } from './actions'

describe('fetch organizations', () => {
  const organizationsGen = fetchOrganizations()

  it('should call the api', () => {
    expect(organizationsGen.next().value)
      .toEqual(call(api, '/api/directory/list-organizations'))
  })

  it('should dispatch the success action when the request succeeds', () => {
    const organizations = {organizations: []}
    expect(organizationsGen.next(organizations).value)
      .toEqual(put({ type: TYPES.FETCH_ORGANIZATIONS_SUCCESS, data: organizations.organizations }))
  })
})
