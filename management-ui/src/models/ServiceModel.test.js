// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { checkPropTypes } from 'prop-types'

import deferredPromise from '../test-utils/deferred-promise'
import ServiceModel, {
  serviceModelPropTypes,
  createService,
} from './ServiceModel'

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

test('createService returns an instance', () => {
  const serviceModel = createService({ store, service })
  expect(serviceModel).toBeInstanceOf(ServiceModel)
})

test('model implements proptypes', () => {
  const errorSpy = jest.spyOn(console, 'error').mockImplementation(() => {})
  const serviceModel = new ServiceModel({ store, service })

  checkPropTypes(serviceModelPropTypes, serviceModel, 'prop', 'ServiceModel')

  expect(errorSpy).not.toHaveBeenCalled()
  errorSpy.mockRestore()
})

test('fetches data', async () => {
  const request = deferredPromise()
  store = {
    serviceRepository: {
      getByName: jest.fn(() => request),
    },
  }

  const serviceModel = new ServiceModel({ store, service })
  serviceModel.fetch()

  await request.resolve({
    ...service,
    internal: true,
  })

  expect(store.serviceRepository.getByName).toHaveBeenCalled()
  expect(serviceModel.internal).toBe(true)
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
