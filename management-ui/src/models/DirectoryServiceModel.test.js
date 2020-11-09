// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'
import { RootStore } from '../stores'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'
import DirectoryServiceModel, {
  directoryServicePropTypes,
} from './DirectoryServiceModel'
import AccessProofModel from './AccessProofModel'

test('createDirectoryService returns an instance', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
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
    serviceData: {
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

test('initializing and updating the model', () => {
  const serviceData = {
    organizationName: 'Organization',
    serviceName: 'Service',
    state: 'up',
    apiSpecificationType: 'API',
    latestAccessRequest: {
      state: ACCESS_REQUEST_STATES.CREATED,
    },
    latestAccessProof: {
      id: 'abc',
    },
  }

  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData,
  })

  expect(directoryService.latestAccessRequest).toBeNull()
  expect(directoryService.latestAccessProof).toBeNull()

  const latestAccessRequest = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: {},
    accessRequestData: serviceData.latestAccessRequest,
  })

  const latestAccessProof = new AccessProofModel({
    accessProofData: serviceData.latestAccessProof,
  })

  directoryService.update(
    { state: 'down' },
    latestAccessRequest,
    latestAccessProof,
  )

  expect(directoryService.state).toBe('down')
  expect(directoryService.latestAccessRequest).toBeInstanceOf(
    OutgoingAccessRequestModel,
  )
  expect(directoryService.latestAccessProof).toBeInstanceOf(AccessProofModel)
})

test('updating the model with an invalid latest access request', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(() => directoryService.update({}, 'invalid')).toThrow()
})

test('initializing the model with an invalid access proof', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(() => directoryService.update({}, null, 'invalid')).toThrow()
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
    serviceData: {
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

  await directoryService.fetch()

  expect(directoryServicesStoreFetchSpy).toHaveBeenCalledWith(directoryService)
})

describe('requesting access to a service', () => {
  it('should request access via the directory service store', async () => {
    const rootStore = new RootStore({})

    const directoryService = new DirectoryServiceModel({
      directoryServicesStore: rootStore.directoryServicesStore,
      serviceData: {},
    })

    const spy = jest
      .spyOn(rootStore.directoryServicesStore, 'requestAccess')
      .mockResolvedValue(null)

    await directoryService.requestAccess()

    expect(spy).toHaveBeenCalledWith(directoryService)
  })
})
