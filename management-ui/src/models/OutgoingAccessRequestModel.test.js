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

test('verifies object as instance', () => {
  const data = { id: 'accessProof' }
  const instance = new OutgoingAccessRequestModel({
    accessRequestData: data,
    outgoingAccessRequestStore: {},
  })

  expect(() => OutgoingAccessRequestModel.verifyInstance(data)).toThrow()
  expect(() =>
    OutgoingAccessRequestModel.verifyInstance(instance),
  ).not.toThrow()
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

test('detect if current state is cancelled or rejected', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    accessRequestData,
    outgoingAccessRequestStore: {},
  })

  expect(accessRequest.isCancelledOrRejected).toBe(false)

  accessRequest.update({
    state: ACCESS_REQUEST_STATES.CANCELLED,
  })

  expect(accessRequest.isCancelledOrRejected).toBe(true)

  accessRequest.update({
    state: ACCESS_REQUEST_STATES.REJECTED,
  })

  expect(accessRequest.isCancelledOrRejected).toBe(true)
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
