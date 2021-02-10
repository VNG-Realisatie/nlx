// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { configure } from 'mobx'
import ServiceStore from '../ServiceStore'
import { RootStore } from '../index'
import { ManagementApi } from '../../api'
import IncomingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from '../../stores/models/IncomingAccessRequestModel'

import AccessGrantModel from './AccessGrantModel'
import ServiceModel from './ServiceModel'

let serviceData

beforeEach(() => {
  serviceData = {
    name: 'Service',
    endpointURL: '',
    documentationURL: '',
    apiSpecificationURL: '',
    internal: false,
    techSupportContact: '',
    publicSupportContact: '',
    inways: [],
  }
})

afterEach(() => {
  jest.restoreAllMocks()
})

test('initialize and update the service', async () => {
  jest
    .spyOn(ServiceModel.prototype, 'incomingAccessRequests', 'get')
    .mockReturnValue([])

  const serviceModel = new ServiceModel({
    servicesStore: {},
    serviceData,
  })

  serviceModel.update({ ...serviceData, internal: true })

  expect(serviceModel.internal).toBe(true)
})

test('(re-)fetching the model should call fetch on store', async () => {
  configure({ safeDescriptors: false })
  jest
    .spyOn(ServiceModel.prototype, 'incomingAccessRequests', 'get')
    .mockReturnValue([])

  const servicesStore = new ServiceStore({
    serviceRepository: {
      getByName: jest.fn().mockResolvedValue({
        name: 'Service',
      }),
    },
    rootStore: {},
  })

  const serviceModel = new ServiceModel({
    servicesStore,
    serviceData,
  })

  const servicesStoreFetchSpy = jest
    .spyOn(servicesStore, 'fetch')
    .mockResolvedValue()
  await serviceModel.fetch()

  expect(servicesStoreFetchSpy).toHaveBeenCalledWith(serviceModel)
})

test('get related incoming access requests', () => {
  const getForService = jest.fn(() => [
    new IncomingAccessRequestModel({
      accessRequestData: {
        id: '1',
        state: 'REJECTED',
      },
    }),
    new IncomingAccessRequestModel({
      accessRequestData: {
        id: '2',
        state: 'RECEIVED',
      },
    }),
  ])

  const servicesStore = new ServiceStore({
    serviceRepository: {},
    rootStore: {
      incomingAccessRequestsStore: {
        getForService,
      },
    },
  })

  const serviceModel = new ServiceModel({
    servicesStore,
    serviceData,
  })

  const incomingAccessRequests = serviceModel.incomingAccessRequests

  expect(getForService).toHaveBeenCalledWith(serviceModel)
  expect(incomingAccessRequests).toHaveLength(1)
  expect(incomingAccessRequests[0]).toEqual(
    expect.objectContaining({
      id: '2',
      state: 'RECEIVED',
    }),
  )
})

test('automatically update incomingAccessRequestCount when related incoming access requests changes', async () => {
  const managementApiClient = new ManagementApi()

  managementApiClient.managementListServices = jest.fn().mockResolvedValue({
    services: [
      {
        name: 'service-a',
        incomingAccessRequestCount: 1,
      },
    ],
  })

  managementApiClient.managementListIncomingAccessRequest = jest
    .fn()
    .mockResolvedValueOnce({
      accessRequests: [
        {
          id: '1',
          serviceName: 'service-a',
          organizationName: 'X',
          state: ACCESS_REQUEST_STATES.RECEIVED,
        },
        {
          id: '2',
          serviceName: 'service-a',
          organizationName: 'Y',
          state: ACCESS_REQUEST_STATES.RECEIVED,
        },
      ],
    })

  const rootStore = new RootStore({ managementApiClient })
  await rootStore.servicesStore.fetchAll()
  const serviceModel = rootStore.servicesStore.getService('service-a')

  expect(serviceModel.incomingAccessRequests).toHaveLength(0)
  expect(serviceModel.incomingAccessRequestCount).toBe(1)

  await rootStore.incomingAccessRequestsStore.fetchForService({
    name: 'service-a',
  })

  expect(
    managementApiClient.managementListIncomingAccessRequest,
  ).toHaveBeenCalled()
  expect(serviceModel.incomingAccessRequests).toHaveLength(2)
  expect(serviceModel.incomingAccessRequestCount).toBe(2)
})

test('get related access grants', async () => {
  jest
    .spyOn(ServiceModel.prototype, 'incomingAccessRequests', 'get')
    .mockReturnValue([])

  const getForService = jest.fn(() => [
    new AccessGrantModel({
      accessGrantData: {
        id: '1',
        createdAt: '2020-10-01',
        revokedAt: null,
      },
    }),
    new AccessGrantModel({
      accessGrantData: {
        id: '2',
        createdAt: '2020-10-01',
        revokedAt: '2020-10-02',
      },
    }),
  ])

  const servicesStore = new ServiceStore({
    serviceRepository: {},
    rootStore: {
      accessGrantStore: {
        getForService,
      },
    },
  })

  const serviceModel = new ServiceModel({
    servicesStore,
    serviceData,
  })

  const accessGrants = serviceModel.accessGrants

  expect(getForService).toHaveBeenCalledWith(serviceModel)
  expect(accessGrants).toHaveLength(1)
  expect(accessGrants[0]).toEqual(
    expect.objectContaining({
      id: '1',
      revokedAt: null,
    }),
  )
})
