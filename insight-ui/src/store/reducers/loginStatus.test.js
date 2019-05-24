import loginStatusReducer from './loginStatus'
import * as TYPES from '../types'

describe('the loginStatus reducer', () => {
  it('should return the initial state', () => {
    expect(loginStatusReducer(undefined, {}))
      .toEqual(null)
  })

  it('should handle IRMA_LOGIN_REQUEST_SUCCESS', () => {
    expect(loginStatusReducer(undefined, {
      type: TYPES.IRMA_LOGIN_REQUEST_SUCCESS,
      data: 'foo'
    }))
      .toEqual('foo')
  })

  it('should handle IRMA_LOGIN_REQUEST_FAILED', () => {
    expect(loginStatusReducer(undefined, {
      type: TYPES.IRMA_LOGIN_REQUEST_FAILED,
      data: 'foo'
    }))
      .toEqual('foo')
  })

  it('should handle RESET_LOGIN_INFORMATION', () => {
    expect(loginStatusReducer(undefined, {
      type: TYPES.RESET_LOGIN_INFORMATION
    }))
      .toEqual(null)
  })
})
