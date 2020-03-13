// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import * as TYPES from '../types'
import proofReducer from './proof'

describe('the proof reducer', () => {
  it('should return the initial state', () => {
    expect(proofReducer(undefined, {})).toEqual({
      loaded: false,
      error: false,
      value: null,
      response: null,
    })
  })

  it('should handle FETCH_PROOF_SUCCESS', () => {
    expect(
      proofReducer(undefined, {
        type: TYPES.FETCH_PROOF_SUCCESS,
        data: 'xyz',
      }),
    ).toEqual({
      loaded: true,
      error: false,
      value: 'xyz',
      message: null,
    })
  })

  it('should handle FETCH_PROOF_FAILED', () => {
    expect(
      proofReducer(undefined, {
        type: TYPES.FETCH_PROOF_FAILED,
        data: {
          response: 'an error occured',
        },
      }),
    ).toEqual({
      loaded: true,
      error: true,
      value: null,
      message: 'an error occured',
    })
  })

  it('should handle RESET_LOGIN_INFORMATION', () => {
    expect(
      proofReducer(undefined, {
        type: TYPES.RESET_LOGIN_INFORMATION,
      }),
    ).toEqual({
      loaded: false,
      error: false,
      value: null,
      response: null,
    })
  })
})
