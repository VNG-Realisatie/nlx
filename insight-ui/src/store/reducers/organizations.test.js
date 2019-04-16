import organizationsReducer from './organizations'
import * as TYPES from "../types";

describe('the organizations reducer', () => {
  it('should return the initial state', () => {
    expect(organizationsReducer(undefined, {}))
      .toEqual([])
  })

  describe('handling FETCH_ORGANIZATIONS_SUCCESS', () => {
    it('should return the organizations', () => {
      expect(organizationsReducer(undefined, {
        type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
        data: [ { name: 'foo', insight_irma_endpoint: 'irma_endpoint', insight_log_endpoint: 'log_endpoint'} ]
      }))
        .toEqual([
          { name: 'foo', insight_irma_endpoint: 'irma_endpoint', insight_log_endpoint: 'log_endpoint'}
        ])
    })

    it('should filter out invalid organizations', () => {
      expect(organizationsReducer(undefined, {
        type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
        data: [
          { name: 'foo', insight_irma_endpoint: 'irma_endpoint', insight_log_endpoint: 'log_endpoint'},
          { name: 'foo' },
        ]
      }))
        .toEqual([
          { name: 'foo', insight_irma_endpoint: 'irma_endpoint', insight_log_endpoint: 'log_endpoint'}
        ])
    })
  })
})
