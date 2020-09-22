// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
// So we can mock AND spy { createAccessRequestInstance }
import * as outgoingAccessRequestLib from './OutgoingAccessRequestModel'
import DirectoryServiceModel, {
  directoryServicePropTypes,
  createDirectoryService,
} from './DirectoryServiceModel'

let store
let service
let createAccessRequestInstanceMock

beforeEach(() => {
  store = {}
  service = {
    organizationName: 'Organization',
    serviceName: 'Service',
    state: 'up',
    apiSpecificationType: 'API',
  }

  createAccessRequestInstanceMock = jest
    .spyOn(outgoingAccessRequestLib, 'createAccessRequestInstance')
    .mockImplementation((obj) => obj)
})

afterEach(() => {
  createAccessRequestInstanceMock.mockReset()
})

test('createDirectoryService returns an instance', () => {
  const directoryService = createDirectoryService({ store, service })
  expect(directoryService).toBeInstanceOf(DirectoryServiceModel)
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const directoryService = new DirectoryServiceModel({ store, service })

  checkPropTypes(
    directoryServicePropTypes,
    directoryService,
    'prop',
    'DirectoryServiceModel',
  )

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('fetches data', async () => {
  const request = deferredPromise()
  store = {
    directoryRepository: {
      getByName: jest.fn(() => request),
    },
  }

  const directoryService = new DirectoryServiceModel({ store, service })

  expect(directoryService.latestAccessRequest).toBeNull()

  directoryService.fetch()

  expect(store.directoryRepository.getByName).toHaveBeenCalled()

  await request.resolve({
    state: 'down',
    latestAccessRequest: {},
  })

  expect(directoryService.state).toBe('down')
  expect(directoryService.latestAccessRequest).toEqual({})
})

describe('creating access request', () => {
  it('should fail silently when an access request is already open', () => {
    service.latestAccessRequest = {
      isOpen: true,
    }

    const directoryService = new DirectoryServiceModel({ store, service })

    // createAccessRequestInstance is called in constructor, so clear it
    createAccessRequestInstanceMock.mockClear()

    directoryService.requestAccess()
    expect(createAccessRequestInstanceMock).not.toHaveBeenCalled()
  })

  it('should send access request', async () => {
    service.latestAccessRequest = {
      isOpen: false,
    }

    const directoryService = new DirectoryServiceModel({ store, service })

    const send = jest.fn()
    // Rebuild implmentation to add `send`
    // Note: send is yielded, so we need to await `requestAccess`
    createAccessRequestInstanceMock.mockImplementation((obj) => ({
      ...obj,
      send,
    }))

    await directoryService.requestAccess()

    expect(send).toHaveBeenCalled()
  })
})
