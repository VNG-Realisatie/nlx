// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//
import OutgoingAccessRequestSyncErrorStore from './OutgoingAccessRequestSyncErrorStore'
import OutgoingAccessRequestSyncErrorModel from './models/OutgoingAccessRequestSyncErrorModel'

test('load from sync response', () => {
  const syncErrorStore = new OutgoingAccessRequestSyncErrorStore()

  expect(() => {
    syncErrorStore.loadFromSyncResponse('', '')
  }).toThrowError('please provide the JSON response')

  syncErrorStore.loadFromSyncResponse('00000000000000000001', 'my-service', {
    message: 'service_provider_no_organization_inway_specified',
  })

  let got = syncErrorStore.getForService('00000000000000000001', 'my-service')

  expect(got).toBeInstanceOf(OutgoingAccessRequestSyncErrorModel)
  expect(got.message).toEqual(
    'The organization has not specified an organization Inway. We are unable to retrieve the current state of your access requests.',
  )

  syncErrorStore.clearForService('00000000000000000001', 'my-service')

  got = syncErrorStore.getForService('00000000000000000001', 'my-service')
  expect(got).toBeUndefined()
})

test('load from sync all response', () => {
  const syncErrorStore = new OutgoingAccessRequestSyncErrorStore()

  expect(() => {
    syncErrorStore.loadFromSyncAllResponse()
  }).toThrowError('please provide the JSON response')

  syncErrorStore.loadFromSyncAllResponse({
    details: [
      {
        metadata: {
          '00000000000000000001': 'internal_error',
          '00000000000000000002':
            'service_provider_organization_inway_unreachable',
        },
      },
    ],
  })

  let got = syncErrorStore.getForService('00000000000000000001', 'arbitrary')

  expect(got).toBeInstanceOf(OutgoingAccessRequestSyncErrorModel)
  expect(got.message).toEqual(
    'Internal error while trying to retrieve the current state of your access request. Please consult your system administrator.',
  )

  got = syncErrorStore.getForService('00000000000000000002', 'arbitrary')

  expect(got).toBeInstanceOf(OutgoingAccessRequestSyncErrorModel)
  expect(got.message).toEqual(
    'The organization Inway of this organization is unreachable. We are unable to retrieve the current state of your access requests.',
  )

  syncErrorStore.clearAll()
  got = syncErrorStore.getForService('00000000000000000002', 'arbitrary')
  expect(got).toBeUndefined()
})
