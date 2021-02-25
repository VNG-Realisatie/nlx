// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import IncomingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './IncomingAccessRequestModel'

let accessRequestData

beforeEach(() => {
  accessRequestData = {
    id: '1a2B',
    organizationName: 'Organization A',
    serviceName: 'Servicio',
    state: 'RECEIVED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:01Z',
  }
})

test('initialize the model', () => {
  const model = new IncomingAccessRequestModel({
    accessRequestData: accessRequestData,
  })

  expect(model).toBeInstanceOf(IncomingAccessRequestModel)
})

test('approving request handles as expected', async () => {
  const incomingAccessRequestStore = {
    approveAccessRequest: jest.fn(),
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  accessRequest.approve()

  expect(incomingAccessRequestStore.approveAccessRequest).toHaveBeenCalledWith(
    accessRequest,
  )
})

test('rejecting request handles as expected', async () => {
  const incomingAccessRequestStore = {
    rejectAccessRequest: jest.fn(),
    fetchForService: jest.fn(),
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  accessRequest.reject()

  expect(incomingAccessRequestStore.rejectAccessRequest).toHaveBeenCalled()
})

test('when approving or rejecting fails', async () => {
  jest.spyOn(global.console, 'error').mockImplementation(() => {})

  const incomingAccessRequestStore = {
    approveAccessRequest: jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error')),
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  await expect(accessRequest.approve()).rejects.toThrow('arbitrary error')

  global.console.error.mockRestore()
})

test('when rejecting fails', async () => {
  jest.spyOn(global.console, 'error').mockImplementation(() => {})

  const incomingAccessRequestStore = {
    rejectAccessRequest: jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error')),
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  await expect(accessRequest.reject()).rejects.toThrow('arbitrary error')

  global.console.error.mockRestore()
})

test('returns proper isResolved value', () => {
  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData,
  })
  expect(accessRequest.isResolved).toBe(false)

  accessRequest.update({ state: ACCESS_REQUEST_STATES.RECEIVED })
  expect(accessRequest.isResolved).toBe(false)

  accessRequest.update({ state: ACCESS_REQUEST_STATES.FAILED })
  expect(accessRequest.isResolved).toBe(true)

  accessRequest.update({ state: ACCESS_REQUEST_STATES.APPROVED })
  expect(accessRequest.isResolved).toBe(true)

  accessRequest.update({ state: ACCESS_REQUEST_STATES.REJECTED })
  expect(accessRequest.isResolved).toBe(true)
})
