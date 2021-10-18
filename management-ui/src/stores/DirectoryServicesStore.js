// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import DirectoryServiceModel from '../stores/models/DirectoryServiceModel'

class DirectoryServicesStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchAll calls being in flight.
  isFetching = false

  constructor({ rootStore, directoryApiClient }) {
    makeAutoObservable(this)

    this._rootStore = rootStore
    this._directoryApiClient = directoryApiClient
  }

  fetch = flow(function* fetch(organizationSerialNumber, serviceName) {
    const serviceData =
      yield this._directoryApiClient.directoryGetOrganizationService({
        organizationSerialNumber,
        serviceName,
      })

    let directoryService = this.getService(
      organizationSerialNumber,
      serviceName,
    )

    if (!directoryService) {
      directoryService = this._updateFromServer(serviceData)
      this.services.push(directoryService)
      return directoryService
    }

    return this._updateFromServer(serviceData)
  })

  fetchAll = flow(function* fetchAll() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true
    this.error = ''

    try {
      const servicesData =
        yield this._directoryApiClient.directoryListServices()

      this.services = servicesData.services.map((serviceData) =>
        this._updateFromServer(serviceData),
      )
    } catch (e) {
      this.error = e
      console.error(e)
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

  getService = (organizationSerialNumber, serviceName) => {
    return this.services.find(
      (service) =>
        service.organization.serialNumber === organizationSerialNumber &&
        service.serviceName === serviceName,
    )
  }

  async requestAccess(organizationSerialNumber, serviceName) {
    return this._rootStore.outgoingAccessRequestStore.create(
      organizationSerialNumber,
      serviceName,
    )
  }

  _updateFromServer(serviceData) {
    const latestAccessRequest =
      this._rootStore.outgoingAccessRequestStore.updateFromServer(
        serviceData.latestAccessRequest,
      )
    const latestAccessProof = this._rootStore.accessProofStore.updateFromServer(
      serviceData.latestAccessProof,
    )

    const cachedDirectoryService = this.getService(
      serviceData.organizationSerialNumber,
      serviceData.serviceName,
    )

    if (cachedDirectoryService) {
      return cachedDirectoryService.update({
        serviceData,
        latestAccessRequest,
        latestAccessProof,
      })
    }

    return new DirectoryServiceModel({
      directoryServicesStore: this,
      serviceData,
      latestAccessProof,
      latestAccessRequest,
    })
  }

  get servicesWithAccess() {
    return this.services.filter((service) => service.hasAccess)
  }
}

export default DirectoryServicesStore
