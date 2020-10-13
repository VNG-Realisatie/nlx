// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
import OutgoingAccessRequestModel, {
  createAccessRequestInstance,
  outgoingAccessRequestPropTypes,
} from './OutgoingAccessRequestModel'

jest.mock('../domain/access-request-repository', (obj) => obj)

let serviceData
let accessRequestJson
let accessRequestRepository

beforeEach(() => {
  serviceData = {
    organizationName: 'Organization',
    serviceName: 'Service',
  }

  accessRequestJson = {
    ...serviceData,
    id: 'abcd',
    createdAt: 'datetime1',
    updatedAt: 'datetime2',
  }

  accessRequestRepository = {
    createAccessRequest: jest.fn(),
  }
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const accessRequest = new OutgoingAccessRequestModel({
    json: accessRequestJson,
    accessRequestRepository,
  })

  checkPropTypes(
    outgoingAccessRequestPropTypes,
    accessRequest,
    'prop',
    'OutgoingAccessRequestModel',
  )

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('createAccessRequestInstance creates an instance', () => {
  expect(createAccessRequestInstance(serviceData)).toBeInstanceOf(
    OutgoingAccessRequestModel,
  )
})

test('sending a request', async () => {
  const request = deferredPromise()
  accessRequestRepository = {
    createAccessRequest: jest.fn(() => request),
  }

  const accessRequest = new OutgoingAccessRequestModel({
    json: serviceData,
    accessRequestRepository,
  })

  expect(accessRequest.state).toBe('')

  accessRequest.send()

  expect(accessRequest.state).toBe('CREATED')
  expect(accessRequestRepository.createAccessRequest).toHaveBeenCalled()

  await request.resolve(accessRequestJson)

  expect(accessRequest.id).toBe('abcd')
})

test('update should ignore properties that do not belong on object', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    json: accessRequestJson,
    accessRequestRepository,
  })

  accessRequest.update({ yada: 'blada' })

  expect('yada' in accessRequest).toBe(false)
})
