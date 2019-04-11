import { organizations } from './index'
import * as TYPES from '../types'

describe('organizations reducer', () => {
  it('should return the initial state', () => {
    expect(organizations(undefined, {})).toEqual([])
  })

  describe('the FETCH_ORGANIZATIONS_SUCCESS action', () => {
    it('should replace the organizations with the data', () => {
      const action = {
        type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
        data: [{
          name: 'foo',
          insight_irma_endpoint: 'irma_endpoint',
          insight_log_endpoint: 'log_endpoint'
        }]
      }
      expect(organizations(undefined, action)).toEqual([{
        name: 'foo',
        insight_irma_endpoint: 'irma_endpoint',
        insight_log_endpoint: 'log_endpoint'
      }])
    })

    it('should filter out the organizations without irma and log endpoint', () => {
      const action = {
        type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
        data: [{
          name: 'valid',
          insight_irma_endpoint: 'irma_endpoint',
          insight_log_endpoint: 'log_endpoint'
        }, {
          name: 'invalid'
        }]
      }
      expect(organizations(undefined, action)).toEqual([{
        name: 'valid',
        insight_irma_endpoint: 'irma_endpoint',
        insight_log_endpoint: 'log_endpoint'
      }])
    })
  })
})
