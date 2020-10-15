// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { act } from '@testing-library/react'
import { checkPropTypes } from 'prop-types'
import * as outgoingAccessRequestLib from './OutgoingAccessRequestModel'
import DirectoryServiceModel, {
  directoryServicePropTypes,
  createDirectoryService,
} from './DirectoryServiceModel'
import { ACCESS_REQUEST_STATES } from './OutgoingAccessRequestModel'

let createAccessRequestInstanceMock

beforeEach(() => {
  createAccessRequestInstanceMock = jest
    .spyOn(outgoingAccessRequestLib, 'createAccessRequestInstance')
    .mockImplementation((obj) => obj)
})

afterEach(() => {
  createAccessRequestInstanceMock.mockReset()
})

test('createDirectoryService returns an instance', () => {
  const directoryService = createDirectoryService({
    store: {},
    service: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })
  expect(directoryService).toBeInstanceOf(DirectoryServiceModel)
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const directoryService = createDirectoryService({
    store: {},
    service: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  checkPropTypes(
    directoryServicePropTypes,
    directoryService,
    'prop',
    'DirectoryServiceModel',
  )

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('initializing the model', async () => {
  const directoryService = new DirectoryServiceModel({
    store: {},
    service: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(directoryService.latestAccessRequest).toBeNull()
})

test('(re-)fetching the model', async () => {
  const store = {
    directoryRepository: {
      getByName: jest.fn().mockResolvedValue({
        state: 'down',
        latestAccessRequest: {
          id: '42',
        },
      }),
    },
  }

  const directoryService = new DirectoryServiceModel({
    store,
    service: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(directoryService.state).toBe('up')
  expect(directoryService.latestAccessRequest).toBeNull()

  await act(async () => {
    await directoryService.fetch()
  })

  expect(store.directoryRepository.getByName).toHaveBeenCalledWith(
    'Organization',
    'Service',
  )

  expect(directoryService.state).toBe('down')
  expect(directoryService.latestAccessRequest).toEqual({
    id: '42',
  })
})

describe('requesting access to a service', () => {
  it('should not create a new access request an access request is already created', () => {
    const service = {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      latestAccessRequest: {
        state: ACCESS_REQUEST_STATES.CREATED,
        send: jest.fn(),
      },
    }

    const directoryService = new DirectoryServiceModel({
      store: {},
      service: service,
    })

    const sendHandlerSpy = jest.fn()
    createAccessRequestInstanceMock.mockImplementation((obj) => ({
      ...obj,
      send: sendHandlerSpy,
    }))

    directoryService.requestAccess()

    expect(sendHandlerSpy).not.toHaveBeenCalled()
  })

  it('should create access request if it was cancelled before', async () => {
    const service = {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      latestAccessRequest: {
        state: ACCESS_REQUEST_STATES.CANCELLED,
      },
    }

    const directoryService = new DirectoryServiceModel({
      store: {},
      service: service,
    })

    const sendHandlerSpy = jest.fn()
    createAccessRequestInstanceMock.mockImplementation((obj) => ({
      ...obj,
      send: sendHandlerSpy,
    }))

    directoryService.requestAccess()

    expect(sendHandlerSpy).not.toHaveBeenCalled()
  })

  it('should create access request if it was rejected before', async () => {
    const service = {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      latestAccessRequest: {
        state: ACCESS_REQUEST_STATES.REJECTED,
      },
    }

    const directoryService = new DirectoryServiceModel({
      store: {},
      service: service,
    })

    const sendHandlerSpy = jest.fn()
    createAccessRequestInstanceMock.mockImplementation((obj) => ({
      ...obj,
      send: sendHandlerSpy,
    }))

    await directoryService.requestAccess()

    expect(sendHandlerSpy).toHaveBeenCalled()
  })
})
