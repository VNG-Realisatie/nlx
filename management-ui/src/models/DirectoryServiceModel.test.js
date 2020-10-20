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
    directoryServiceStore: {},
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
    directoryServiceStore: {},
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
    directoryServiceStore: {},
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
      directoryServiceStore: {},
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
    directoryServiceStore: rootStore.directoryStore,
    service: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(directoryService.state).toBe('up')
  expect(directoryService.latestAccessRequest).toBeNull()

  const directoryStoreFetchSpy = jest.spyOn(rootStore.directoryStore, 'fetch')

  await act(async () => {
    await directoryService.fetch()
  })

  expect(directoryStoreFetchSpy).toHaveBeenCalledWith(directoryService)
})

describe('requesting access to a service', () => {
  it('should not create a new access request an access request is already created', async () => {
    const accessRequestRepository = {
      createAccessRequest: jest.fn().mockResolvedValue({}),
    }

    const service = {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      latestAccessRequest: new OutgoingAccessRequestModel({
        accessRequestData: {
          state: ACCESS_REQUEST_STATES.CREATED,
        },
        accessRequestRepository: accessRequestRepository,
      }),
    }

    const directoryService = new DirectoryServiceModel({
      directoryServiceStore: {},
      service: service,
      accessRequestRepository: accessRequestRepository,
    })

    await directoryService.requestAccess()

    expect(
      accessRequestRepository.createAccessRequest,
    ).not.toHaveBeenCalledWith()
  })

  it('should create an access request if it was cancelled before', async () => {
    const directoryServiceStore = {
      requestAccess: jest.fn().mockResolvedValue(),
    }

    const outgoingAccessRequestStore = new OutgoingAccessRequestStore({
      rootStore: {},
    })

    const latestAccessRequest = await outgoingAccessRequestStore.loadOutgoingAccessRequest(
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
      directoryServiceStore: directoryServiceStore,
      service: service,
    })

    await directoryService.requestAccess()
    expect(directoryServiceStore.requestAccess).toHaveBeenCalledWith(
      directoryService,
    )
  })

  it('should create access request if it was rejected before', async () => {
    const directoryServiceStore = {
      requestAccess: jest.fn().mockResolvedValue(),
    }

    const outgoingAccessRequestsStore = new OutgoingAccessRequestStore({
      rootStore: {},
    })

    const latestAccessRequest = await outgoingAccessRequestsStore.loadOutgoingAccessRequest(
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
      directoryServiceStore: directoryServiceStore,
      service: service,
    })

    await directoryService.requestAccess()
    expect(directoryServiceStore.requestAccess).toHaveBeenCalledWith(
      directoryService,
    )
  })
})
