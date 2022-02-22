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
      organization: {
        name: 'Organization',
        serialNumber: '00000000000000000001',
      },
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()

  expect(directoryService.organization.name).toEqual('Organization')
  expect(directoryService.organization.serialNumber).toEqual(
    '00000000000000000001',
  )

  const accessRequest = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: {},
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.APPROVED,
      publicKeyFingerprint: 'public-key-fingerprint',
    },
  })

  const accessProof = new AccessProofModel({
    accessProofData: {
      id: 'abc',
    },
  })

  directoryService.update({
    serviceData: { state: 'down' },
    accessStates: [
      {
        accessRequest: accessRequest,
        accessProof: accessProof,
      },
    ],
  })

  expect(directoryService.state).toBe('down')
  expect(directoryService.hasAccess('public-key-fingerprint')).toEqual(true)

  expect(
    directoryService.getAccessStateFor(accessRequest.publicKeyFingerprint),
  ).toEqual({
    accessRequest: accessRequest,
    accessProof: accessProof,
  })
})

test('updating the model with an invalid latest access request and access proof', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organization: {
        name: 'Organization',
        serialNumber: '00000000000000000001',
      },
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  expect(() =>
    directoryService.update({
      accessStates: 'invalid',
    }),
  ).toThrow()

  expect(() =>
    directoryService.update({
      accessStates: [
        {
          accessRequest: 'invalid',
          accessProof: new AccessProofModel({}),
        },
      ],
    }),
  ).toThrow()

  expect(() =>
    directoryService.update({
      accessStates: [
        {
          accessRequest: new OutgoingAccessRequestModel({}),
          accessProof: 'invalid',
        },
      ],
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
      organization: {
        name: 'Organization',
        serialNumber: '00000000000000000001',
      },
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
  })

  await directoryService.fetch()

  expect(rootStore.directoryServicesStore.fetch).toHaveBeenCalledWith(
    '00000000000000000001',
    'Service',
  )
})

test('requesting access to a service', async () => {
  configure({ safeDescriptors: false })

  const rootStore = new RootStore({})

  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: rootStore.directoryServicesStore,
    serviceData: {
      organization: {
        serialNumber: '00000000000000000001',
        name: 'organization',
      },
      serviceName: 'service',
    },
  })

  jest
    .spyOn(rootStore.directoryServicesStore, 'requestAccess')
    .mockResolvedValue(null)

  jest.spyOn(rootStore.directoryServicesStore, 'fetch').mockResolvedValue(null)

  await directoryService.requestAccess('public-key-fingerprint')

  expect(rootStore.directoryServicesStore.requestAccess).toHaveBeenCalledWith(
    '00000000000000000001',
    'service',
    'public-key-fingerprint',
  )

  expect(rootStore.directoryServicesStore.fetch).toHaveBeenCalledWith(
    '00000000000000000001',
    'service',
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
        state: ACCESS_REQUEST_STATES.APPROVED,
        publicKeyFingerprint: 'public-key-fingerprint',
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
        organization: {
          name: 'Organization',
          serialNumber: '00000000000000000001',
        },
        serviceName: 'Service',
        state: 'up',
        apiSpecificationType: 'API',
      },
      accessStates: [
        {
          accessRequest: outgoingAccessRequest,
          accessProof: accessProof,
        },
      ],
    })
  })

  it('when there is an access request and valid proof', () => {
    directoryService.update({
      serviceData: {},
      accessStates: [
        {
          accessRequest: outgoingAccessRequest,
          accessProof: accessProof,
        },
      ],
    })

    expect(directoryService.hasAccess('public-key-fingerprint')).toEqual(true)
  })

  it('when there is no proof for this service', () => {
    accessProof.update({
      accessRequestId: 'unrelated-access-request-id',
    })

    directoryService.update({
      serviceData: {},
      accessStates: [
        {
          accessRequest: outgoingAccessRequest,
        },
      ],
    })

    expect(directoryService.hasAccess('public-key-fingerprint')).toEqual(false)
  })

  it('when the access proof has been revoked', () => {
    accessProof.update({
      revokedAt: new Date(),
    })

    directoryService.update({
      serviceData: {},
      accessStates: [
        {
          accessRequest: outgoingAccessRequest,
          accessProof: accessProof,
        },
      ],
    })

    expect(directoryService.hasAccess('public-key-fingerprint')).toEqual(false)
  })
})

test('retrieving the first failing access state for this service', () => {
  const accessRequest = new OutgoingAccessRequestModel({
    outgoingAccessRequestStore: {},
    accessRequestData: {
      state: ACCESS_REQUEST_STATES.FAILED,
      publicKeyFingerprint: 'public-key-fingerprint',
      errorDetails: {
        cause: 'failed to request access',
      },
    },
  })

  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organization: {
        name: 'Organization',
        serialNumber: '00000000000000000001',
      },
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
    accessStates: [
      {
        accessRequest: accessRequest,
      },
    ],
  })

  expect(directoryService.getFailingAccessStates()).toEqual([
    {
      accessRequest: accessRequest,
      accessProof: undefined,
    },
  ])
})

test('access states with access', () => {
  const directoryService = new DirectoryServiceModel({
    directoryServicesStore: {},
    serviceData: {
      organization: {
        name: 'Organization',
        serialNumber: '00000000000000000001',
      },
      serviceName: 'Service',
      state: 'up',
      apiSpecificationType: 'API',
    },
    accessStates: [
      {
        accessRequest: new OutgoingAccessRequestModel({
          outgoingAccessRequestStore: {},
          accessRequestData: {
            state: ACCESS_REQUEST_STATES.APPROVED,
            publicKeyFingerprint: 'public-key-fingerprint',
          },
        }),
        accessProof: new AccessProofModel({
          accessProofData: {
            id: '42',
          },
        }),
      },
    ],
  })

  expect(directoryService.accessStatesWithAccess).toHaveLength(1)
})
