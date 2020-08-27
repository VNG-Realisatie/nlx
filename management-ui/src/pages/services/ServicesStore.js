// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, decorate, flow, observable } from 'mobx'
import ServiceRepository from '../../domain/service-repository'
import { createService } from '../../models/ServiceModel'

class ServicesStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchServices calls being in flight.
  isFetching = false

  constructor({ rootStore, domain = ServiceRepository }) {
    this.rootStore = rootStore
    this.domain = domain

    this.services = []
    this.error = ''
    this.isInitiallyFetched = false
    this.isFetching = false
  }

  fetchServices = flow(function* fetchServices() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true
    this.error = ''

    try {
      const services = yield this.domain.getAll()
      this.services = services.map((service) =>
        createService({ store: this, service }),
      )
    } catch (e) {
      this.error = e
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  })

  selectService = (serviceName) => {
    const serviceModel = this.services.find(
      (service) => service.name === serviceName,
    )

    if (serviceModel) {
      serviceModel.fetch()
    }

    return serviceModel
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
  isInitiallyFetched: observable,
  error: observable,
  fetchServices: action.bound,
  removeService: action.bound,
  addService: action.bound,
})

export const createServicesStore = (...args) => new ServicesStore(...args)

export default ServicesStore
