// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
import ServicesStore from '../stores/ServicesStore'
import ServiceModel, { serviceModelPropTypes } from './ServiceModel'

let store
let service

beforeEach(() => {
  store = {}
  service = {
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
  const serviceModel = new ServiceModel({ store, service })

  checkPropTypes(serviceModelPropTypes, serviceModel, 'prop', 'ServiceModel')

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('(re-)fetching the model', async () => {
  const serviceStore = new ServicesStore({
    serviceRepository: {
      getByName: jest.fn().mockResolvedValue({
        name: 'Service',
      }),
    },
  })

  const serviceModel = new ServiceModel({
    store: serviceStore,
    service: {
      name: 'Service',
      internal: true,
    },
  })

  expect(serviceModel.name).toBe('Service')
  expect(serviceModel.internal).toBe(true)

  const servicesStoreFetchSpy = jest
    .spyOn(serviceStore, 'fetch')
    .mockResolvedValue()
  await serviceModel.fetch()

  expect(servicesStoreFetchSpy).toHaveBeenCalledWith(serviceModel)
})

test('updates service', async () => {
  store = {
    serviceRepository: {
      update: jest.fn(async (name, service) => ({ ...service })),
    },
  }
  const serviceModel = new ServiceModel({ store, service })

  await serviceModel.update({ ...service, internal: true })
  await expect(store.serviceRepository.update).toHaveBeenCalledWith(
    service.name,
    expect.objectContaining({ name: service.name, internal: true }),
  )

  expect(serviceModel.internal).toBe(true)
})

test('fetching access grants', async () => {
  const request = deferredPromise()
  store = {
    accessGrantRepository: {
      getByServiceName: jest.fn(() => request),
    },
  }

  const serviceModel = new ServiceModel({ store, service })
  serviceModel.fetchAccessGrants()

  await request.resolve([{ id: 'somegrant' }])

  expect(store.accessGrantRepository.getByServiceName).toHaveBeenCalled()
  expect(serviceModel.accessGrants).toEqual([{ id: 'somegrant' }])
})
