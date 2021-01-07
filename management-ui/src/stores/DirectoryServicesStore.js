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

  fetch = flow(function* fetch({ organizationName, serviceName }) {
    const serviceData = yield this._directoryApiClient.directoryGetOrganizationService(
      {
        organizationName,
        serviceName,
      },
    )

    let directoryService = this.getService({
      organizationName,
      serviceName,
    })

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
      const servicesData = yield this._directoryApiClient.directoryListServices()

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

  getService = ({ organizationName, serviceName }) => {
    return this.services.find(
      (service) =>
        service.organizationName === organizationName &&
        service.serviceName === serviceName,
    )
  }

  async requestAccess(directoryService) {
    return this._rootStore.outgoingAccessRequestStore.create({
      organizationName: directoryService.organizationName,
      serviceName: directoryService.serviceName,
    })
  }

  _updateFromServer(serviceData) {
    const latestAccessRequest = this._rootStore.outgoingAccessRequestStore.updateFromServer(
      serviceData.latestAccessRequest,
    )
    const latestAccessProof = this._rootStore.accessProofStore.updateFromServer(
      serviceData.latestAccessProof,
    )

    const cachedDirectoryService = this.getService({
      organizationName: serviceData.organizationName,
      serviceName: serviceData.serviceName,
    })

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
}

export default DirectoryServicesStore
