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

  async requestAccess(
    organizationSerialNumber,
    serviceName,
    publicKeyFingerprint,
  ) {
    return this._rootStore.outgoingAccessRequestStore.create(
      organizationSerialNumber,
      serviceName,
      publicKeyFingerprint,
    )
  }

  _updateFromServer(serviceData) {
    const cachedDirectoryService = this.getService(
      serviceData.organization.serialNumber,
      serviceData.serviceName,
    )

    let accessStates = []

    if (serviceData.accessStates) {
      accessStates = serviceData.accessStates.map((accessState) => {
        return {
          accessRequest:
            this._rootStore.outgoingAccessRequestStore.updateFromServer(
              accessState.accessRequest,
            ),
          accessProof: this._rootStore.accessProofStore.updateFromServer(
            accessState.accessProof,
          ),
        }
      })
    }

    if (cachedDirectoryService) {
      return cachedDirectoryService.update({
        serviceData,
        accessStates,
      })
    }

    return new DirectoryServiceModel({
      directoryServicesStore: this,
      serviceData,
      accessStates,
    })
  }

  // TODO: pass fingerprint
  get servicesWithAccess() {
    return this.services.filter((service) => service.hasAccess)
  }
}

export default DirectoryServicesStore
