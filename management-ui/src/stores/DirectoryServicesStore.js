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

    let directoryService = this.getService({ organizationName, serviceName })
    if (!directoryService) {
      directoryService = new DirectoryServiceModel({
        directoryServicesStore: this,
        serviceData,
      })
      this.services.push(directoryService)
    }

    const {
      latestAccessRequest,
      latestAccessProof,
    } = yield this.syncStoresWithServiceData(serviceData)

    directoryService.update(serviceData, latestAccessRequest, latestAccessProof)
  })

  fetchAll = flow(function* fetchAll() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true
    this.error = ''

    try {
      const servicesData = yield this.directoryRepository.getAll()

      const serviceModelsLoaded = servicesData.map(async (serviceData) => {
        const {
          latestAccessRequest,
          latestAccessProof,
        } = await this.syncStoresWithServiceData(serviceData)

        const directoryService = new DirectoryServiceModel({
          directoryServicesStore: this,
          serviceData,
        })

        directoryService.update({}, latestAccessRequest, latestAccessProof)
        return directoryService
      })

      this.services = yield Promise.all(serviceModelsLoaded)
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
    return this.rootStore.outgoingAccessRequestsStore.create({
      organizationName: directoryService.organizationName,
      serviceName: directoryService.serviceName,
    })
  }

  async syncStoresWithServiceData(serviceData) {
    const latestAccessRequest = await this.rootStore.outgoingAccessRequestsStore.updateFromServer(
      serviceData.latestAccessRequest,
    )
    const latestAccessProof = await this.rootStore.accessProofStore.updateFromServer(
      serviceData.latestAccessProof,
    )

    return {
      latestAccessRequest,
      latestAccessProof,
    }
  }
}

export default DirectoryServicesStore
