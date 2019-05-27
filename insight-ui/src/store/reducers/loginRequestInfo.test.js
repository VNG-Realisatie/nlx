// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import loginRequestInfoReducer from './loginRequestInfo'
import * as TYPES from '../types'

describe('the loginRequestInfo reducer', () => {
  it('should return the initial state', () => {
    expect(loginRequestInfoReducer(undefined, {}))
      .toEqual({})
  })

  it('should handle FETCH_IRMA_LOGIN_INFORMATION_SUCCESS', () => {
    expect(loginRequestInfoReducer(undefined, {
      type: TYPES.FETCH_IRMA_LOGIN_INFORMATION_SUCCESS,
      data: 'foo'
    }))
      .toEqual('foo')
  })

  it('should handle RESET_LOGIN_INFORMATION', () => {
    expect(loginRequestInfoReducer(undefined, {
      type: TYPES.RESET_LOGIN_INFORMATION
    }))
      .toEqual({})
  })
})
