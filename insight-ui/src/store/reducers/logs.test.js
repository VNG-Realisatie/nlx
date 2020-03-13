// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

import * as TYPES from '../types'
import logsReducer from './logs'

describe('the logs reducer', () => {
  it('should return the initial state', () => {
    expect(logsReducer(undefined, {})).toEqual({
      records: [],
      rowCount: 0,
      rowsPerPage: 0,
    })
  })

  it('should handle FETCH_ORGANIZATION_LOGS_SUCCESS', () => {
    expect(
      logsReducer(undefined, {
        type: TYPES.FETCH_ORGANIZATION_LOGS_SUCCESS,
        data: {
          rowCount: 11,
          rowsPerPage: 5,
          records: [
            {
              data: {
                'doelbinding-data-elements': 'a,b',
                'doelbinding-process-id': 'process id',
                'doelbinding-application-id': 'application',
              },
              created: '2019-05-17T07:22:49.996932Z',
              source_organization: 'source organization',
              destination_organization: 'destination organization',
              'logrecord-id': 'id',
            },
          ],
        },
      }),
    ).toEqual({
      rowCount: 11,
      rowsPerPage: 5,
      records: [
        {
          id: 'id',
          subjects: ['a', 'b'],
          requestedBy: 'source organization',
          requestedAt: 'destination organization',
          application: 'application',
          reason: 'process id',
          date: new Date(Date.UTC(2019, 4, 17, 7, 22, 49, 996)),
        },
      ],
    })
  })

  it('should handle RESET_LOGIN_INFORMATION', () => {
    expect(
      logsReducer(undefined, {
        type: TYPES.RESET_LOGIN_INFORMATION,
      }),
    ).toEqual({
      records: [],
      rowCount: 0,
      rowsPerPage: 0,
    })
  })
})
