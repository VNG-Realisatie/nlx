// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import * as TYPES from '../types'
import organizationsReducer from './organizations'

describe('the organizations reducer', () => {
  it('should return the initial state', () => {
    expect(organizationsReducer(undefined, {})).toEqual([])
  })

  describe('handling FETCH_ORGANIZATIONS_SUCCESS', () => {
    it('should return the organizations', () => {
      expect(
        organizationsReducer(undefined, {
          type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
          data: [
            {
              name: 'foo',
              insight_irma_endpoint: 'irma_endpoint', // eslint-disable-line camelcase
              insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
            },
          ],
        }),
      ).toEqual([
        {
          name: 'foo',
          insight_irma_endpoint: 'irma_endpoint', // eslint-disable-line camelcase
          insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
        },
      ])
    })

    it('should filter out invalid organizations', () => {
      expect(
        organizationsReducer(undefined, {
          type: TYPES.FETCH_ORGANIZATIONS_SUCCESS,
          data: [
            {
              name: 'foo',
              insight_irma_endpoint: 'irma_endpoint', // eslint-disable-line camelcase
              insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
            },
            { name: 'foo' },
          ],
        }),
      ).toEqual([
        {
          name: 'foo',
          insight_irma_endpoint: 'irma_endpoint', // eslint-disable-line camelcase
          insight_log_endpoint: 'log_endpoint', // eslint-disable-line camelcase
        },
      ])
    })
  })
})
