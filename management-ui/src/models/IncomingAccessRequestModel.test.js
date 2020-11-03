// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
import IncomingAccessRequestModel, {
  createIncomingAccessRequest,
  incomingAccessRequestPropTypes,
} from './IncomingAccessRequestModel'

let incomingAccessRequestStore
let accessRequestData

beforeEach(() => {
  incomingAccessRequestStore = {}

  accessRequestData = {
    id: '1a2B',
    organizationName: 'Organization A',
    serviceName: 'Servicio',
    state: 'RECEIVED',
    createdAt: '2020-10-01T12:00:00Z',
    updatedAt: '2020-10-01T12:00:01Z',
  }
})

test('createIncomingAccessRequest returns an instance', () => {
  const directoryService = createIncomingAccessRequest({
    incomingAccessRequestStore,
    accessRequestData,
  })

  expect(directoryService).toBeInstanceOf(IncomingAccessRequestModel)
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData,
  })

  checkPropTypes(
    incomingAccessRequestPropTypes,
    accessRequest,
    'prop',
    'IncomingAccessRequestModel',
  )

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('approving request handles as expected', async () => {
  const request = deferredPromise()
  const approveIncomingAccessRequest = jest.fn(() => request)

  incomingAccessRequestStore = {
    accessRequestRepository: {
      approveIncomingAccessRequest,
    },
    fetchForService: jest.fn(),
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  accessRequest.approve()
  await request.resolve()

  expect(approveIncomingAccessRequest).toHaveBeenCalled()
  expect(incomingAccessRequestStore.fetchForService).toHaveBeenCalled()
})

test('rejecting request handles as expected', async () => {
  const request = deferredPromise()
  const rejectIncomingAccessRequest = jest.fn(() => request)

  incomingAccessRequestStore = {
    accessRequestRepository: {
      rejectIncomingAccessRequest,
    },
    fetchForService: jest.fn(),
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  accessRequest.reject()
  await request.resolve()

  expect(rejectIncomingAccessRequest).toHaveBeenCalled()
  expect(incomingAccessRequestStore.fetchForService).toHaveBeenCalled()
})

test('set an error', async () => {
  const approveIncomingAccessRequest = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))
  const rejectIncomingAccessRequest = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  incomingAccessRequestStore = {
    accessRequestRepository: {
      approveIncomingAccessRequest,
      rejectIncomingAccessRequest,
    },
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  await accessRequest.approve()
  expect(accessRequest.error).toEqual('arbitrary error')

  accessRequest.error = ''

  await accessRequest.reject()
  expect(accessRequest.error).toEqual('arbitrary error')
})

test('returns proper isResolved value', () => {
  const accessRequest = new IncomingAccessRequestModel({
    accessRequestData,
  })
  expect(accessRequest.isResolved).toBe(false)

  accessRequest.update({ state: 'RECEIVED' })
  expect(accessRequest.isResolved).toBe(false)

  accessRequest.update({ state: 'FAILED' })
  expect(accessRequest.isResolved).toBe(true)

  accessRequest.update({ state: 'ACCEPTED' })
  expect(accessRequest.isResolved).toBe(true)

  accessRequest.update({ state: 'REJECTED' })
  expect(accessRequest.isResolved).toBe(true)
})
