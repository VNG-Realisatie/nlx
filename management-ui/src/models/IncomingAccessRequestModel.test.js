// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
import IncomingAccessRequestModel, {
  createIncomingAccessRequest,
  incomingAccessRequestPropTypes,
} from './IncomingAccessRequestModel'

let store
let accessRequestData
let accessRequestRepository

beforeEach(() => {
  store = {
    fetchIncomingAccessRequests: jest.fn(),
  }

  accessRequestData = {
    id: '1a2B',
    organizationName: 'Organization A',
    serviceName: 'Servicio',
    state: 'RECEIVED',
    createdAt: '2020-08-25T13:30:43.480155Z',
    updatedAt: '2020-08-25T13:30:43.480155Z',
  }

  accessRequestRepository = {
    approveIncomingAccessRequest: jest.fn(),
  }
})

test('createIncomingAccessRequest returns an instance', () => {
  const directoryService = createIncomingAccessRequest({
    store,
    accessRequestData,
    accessRequestRepository,
  })

  expect(directoryService).toBeInstanceOf(IncomingAccessRequestModel)
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const accessRequest = new IncomingAccessRequestModel({
    json: accessRequestData,
    accessRequestRepository,
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
  accessRequestRepository = {
    approveIncomingAccessRequest: jest.fn(() => request),
  }

  const accessRequest = new IncomingAccessRequestModel({
    store,
    accessRequestData,
    accessRequestRepository,
  })

  accessRequest.approve()
  await request.resolve()

  expect(store.fetchIncomingAccessRequests).toHaveBeenCalled()
})

test('on error it will reset state to RECEIVED', async () => {
  accessRequestRepository = {
    approveIncomingAccessRequest: jest
      .fn()
      .mockRejectedValue(new Error('arbitrary error')),
  }

  const accessRequest = new IncomingAccessRequestModel({
    store,
    accessRequestData,
    accessRequestRepository,
  })

  await accessRequest.approve()

  expect(accessRequest.error).toEqual('arbitrary error')
})
