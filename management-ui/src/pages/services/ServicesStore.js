// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, decorate, flow, observable } from 'mobx'
import ServiceRepository from '../../domain/service-repository'
import { createService } from '../../models/ServiceModel'

class ServicesStore {
  services = []
  isReady = false
  error = ''

  constructor({ rootStore, domain = ServiceRepository }) {
    this.rootStore = rootStore
    this.domain = domain

    this.isReady = false
    this.error = ''
  }

  fetchServices = flow(function* fetchServices() {
    this.isReady = false
    this.error = ''

    try {
      const services = yield this.domain.getAll()
      this.services = services.map((service) =>
        createService({ store: this, service }),
      )
    } catch (e) {
      this.error = e
    } finally {
      this.isReady = true
    }
  })

  selectService = (serviceName) => {
    return this.services.find((service) => service.name === serviceName)
  }

  removeService = flow(function* removeService(service) {
    yield this.domain.remove(service)
    const removed = this.services.remove(service)
    if (!removed) {
      this.fetchServices()
    }
  })

  addService = flow(function* addService(service) {
    const newService = yield this.domain.create(service)
    const serviceModel = createService({ store: this, service: newService })
    this.services.push(serviceModel)
    return serviceModel
  })
}

decorate(ServicesStore, {
  services: observable,
  isLoading: observable,
  error: observable,
  fetchServices: action.bound,
  removeService: action.bound,
  addService: action.bound,
})

export const createServicesStore = (...args) => new ServicesStore(...args)

export default ServicesStore
