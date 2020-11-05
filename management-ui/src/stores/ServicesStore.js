// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'
import { serialize } from 'serializr'

import ServiceRepository from '../domain/service-repository'
import ServiceModel, { ServiceModelSchema } from '../models/ServiceModel'

class ServicesStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchAll calls being in flight.
  isFetching = false

  constructor({ rootStore, serviceRepository = ServiceRepository }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this.serviceRepository = serviceRepository

    this.services = []
    this.error = ''
    this.isInitiallyFetched = false
    this.isFetching = false
  }

  fetch = flow(function* fetch({ name }) {
    const serviceData = yield this.serviceRepository.getByName(name)

    let service = this.getService(name)
    if (!service) {
      service = new ServiceModel({
        servicesStore: this,
        serviceData,
      })
    } else {
      service.update(serviceData)
    }

    yield Promise.all([
      this.rootStore.incomingAccessRequestsStore.fetchForService(service),
      this.rootStore.accessGrantStore.fetchForService(service),
    ])
  }).bind(this)

  fetchAll = flow(function* fetchAll() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true
    this.error = ''

    try {
      const servicesData = yield this.serviceRepository.getAll()
      this.services = servicesData.map(
        (serviceData) =>
          new ServiceModel({
            servicesStore: this,
            serviceData,
          }),
      )
    } catch (e) {
      this.error = e
      console.error(e)
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

  getService = (serviceName) => {
    return this.services.find((service) => service.name === serviceName)
  }

  create = flow(function* create(formData) {
    const serviceData = yield this.serviceRepository.create(formData)
    const service = new ServiceModel({
      servicesStore: this,
      serviceData,
    })

    this.services.push(service)
    return service
  }).bind(this)

  update = flow(function* update(formData) {
    if (!formData.name) {
      throw new Error('Name required to update service')
    }

    const service = this.getService(formData.name)

    if (!service) {
      throw new Error('Can not edit a service that does not exist')
    }

    const serviceData = yield this.serviceRepository.update(
      formData.name,
      serialize(ServiceModelSchema, formData),
    )

    service.update(serviceData)
  }).bind(this)

  removeService = flow(function* removeService(service) {
    yield this.serviceRepository.remove(service)
    const removed = this.services.remove(service)
    if (!removed) {
      this.fetchAll()
    }
  }).bind(this)
}

export default ServicesStore
