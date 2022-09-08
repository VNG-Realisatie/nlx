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

  constructor({ rootStore, directoryApiClient, managementApiClient }) {
    makeAutoObservable(this)

    this._rootStore = rootStore
    this._directoryApiClient = directoryApiClient
    this._managementApiClient = managementApiClient
  }

  fetch = flow(function* fetch(organizationSerialNumber, serviceName) {
    try {
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
    } catch (error) {
      if (error.status === 404) {
        const service = this.getService(organizationSerialNumber, serviceName)
        this.services.remove(service)
      }
    }
  })
  //TODO: add tests
  syncOutgoingAccessRequests = flow(function* fetch(
    organizationSerialNumber,
    serviceName,
  ) {
    yield this._managementApiClient.managementSynchronizeOutgoingAccessRequests(
      {
        organizationSerialNumber,
        serviceName,
      },
    )
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

  async requestAccess(organizationSerialNumber, serviceName, publicKeyPEM) {
    return this._managementApiClient.managementSendAccessRequest({
      organizationSerialNumber: organizationSerialNumber,
      serviceName: serviceName,
      publicKeyPEM: publicKeyPEM,
    })
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

  get servicesWithAccess() {
    return this._rootStore.accessProofStore.accessProofs
      .map((accessProof) =>
        this._rootStore.directoryServicesStore.getService(
          accessProof.organization.serialNumber,
          accessProof.serviceName,
        ),
      )
      .filter((service) => !!service)
      .filter((directoryService, i, arr) => arr.indexOf(directoryService) === i)
  }
}

export default DirectoryServicesStore
