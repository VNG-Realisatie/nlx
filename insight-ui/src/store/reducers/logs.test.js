import logsReducer from './logs'
import * as TYPES from '../types'

describe('the logs reducer', () => {
  it('should return the initial state', () => {
    expect(logsReducer(undefined, {}))
      .toEqual([])
  })


  it('should handle FETCH_ORGANIZATION_LOGS_SUCCESS', () => {
    expect(logsReducer(undefined, {
      type: TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS,
      data: {
        records: ['foo', 'bar']
      }
    }))
      .toEqual([ 'foo', 'bar' ])
  })
})
