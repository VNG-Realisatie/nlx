// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

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
  const approveAccessRequest = jest.fn().mockResolvedValue(null)

  incomingAccessRequestStore = {
    approveAccessRequest,
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  accessRequest.approve()

  expect(approveAccessRequest).toHaveBeenCalled()
})

test('rejecting request handles as expected', async () => {
  const rejectAccessRequest = jest.fn().mockResolvedValue(null)

  incomingAccessRequestStore = {
    rejectAccessRequest,
    fetchForService: jest.fn(),
  }

  const accessRequest = new IncomingAccessRequestModel({
    incomingAccessRequestStore,
    accessRequestData,
  })

  accessRequest.reject()

  expect(rejectAccessRequest).toHaveBeenCalled()
})

test('when approving or rejecting fails', async () => {
  const approveAccessRequest = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))
  const rejectAccessRequest = jest
    .fn()
    .mockRejectedValue(new Error('arbitrary error'))

  incomingAccessRequestStore = {
    approveAccessRequest,
    rejectAccessRequest,
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
