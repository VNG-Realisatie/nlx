// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow, action } from 'mobx'
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
    outgoingAccessRequestsStore,
    directoryRepository = DirectoryRepository,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this, {
      selectService: action.bound,
    })

    this.outgoingAccessRequestsStore = outgoingAccessRequestsStore
    this.directoryRepository = directoryRepository
    this.accessRequestRepository = accessRequestRepository
  }

  fetchServices = flow(function* fetchServices() {
    if (this.isFetching) {
      return
    }
    this.isFetching = true
    this.error = ''

    try {
      const services = yield this.directoryRepository.getAll()
      this.services = services.map((service) =>
        mapDirectoryServiceFromApiToModel(
          this.directoryServiceStore,
          this.accessRequestRepository,
          service,
        ),
      )
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
    return this.outgoingAccessRequestsStore.create({
      organizationName: directoryService.organizationName,
      serviceName: directoryService.serviceName,
    })
  }
}

function mapDirectoryServiceFromApiToModel(
  directoryServiceStore,
  accessRequestRepository,
  service,
) {
  const latestAccessRequest = service.latestAccessRequest
    ? this.outgoingAccessRequestsStore.loadOutgoingAccessRequest(
        service.latestAccessRequest,
      )
    : null

  return new DirectoryServiceModel({
    directoryServiceStore,
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
