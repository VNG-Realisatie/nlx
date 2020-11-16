// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import DirectoryRepository from '../domain/directory-repository'
import DirectoryServiceModel from '../models/DirectoryServiceModel'

class DirectoryServicesStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchAll calls being in flight.
  isFetching = false

  constructor({ rootStore, directoryRepository = DirectoryRepository }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this.directoryRepository = directoryRepository
  }

  fetch = flow(function* fetch({ organizationName, serviceName }) {
    const serviceData = yield this.directoryRepository.getByName(
      organizationName,
      serviceName,
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
      const servicesData = yield this.directoryRepository.getAll()

      this.services = servicesData.map((serviceData) =>
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
    return this.rootStore.outgoingAccessRequestStore.create({
      organizationName: directoryService.organizationName,
      serviceName: directoryService.serviceName,
    })
  }

  _updateFromServer(serviceData) {
    const latestAccessRequest = this.rootStore.outgoingAccessRequestStore.updateFromServer(
      serviceData.latestAccessRequest,
    )
    const latestAccessProof = this.rootStore.accessProofStore.updateFromServer(
      serviceData.latestAccessProof,
    )

    const cachedDirectoryService = this.getService({
      organizationName: serviceData.organizationName,
      serviceName: serviceData.serviceName,
    })

    if (cachedDirectoryService) {
      return cachedDirectoryService.update(
        serviceData,
        latestAccessRequest,
        latestAccessProof,
      )
    }

    return new DirectoryServiceModel({
      directoryServicesStore: this,
      serviceData: Object.assign({}, serviceData, {
        latestAccessProof: latestAccessProof,
        latestAccessRequest: latestAccessRequest,
      }),
    })
  }
}

export default DirectoryServicesStore
