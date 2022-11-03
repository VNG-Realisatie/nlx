// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, reaction } from 'mobx'

class ServiceModel {
  name = ''
  endpointUrl = ''
  documentationUrl = ''
  apiSpecificationUrl = ''
  internal = false
  techSupportContact = ''
  publicSupportContact = ''
  inways = []
  incomingAccessRequestCount = 0
  oneTimeCosts = 0
  monthlyCosts = 0
  requestCosts = 0

  constructor({ servicesStore, serviceData }) {
    makeAutoObservable(this)

    this.servicesStore = servicesStore
    this.name = serviceData.name
    this.update(serviceData)

    reaction(
      () => this.incomingAccessRequests.length,
      (incomingAccessRequestCount) => {
        this.update({ incomingAccessRequestCount })
      },
    )
  }

  get incomingAccessRequests() {
    const allIncomingAccessRequests =
      this.servicesStore.rootStore.incomingAccessRequestsStore.getForService(
        this,
      )
    return allIncomingAccessRequests.filter(
      (accessRequest) => !accessRequest.isResolved,
    )
  }

  get accessGrants() {
    const allAccessGrants =
      this.servicesStore.rootStore.accessGrantStore.getForService(this)
    return allAccessGrants.filter(
      (accessGrant) =>
        accessGrant.revokedAt === null && accessGrant.terminatedAt === null,
    )
  }

  fetch = async () => {
    await this.servicesStore.fetch(this)
  }

  update = (service) => {
    if (service.endpointUrl) {
      this.endpointUrl = service.endpointUrl
    }

    if (service.documentationUrl) {
      this.documentationUrl = service.documentationUrl
    }

    if (service.apiSpecificationUrl) {
      this.apiSpecificationUrl = service.apiSpecificationUrl
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

    if (service.incomingAccessRequestCount !== undefined) {
      this.incomingAccessRequestCount = service.incomingAccessRequestCount
    }

    if ('oneTimeCosts' in service) {
      this.oneTimeCosts = service.oneTimeCosts / 100
    }

    if ('monthlyCosts' in service) {
      this.monthlyCosts = service.monthlyCosts / 100
    }

    if ('requestCosts' in service) {
      this.requestCosts = service.requestCosts / 100
    }
  }
}

export default ServiceModel
