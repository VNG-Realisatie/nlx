// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'

jest.mock('../domain/access-request-repository', (obj) => obj)

let accessRequestData

beforeEach(() => {
  accessRequestData = {
    id: 'abcd',
    organizationName: 'Organization',
    serviceName: 'Service',
    state: ACCESS_REQUEST_STATES.RECEIVED,
    createdAt: '2020-10-01',
    updatedAt: '2020-10-02',
  }
})

test('should properly construct object', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData,
    outgoingAccessRequestStore: {},
  })

  expect(accessRequest.id).toBe(accessRequestData.id)
  expect(accessRequest.organizationName).toBe(
    accessRequestData.organizationName,
  )
  expect(accessRequest.serviceName).toBe(accessRequestData.serviceName)
  expect(accessRequest.state).toBe(accessRequestData.state)
  expect(accessRequest.createdAt).toEqual(new Date(accessRequestData.createdAt))
  expect(accessRequest.updatedAt).toEqual(new Date(accessRequestData.updatedAt))
})

test('calling retry should pass instance to store', () => {
  const storeRetryMock = jest.fn()
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData,
    outgoingAccessRequestStore: { retry: storeRetryMock },
  })

  accessRequest.retry()

  expect(storeRetryMock).toHaveBeenCalledWith(accessRequest)
})
