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
        records: [{
          data: {
            'doelbinding-data-elements': 'a,b',
            'doelbinding-process-id': 'process id',
            'doelbinding-application-id': 'application'
          },
          created: '2019-05-17T07:22:49.996932Z',
          source_organization: 'source organization',
          destination_organization: 'destination organization',
          'logrecord-id': 'id'
        }]
      }
    }))
      .toEqual([{
        id: 'id',
        subjects: ['a', 'b'],
        requestedBy: 'source organization',
        requestedAt: 'destination organization',
        application: 'application',
        reason: 'process id',
        date: new Date(2019, 4, 17, 9, 22, 49, 996)
      }])
  })
})
