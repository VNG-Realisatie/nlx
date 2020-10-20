// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, flow, makeAutoObservable } from 'mobx'
import DirectoryRepository from '../domain/directory-repository'
import DirectoryServiceModel from '../models/DirectoryServiceModel'
import AccessRequestRepository from '../domain/access-request-repository'

class DirectoryStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchServices calls being in flight.
  isFetching = false

  constructor({
    rootStore,
    directoryRepository = DirectoryRepository,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this, {
      selectService: action.bound,
    })

    this.rootStore = rootStore
    this.directoryRepository = directoryRepository
    this.accessRequestRepository = accessRequestRepository
  }

  fetch = flow(function* fetchService(directoryServiceModel) {
    const response = yield this.directoryRepository.getByName(
      directoryServiceModel.organizationName,
      directoryServiceModel.serviceName,
    )

    let outgoingAccessRequestModel = null
    if (response.latestAccessRequest) {
      outgoingAccessRequestModel = yield this.rootStore.outgoingAccessRequestsStore.loadOutgoingAccessRequest(
        response.latestAccessRequest,
      )
    }

    directoryServiceModel.update({
      ...response,
      latestAccessRequest: outgoingAccessRequestModel,
    })
  })

  fetchServices = flow(function* fetchServices() {
    if (this.isFetching) {
      return
    }
    this.isFetching = true
    this.error = ''

    try {
      const services = yield this.directoryRepository.getAll()
      const loadServiceModels = services.map(
        async (service) =>
          await mapDirectoryServiceFromApiToModel(
            this.rootStore,
            this.accessRequestRepository,
            service,
          ),
      )

      const serviceModels = yield Promise.all(loadServiceModels)

      this.services = serviceModels
    } catch (e) {
      this.error = e
      console.error(e)
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

  selectService({ organizationName, serviceName }) {
    const directoryServiceModel = this.services.find(
      (service) =>
        service.organizationName === organizationName &&
        service.serviceName === serviceName,
    )
    if (directoryServiceModel) {
      directoryServiceModel.fetch()
    }
    return directoryServiceModel
  }

  async requestAccess(directoryService) {
    return this.rootStore.outgoingAccessRequestsStore.create({
      organizationName: directoryService.organizationName,
      serviceName: directoryService.serviceName,
    })
  }
}

async function mapDirectoryServiceFromApiToModel(
  rootStore,
  accessRequestRepository,
  service,
) {
  const latestAccessRequest = service.latestAccessRequest
    ? await rootStore.outgoingAccessRequestsStore.loadOutgoingAccessRequest(
        service.latestAccessRequest,
      )
    : null

  return new DirectoryServiceModel({
    directoryServiceStore: rootStore.directoryStore,
    accessRequestRepository,
    service: {
      id: service.id,
      organizationName: service.organizationName,
      serviceName: service.serviceName,
      state: service.state,
      apiSpecificationType: service.apiSpecificationType,
      latestAccessRequest: latestAccessRequest,
    },
  })
}

export default DirectoryStore
