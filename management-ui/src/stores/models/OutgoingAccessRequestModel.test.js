// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'

let accessRequestData

beforeEach(() => {
  accessRequestData = {
    id: 'abcd',
    organization: {
      serialNumber: '00000000000000000001',
      name: 'Organization',
    },
    serviceName: 'Service',
    state: ACCESS_REQUEST_STATES.RECEIVED,
    createdAt: '2020-10-01',
    updatedAt: '2020-10-02',
    errorDetails: {
      cause: 'the cause of an error',
      stackTrace: 'a stack trace of the error',
    },
  }
})

test('should properly construct object', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData,
    outgoingAccessRequestStore: {},
  })

  expect(accessRequest.id).toBe(accessRequestData.id)
  expect(accessRequest.organization.name).toBe(
    accessRequestData.organization.name,
  )
  expect(accessRequest.serviceName).toBe(accessRequestData.serviceName)
  expect(accessRequest.state).toBe(accessRequestData.state)
  expect(accessRequest.createdAt).toEqual(new Date(accessRequestData.createdAt))
  expect(accessRequest.updatedAt).toEqual(new Date(accessRequestData.updatedAt))
  expect(accessRequest.errorDetails.cause).toEqual('the cause of an error')
  expect(accessRequest.errorDetails.stackTrace).toEqual(
    'a stack trace of the error',
  )
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
