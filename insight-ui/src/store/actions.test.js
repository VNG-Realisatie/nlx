import { call, put } from 'redux-saga/effects'

import * as TYPES from './types'
import { fetchOrganizations, fetchIrmaLoginInformation, api, apiPostWithJSONResponse, apiPostWithTextResponse } from './actions'

describe('fetch organizations', () => {
  const organizationsGen = fetchOrganizations()

  it('should call the API', () => {
    expect(organizationsGen.next().value)
      .toEqual(call(api, '/api/directory/list-organizations'))
  })

  it('should dispatch the success action when the request succeeds', () => {
    const organizations = {organizations: []}
    expect(organizationsGen.next(organizations).value)
      .toEqual(put({ type: TYPES.FETCH_ORGANIZATIONS_SUCCESS, data: organizations.organizations }))
  })
})

describe('fetch IRMA login information', () => {
  const irmaLoginInformationGen = fetchIrmaLoginInformation({
    insight_log_endpoint: 'log_endpoint',
    insight_irma_endpoint: 'irma_endpoint'
  })

  it('should call the insights API to get the data subjects', () => {
    expect(irmaLoginInformationGen.next().value)
      .toEqual(call(api, 'log_endpoint/getDataSubjects'))
  })

  it('should call the insights API to generate a JWT token for the data subjects', () => {
    const dataSubjectsResponse = {dataSubjects: {foo: '', bar: ''}}
    expect(irmaLoginInformationGen.next(dataSubjectsResponse).value)
      .toEqual(call(apiPostWithTextResponse, 'log_endpoint/generateJWT', {dataSubjects: ['foo', 'bar']}))
  })

  it('should call the IRMA API to verify the JWT token', () => {
    const jwtTokenResponse = 'dummy-jt-token'
    expect(irmaLoginInformationGen.next(jwtTokenResponse).value)
      .toEqual(call(apiPostWithJSONResponse, 'irma_endpoint/api/v2/verification/', jwtTokenResponse))
  })

  it('should dispatch the success action when the login flow succeeds', () => {
    const jwtVerificationResponse = {
      u: 'u',
      v: 'v',
      vmax: 'vmax',
      irmaqr: 'irmaqr',
    }

    expect(irmaLoginInformationGen.next(jwtVerificationResponse).value)
      .toEqual(put({
        type: TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS,
        data: {
          dataSubjects: ['foo', 'bar'],
          qrCodeValue: '{"u":"irma_endpoint/api/v2/verification/u","v":"v","vmax":"vmax","irmaqr":"irmaqr"}',
          statusUrl: 'irma_endpoint/api/v2/verification/u/status',
          proofUrl: 'irma_endpoint/api/v2/verification/u/getproof',
          JWT: 'u'
        }
      }))
  })
})
