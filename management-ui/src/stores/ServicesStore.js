// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow, action } from 'mobx'
import ServiceRepository from '../domain/service-repository'
import AccessRequestRepository from '../domain/access-request-repository'
import AccessGrantRepository from '../domain/access-grant-repository'
import { createService } from '../models/ServiceModel'

class ServicesStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchServices calls being in flight.
  isFetching = false

  constructor({
    rootStore,
    serviceRepository = ServiceRepository,
    accessRequestRepository = AccessRequestRepository,
    accessGrantRepository = AccessGrantRepository,
  }) {
    makeAutoObservable(this, {
      selectService: action.bound,
    })

    this.rootStore = rootStore
    this.serviceRepository = serviceRepository
    this.accessRequestRepository = accessRequestRepository
    this.accessGrantRepository = accessGrantRepository

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
      const services = yield this.serviceRepository.getAll()
      this.services = services.map((service) =>
        createService({ store: this, service }),
      )
    } catch (e) {
      this.error = e
      console.error(e)
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

  selectService(serviceName) {
    const serviceModel = this.services.find(
      (service) => service.name === serviceName,
    )

    if (serviceModel) {
      serviceModel.fetch()
    }

    return serviceModel
  }

  removeService = flow(function* removeService(service) {
    yield this.serviceRepository.remove(service)
    const removed = this.services.remove(service)
    if (!removed) {
      this.fetchServices()
    }
  }).bind(this)

  addService = flow(function* addService(service) {
    const newService = yield this.serviceRepository.create(service)
    const serviceModel = createService({ store: this, service: newService })

    this.services.push(serviceModel)
    return serviceModel
  }).bind(this)
}

export default ServicesStore
