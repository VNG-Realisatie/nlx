// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'
import ServiceModel from '../stores/models/ServiceModel'

class ServiceStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchAll calls being in flight.
  isFetching = false

  constructor({ rootStore, managementApiClient }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this._managementApiClient = managementApiClient

    this.services = []
    this.error = ''
    this.isInitiallyFetched = false
    this.isFetching = false
  }

  fetch = flow(function* fetch({ name }) {
    const serviceData = yield this._managementApiClient.managementGetService({
      name,
    })

    let service = this.getService(name)
    if (!service) {
      service = new ServiceModel({
        servicesStore: this,
        serviceData,
      })
      this.services.push(service)
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
      const servicesData = yield this._managementApiClient.managementListServices(
        {},
      )

      this.services = servicesData.services.map(
        (serviceData) =>
          new ServiceModel({
            servicesStore: this,
            serviceData,
          }),
      )
    } catch (e) {
      this.error = e
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

  fetchStats = flow(function* fetchStats() {
    const result = yield this._managementApiClient.managementGetStatisticsOfServices()
    const stats = result.services
    if (stats.length < 1) return

    stats.forEach((statistic) => {
      const service = this.getService(statistic.name)
      if (
        service &&
        service.incomingAccessRequestCount !==
          statistic.incomingAccessRequestCount
      ) {
        service.update({
          incomingAccessRequestCount: statistic.incomingAccessRequestCount,
        })
      }
    })
  }).bind(this)

  getService = (serviceName) => {
    return this.services.find((service) => service.name === serviceName)
  }

  create = flow(function* create({
    name,
    endpointURL,
    documentationURL,
    apiSpecificationURL,
    internal,
    techSupportContact,
    publicSupportContact,
    inways,
    oneTimeCosts,
    monthlyCosts,
    requestCosts,
  }) {
    const serviceData = yield this._managementApiClient.managementCreateService(
      {
        body: {
          name,
          endpointURL,
          documentationURL,
          apiSpecificationURL,
          internal,
          techSupportContact,
          publicSupportContact,
          inways,
          oneTimeCosts: oneTimeCosts * 100,
          monthlyCosts: monthlyCosts * 100,
          requestCosts: requestCosts * 100,
        },
      },
    )
    const service = new ServiceModel({
      servicesStore: this,
      serviceData,
    })

    this.services.push(service)
    return service
  }).bind(this)

  update = flow(function* update({
    name,
    endpointURL,
    documentationURL,
    apiSpecificationURL,
    internal,
    techSupportContact,
    publicSupportContact,
    inways,
    oneTimeCosts,
    monthlyCosts,
    requestCosts,
  }) {
    if (!name) {
      throw new Error('Name required to update service')
    }

    const service = this.getService(name)

    if (!service) {
      throw new Error('Can not edit a service that does not exist')
    }

    const serviceData = yield this._managementApiClient.managementUpdateService(
      {
        name,
        body: {
          name,
          endpointURL,
          documentationURL,
          apiSpecificationURL,
          internal,
          techSupportContact,
          publicSupportContact,
          inways,
          oneTimeCosts: oneTimeCosts * 100,
          monthlyCosts: monthlyCosts * 100,
          requestCosts: requestCosts * 100,
        },
      },
    )

    service.update(serviceData)
  }).bind(this)

  removeService = flow(function* removeService(name) {
    const service = this.getService(name)
    const index = this.services.indexOf(service)

    yield this._managementApiClient.managementDeleteService({
      name,
    })

    this.services.splice(index, 1)
  }).bind(this)
}

export default ServiceStore
