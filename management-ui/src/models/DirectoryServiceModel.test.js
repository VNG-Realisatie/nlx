// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { act } from '@testing-library/react'
import { checkPropTypes } from 'prop-types'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'
import DirectoryServiceModel, {
  createDirectoryService,
  directoryServicePropTypes,
} from './DirectoryServiceModel'

test('createDirectoryService returns an instance', () => {
  const directoryService = createDirectoryService({
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
  const directoryService = createDirectoryService({
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

test('(re-)fetching the model', async () => {
  const directoryServiceStore = {
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
    directoryServiceStore: directoryServiceStore,
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

  expect(
    directoryServiceStore.directoryRepository.getByName,
  ).toHaveBeenCalledWith('Organization', 'Service')

  expect(directoryService.state).toBe('down')
  expect(directoryService.latestAccessRequest).toEqual(
    new OutgoingAccessRequestModel({
      accessRequestData: {
        id: '42',
      },
    }),
  )
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

  it('should create access request if it was cancelled before', async () => {
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
          state: ACCESS_REQUEST_STATES.CANCELLED,
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

  it('should create access request if it was rejected before', async () => {
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
          state: ACCESS_REQUEST_STATES.REJECTED,
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

    expect(accessRequestRepository.createAccessRequest).toHaveBeenCalledWith({
      organizationName: 'Organization',
      serviceName: 'Service',
    })
  })
})
