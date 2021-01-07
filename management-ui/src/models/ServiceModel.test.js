// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import ServicesStore from '../stores/ServicesStore'
import IncomingAccessRequestModel from './IncomingAccessRequestModel'
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

test('initialize and update the service', async () => {
  const serviceModel = new ServiceModel({
    servicesStore: {},
    serviceData,
  })

  serviceModel.update({ ...serviceData, internal: true })

  expect(serviceModel.internal).toBe(true)
})

test('(re-)fetching the model should call fetch on store', async () => {
  const servicesStore = new ServicesStore({
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

  const servicesStore = new ServicesStore({
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

test('get related access grants', async () => {
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

  const servicesStore = new ServicesStore({
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
