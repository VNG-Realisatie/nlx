// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import ServicesStore from '../stores/ServicesStore'
import IncomingAccessRequestModel from './IncomingAccessRequestModel'
import ServiceModel, { serviceModelPropTypes } from './ServiceModel'

let store
let serviceData

beforeEach(() => {
  store = {}
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

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const serviceModel = new ServiceModel({ store, serviceData })

  checkPropTypes(serviceModelPropTypes, serviceModel, 'prop', 'ServiceModel')

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
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

// test.only('fetching access grants', async () => {
//   const request = deferredPromise()
//   store = {
//     accessGrantRepository: {
//       getByServiceName: jest.fn(() => request),
//     },
//   }

//   const serviceModel = new ServiceModel({ store, serviceData })
//   serviceModel.fetchAccessGrants()

//   await request.resolve([{ id: 'somegrant' }])

//   expect(store.accessGrantRepository.getByServiceName).toHaveBeenCalled()
//   expect(serviceModel.accessGrants).toEqual([{ id: 'somegrant' }])
// })
