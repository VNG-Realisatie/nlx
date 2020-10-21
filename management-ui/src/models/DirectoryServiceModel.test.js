// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { act } from '@testing-library/react'
import { checkPropTypes } from 'prop-types'
import OutgoingAccessRequestStore from '../stores/OutgoingAccessRequestStore'
import { RootStore } from '../stores'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'
import DirectoryServiceModel, {
  directoryServicePropTypes,
} from './DirectoryServiceModel'

test('createDirectoryService returns an instance', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
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
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
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
    directoryServicesStore: {},
    service: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(directoryService.latestAccessRequest).toBeNull()
})

test('initializing the model with an invalid latest access request', () => {
  expect(() => {
    return new DirectoryServiceModel({
      directoryServicesStore: {},
      service: {
        latestAccessRequest: 'invalid',
      },
    })
  }).toThrowError(
    'the latestAccessRequest should be an instance of the OutgoingAccessRequestModel',
  )
})

test('(re-)fetching the model', async () => {
  const rootStore = new RootStore({
    directoryRepository: {
      getByName: jest.fn().mockResolvedValue({
        state: 'down',
        latestAccessRequest: {
          id: '42',
        },
      }),
    },
  })

  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    service: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(directoryService.state).toBe('up')
  expect(directoryService.latestAccessRequest).toBeNull()

  const directoryServicesStoreFetchSpy = jest.spyOn(
    rootStore.directoryServicesStore,
    'fetch',
  )

  await act(async () => {
    await directoryService.fetch()
  })

  expect(directoryServicesStoreFetchSpy).toHaveBeenCalledWith(directoryService)
})

describe('requesting access to a service', () => {
  it('should not create a new access request an access request is already created', async () => {
    const rootStore = new RootStore({
      accessRequestRepository: {
        createAccessRequest: jest.fn().mockResolvedValue({}),
      },
    })

    const service = {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      latestAccessRequest: new OutgoingAccessRequestModel({
        accessRequestData: {
          state: ACCESS_REQUEST_STATES.CREATED,
        },
      }),
    }

    const directoryService = new DirectoryServiceModel({
      directoryServicesStore: rootStore.directoryServicesStore,
      service: service,
    })

    const spy = jest.spyOn(rootStore.directoryServicesStore, 'requestAccess')
    await directoryService.requestAccess()

    expect(spy).not.toHaveBeenCalled()
  })

  it('should create an access request if it was cancelled before', async () => {
    const rootStore = new RootStore({
      accessRequestRepository: {
        createAccessRequest: jest.fn().mockResolvedValue({}),
      },
    })

    const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
      rootStore,
    })

    const latestAccessRequest = await outgoingAccessRequestStore.updateFromServer(
      {
        id: '42',
        state: ACCESS_REQUEST_STATES.CANCELLED,
      },
    )

    const service = {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      latestAccessRequest: latestAccessRequest,
    }

    const directoryService = new DirectoryServiceModel({
      directoryServicesStore: rootStore.directoryServicesStore,
      service: service,
    })

    const spy = jest.spyOn(rootStore.directoryServicesStore, 'requestAccess')

    await act(async () => {
      await directoryService.requestAccess()
    })

    expect(spy).toHaveBeenCalledWith(directoryService)
  })

  it('should create access request if it was rejected before', async () => {
    const rootStore = new RootStore({
      accessRequestRepository: {
        createAccessRequest: jest.fn().mockResolvedValue({}),
      },
    })

    const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
      rootStore,
    })

    const latestAccessRequest = await outgoingAccessRequestStore.updateFromServer(
      {
        id: '42',
        state: ACCESS_REQUEST_STATES.REJECTED,
      },
    )

    const service = {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      latestAccessRequest: latestAccessRequest,
    }

    const directoryService = new DirectoryServiceModel({
      directoryServicesStore: rootStore.directoryServicesStore,
      service: service,
    })

    const spy = jest.spyOn(rootStore.directoryServicesStore, 'requestAccess')

    await act(async () => {
      await directoryService.requestAccess()
    })

    expect(spy).toHaveBeenCalledWith(directoryService)
  })
})
