// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import { RootStore } from '../index'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'
import DirectoryServiceModel from './DirectoryServiceModel'
import AccessProofModel from './AccessProofModel'

test('initializing the model', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
      serialNumber: '',
    },
  })

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()

  expect(directoryService.latestAccessRequest).toBeNull()
  expect(directoryService.latestAccessProof).toBeNull()

  const latestAccessRequest = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: {},
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.CREATED,
    },
  })

  const latestAccessProof = new AccessProofModel({
    accessProofData: {
      id: 'abc',
    },
  })

  directoryService.update({
    serviceData: { state: 'down' },
    latestAccessRequest,
    latestAccessProof,
  })

  expect(directoryService.state).toBe('down')
  expect(directoryService.latestAccessRequest).toEqual(latestAccessRequest)
  expect(directoryService.latestAccessProof).toEqual(latestAccessProof)
})

test('updating the model with an invalid latest access request and access proof', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(() =>
    directoryService.update({
      latestAccessRequest: 'invalid',
    }),
  ).toThrow()

  expect(() =>
    directoryService.update({
      latestAccessProof: 'invalid',
    }),
  ).toThrow()
})

test('(re-)fetching the model', async () => {
  configure({ safeDescriptors: false })
  const rootStore = new RootStore({})

  jest.spyOn(rootStore.directoryServicesStore, 'fetch').mockResolvedValue({})

  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {
      organizationName: 'Organization',
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  await directoryService.fetch()

  expect(rootStore.directoryServicesStore.fetch).toHaveBeenCalledWith(
    directoryService,
  )
})

test('requesting access to a service', async () => {
  configure({ safeDescriptors: false })
  const rootStore = new RootStore({})

  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {},
  })

  jest
    .spyOn(rootStore.directoryServicesStore, 'requestAccess')
    .mockResolvedValue(null)

  await directoryService.requestAccess()

  expect(rootStore.directoryServicesStore.requestAccess).toHaveBeenCalledWith(
    directoryService,
  )
})

describe('access to this service', () => {
  let directoryService
  let outgoingAccessRequest
  let accessProof

  beforeEach(() => {
    outgoingAccessRequest = new OutgoingAccessRequestModel({
      outgoingAccessRequestStore: {},
      accessRequestData: {
        id: 'access-request-id',
        state: ACCESS_REQUEST_STATES.RECEIVED,
      },
    })

    accessProof = new AccessProofModel({
      accessProofData: {
        id: 'abc',
        accessRequestId: 'access-request-id',
      },
    })

    directoryService = new DirectoryServiceModel({
      directoryServicesStore: {},
      serviceData: {
        organizationName: 'Organization',
        serviceName: 'Service',
        state: 'up',
        apiSpecificationType: 'API',
      },
      latestAccessRequest: outgoingAccessRequest,
      latestAccessProof: accessProof,
    })
  })

  it('when there is an access request and valid proof', () => {
    directoryService.update({
      serviceData: {},
      latestAccessRequest: outgoingAccessRequest,
      latestAccessProof: accessProof,
    })

    expect(directoryService.hasAccess).toEqual(true)
  })

  it('when there is no proof for this service', () => {
    accessProof.update({
      accessRequestId: 'unrelated-access-request-id',
    })

    directoryService.update({
      serviceData: {},
      latestAccessRequest: outgoingAccessRequest,
      latestAccessProof: accessProof,
    })

    expect(directoryService.hasAccess).toEqual(false)
  })

  it('when there is no outgoing access request', () => {
    directoryService.update({
      serviceData: {},
      latestAccessRequest: undefined,
      latestAccessProof: accessProof,
    })

    expect(directoryService.hasAccess).toEqual(false)
  })

  it('when the access proof has been revoked', () => {
    accessProof.update({
      revokedAt: new Date(),
    })

    directoryService.update({
      serviceData: {},
      latestAccessRequest: outgoingAccessRequest,
      latestAccessProof: accessProof,
    })

    expect(directoryService.hasAccess).toEqual(false)
  })
})
