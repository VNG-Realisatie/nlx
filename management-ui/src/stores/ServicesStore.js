// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'
import ServiceRepository from '../domain/service-repository'
import ServiceModel from '../models/ServiceModel'

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

  fetch = flow(function* fetch(serviceModel) {
    const response = yield this.serviceRepository.getByName(serviceModel.name)
    serviceModel.with(response)
    yield serviceModel
  }).bind(this)

  fetchAll = flow(function* fetchAll() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true
    this.error = ''

    try {
      const services = yield this.serviceRepository.getAll()
      this.services = services.map(
        (service) => new ServiceModel({ store: this, service }),
      )
    } catch (e) {
      this.error = e
      console.error(e)
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

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
    yield this.serviceRepository.remove(service)
    const removed = this.services.remove(service)
    if (!removed) {
      this.fetchAll()
    }
  }).bind(this)

  create = flow(function* create(service) {
    const response = yield this.serviceRepository.create(service)
    const serviceModel = new ServiceModel({ store: this, service: response })

    this.services.push(serviceModel)
    return serviceModel
  }).bind(this)
}

export default ServicesStore
