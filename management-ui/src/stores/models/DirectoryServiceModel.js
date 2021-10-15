// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'

import OutgoingAccessRequestModel from './OutgoingAccessRequestModel'
import AccessProofModel from './AccessProofModel'

function throwErrorWhenNotInstanceOf(object, model) {
  if (object && !(object instanceof model)) {
    throw new Error(`Object should be an instance of ${model}`)
  }
}

class DirectoryServiceModel {
  organization = {
    name: '',
    serialNumber: '',
  }

  serviceName = ''
  state = ''
  apiSpecificationType = ''
  documentationURL = ''
  publicSupportContact = ''
  latestAccessRequest = null
  latestAccessProof = null
  oneTimeCosts = 0
  monthlyCosts = 0
  requestCosts = 0

  constructor({
    directoryServicesStore,
    serviceData,
    latestAccessRequest,
    latestAccessProof,
  }) {
    makeAutoObservable(this)

    this.directoryServicesStore = directoryServicesStore

    this.update({ serviceData, latestAccessRequest, latestAccessProof })
  }

  update = ({
    serviceData,
    latestAccessRequest = null,
    latestAccessProof = null,
  }) => {
    if (serviceData.serviceName) {
      this.serviceName = serviceData.serviceName
    }

    if (serviceData.state) {
      this.state = serviceData.state
    }

    if (serviceData.apiSpecificationType) {
      this.apiSpecificationType = serviceData.apiSpecificationType
    }

    if (serviceData.documentationURL) {
      this.documentationURL = serviceData.documentationURL
    }

    if (serviceData.publicSupportContact) {
      this.publicSupportContact = serviceData.publicSupportContact
    }

    if (serviceData.oneTimeCosts) {
      this.oneTimeCosts = serviceData.oneTimeCosts / 100
    }

    if (serviceData.monthlyCosts) {
      this.monthlyCosts = serviceData.monthlyCosts / 100
    }

    if (serviceData.requestCosts) {
      this.requestCosts = serviceData.requestCosts / 100
    }

    if (serviceData.organization) {
      this.organization.name = serviceData.organization.name
      this.organization.serialNumber = serviceData.organization.serialNumber
    }

    throwErrorWhenNotInstanceOf(latestAccessRequest, OutgoingAccessRequestModel)
    throwErrorWhenNotInstanceOf(latestAccessProof, AccessProofModel)

    this.latestAccessRequest = latestAccessRequest
    this.latestAccessProof = latestAccessProof

    return this
  }

  fetch = async () => {
    await this.directoryServicesStore.fetch(
      this.organization.serialNumber,
      this.serviceName,
    )
  }

  requestAccess = flow(function* requestAccess() {
    try {
      this.latestAccessRequest =
        yield this.directoryServicesStore.requestAccess(this)
    } catch (error) {
      this.latestAccessRequest = null
    }
  }).bind(this)

  retryRequestAccess = flow(function* retryRequestAccess() {
    if (!this.latestAccessRequest) {
      return false
    }

    yield this.latestAccessRequest.retry()
  }).bind(this)

  get hasAccess() {
    return !!(
      this.latestAccessRequest &&
      this.latestAccessProof &&
      !this.latestAccessProof.revokedAt &&
      this.latestAccessProof.accessRequestId === this.latestAccessRequest.id
    )
  }
}

export default DirectoryServiceModel
