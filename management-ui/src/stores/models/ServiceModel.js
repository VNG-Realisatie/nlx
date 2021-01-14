// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

class ServiceModel {
  name = ''
  endpointURL = ''
  documentationURL = ''
  apiSpecificationURL = ''
  internal = false
  techSupportContact = ''
  publicSupportContact = ''
  inways = []
  incomingAccessRequestsCount = 0

  constructor({ servicesStore, serviceData }) {
    makeAutoObservable(this)

    this.servicesStore = servicesStore
    this.name = serviceData.name
    this.update(serviceData)
  }

  get incomingAccessRequests() {
    const allIncomingAccessRequests = this.servicesStore.rootStore.incomingAccessRequestsStore.getForService(
      this,
    )
    return allIncomingAccessRequests.filter(
      (accessRequest) => !accessRequest.isResolved,
    )
  }

  get accessGrants() {
    const allAccessGrants = this.servicesStore.rootStore.accessGrantStore.getForService(
      this,
    )
    return allAccessGrants.filter(
      (accessGrant) => accessGrant.revokedAt === null,
    )
  }

  fetch = async () => {
    await this.servicesStore.fetch(this)
  }

  update = (service) => {
    if (service.endpointURL) {
      this.endpointURL = service.endpointURL
    }

    if (service.documentationURL) {
      this.documentationURL = service.documentationURL
    }

    if (service.apiSpecificationURL) {
      this.apiSpecificationURL = service.apiSpecificationURL
    }

    if (service.internal) {
      this.internal = service.internal
    }

    if (service.techSupportContact) {
      this.techSupportContact = service.techSupportContact
    }

    if (service.publicSupportContact) {
      this.publicSupportContact = service.publicSupportContact
    }

    if (service.inways) {
      this.inways = service.inways
    }

    if (service.incomingAccessRequestsCount) {
      this.incomingAccessRequestsCount = service.incomingAccessRequestsCount
    }
  }

  hasChangedIncomingAccessRequests = async () => {
    await this.servicesStore.hasChangedIncomingAccessRequests(this)
  }
}

export default ServiceModel
