// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import { call, put, retry } from 'redux-saga/effects'

import * as TYPES from './types'
import {
  fetchOrganizations,
  fetchOrganizationLogs,
  fetchIrmaLoginInformation,
  fetchProof,
  getIrmaLoginStatus,
  api,
  apiHandleLoginStatus,
  apiPostWithTextResponse,
  apiWithResponseDetection,
  apiPostWithTextAndJSONAsOutput,
} from './actions'

describe('fetch organizations', () => {
  const organizationsGen = fetchOrganizations()

  it('should call the API', () => {
    expect(organizationsGen.next().value).toEqual(
      call(api, '/api/directory/list-organizations'),
    )
  })

  it('should dispatch the success action when the request succeeds', () => {
    const organizations = { organizations: [] }
    expect(organizationsGen.next(organizations).value).toEqual(
      put({
        type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
        data: organizations.organizations,
      }),
    )
  })
})

describe('fetch IRMA login information', () => {
  const irmaLoginInformationGen = fetchIrmaLoginInformation({
    insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
    insight_irma_endpoint: 'irma_endpoint', // eslint-disable-line camelcase
  })

  it('should call the insights API to get the data subjects', () => {
    expect(irmaLoginInformationGen.next().value).toEqual(
      call(api, 'log_endpoint/getDataSubjects'),
    )
  })

  it('should call the insights API to generate a JWT token for the data subjects', () => {
    const dataSubjectsResponse = { dataSubjects: { foo: '', bar: '' } }
    expect(irmaLoginInformationGen.next(dataSubjectsResponse).value).toEqual(
      call(apiPostWithTextResponse, 'log_endpoint/generateJWT', {
        dataSubjects: ['foo', 'bar'],
      }),
    )
  })

  it('should call the IRMA API to verify the JWT token', () => {
    const jwtTokenResponse = 'dummy-jt-token'
    expect(irmaLoginInformationGen.next(jwtTokenResponse).value).toEqual(
      call(
        apiPostWithTextAndJSONAsOutput,
        'irma_endpoint/session',
        jwtTokenResponse,
      ),
    )
  })

  it('should dispatch the success action when the login flow succeeds', () => {
    const irmaSessionResponse = {
      token: 'test-token',
      sessionPtr: {
        irmaqr: 'irmaqr',
        u: 'irma_endpoint/irma/session/test-session',
      },
    }

    expect(irmaLoginInformationGen.next(irmaSessionResponse).value).toEqual(
      put({
        type: TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS,
        data: {
          dataSubjects: ['foo', 'bar'],
          qrCodeValue:
            '{"irmaqr":"irmaqr","u":"irma_endpoint/irma/session/test-session"}',
          statusUrl: 'irma_endpoint/session/test-token/status',
          proofUrl: 'irma_endpoint/session/test-token/getproof',
          JWT: 'test-token',
        },
      }),
    )
  })
})

describe('get IRMA login status', () => {
  const irmaLoginStatusGen = getIrmaLoginStatus({
    statusUrl: 'status_url',
  })

  it('should make requests to the status URL to get the login status for 1 minute, with a delay of 1 second', () => {
    expect(irmaLoginStatusGen.next().value).toEqual(
      retry(60, 1000, apiHandleLoginStatus, 'status_url'),
    )
  })

  it('should dispatch a success action with the request response as data', () => {
    expect(irmaLoginStatusGen.next('DONE').value).toEqual(
      put({
        type: TYPES.IRMA_LOGIN_REQUEST_SUCCESS,
        data: {
          error: false,
          response: 'DONE',
        },
      }),
    )
  })
})

describe('fetch proof', () => {
  const fetchProofGen = fetchProof({
    proofUrl: 'proof_url',
  })

  it('should get the proof value', () => {
    expect(fetchProofGen.next().value).toEqual(
      call(apiWithResponseDetection, 'proof_url'),
    )
  })

  it('should dispatch a success action with the proof as data', () => {
    expect(fetchProofGen.next({ response: 'the-proof' }).value).toEqual(
      put({
        type: TYPES.FETCH_PROOF_SUCCESS,
        data: { response: 'the-proof' },
      }),
    )
  })
})

describe('fetch organization logs', () => {
  const fetchOrganizationLogsGen = fetchOrganizationLogs({
    proof: 'the_proof',
    insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
  })

  it('should fetch the logs with the provided proof', () => {
    expect(fetchOrganizationLogsGen.next('the_proof').value).toEqual(
      call(apiPostWithTextAndJSONAsOutput, 'log_endpoint/fetch', 'the_proof'),
    )
  })

  it('should dispatch a success action with the request response as data', () => {
    expect(fetchOrganizationLogsGen.next({ response: [] }).value).toEqual(
      put({
        type: TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS,
        data: { response: [] },
      }),
    )
  })

  describe('with pagination', () => {
    it('should pass the pagination params to the fetch endpoint', () => {
      const fetchOrganizationLogsGen = fetchOrganizationLogs({
        proof: 'the_proof',
        insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
        page: 2,
        rowsPerPage: 42,
      })

      expect(fetchOrganizationLogsGen.next('the_proof').value).toEqual(
        call(
          apiPostWithTextAndJSONAsOutput,
          'log_endpoint/fetch?page=2&rowsPerPage=42',
          'the_proof',
        ),
      )
    })

    it('should pass not pass the pagination params to the fetch endpoint if the params are null values', () => {
      const fetchOrganizationLogsGen = fetchOrganizationLogs({
        proof: 'the_proof',
        insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
        page: null,
        rowsPerPage: null,
      })

      expect(fetchOrganizationLogsGen.next('the_proof').value).toEqual(
        call(apiPostWithTextAndJSONAsOutput, 'log_endpoint/fetch', 'the_proof'),
      )
    })
  })
})
