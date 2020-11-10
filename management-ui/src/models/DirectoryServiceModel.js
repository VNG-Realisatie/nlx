// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import { func, object, string } from 'prop-types'

import OutgoingAccessRequestModel from './OutgoingAccessRequestModel'
import AccessProofModel from './AccessProofModel'

export const directoryServicePropTypes = {
  organizationName: string.isRequired,
  serviceName: string.isRequired,
  state: string.isRequired,
  apiSpecificationType: string,
  latestAccessRequest: object,
  latestAccessProof: object,
  fetch: func.isRequired,
  requestAccess: func.isRequired,
  retryRequestAccess: func.isRequired,
}

function throwErrorWhenNotInstanceOf(object, model) {
  if (object && !(object instanceof model)) {
    throw new Error(
      `Object should be an instance of OutgoingAccessRequestModel`,
    )
  }
}

class DirectoryServiceModel {
  organizationName = ''
  serviceName = ''
  state = ''
  apiSpecificationType = ''
  latestAccessRequest = null
  latestAccessProof = null

  constructor({ directoryServicesStore, serviceData }) {
    makeAutoObservable(this)

    this.directoryServicesStore = directoryServicesStore

    this.update(serviceData)
  }

  update = (
    directoryServiceData,
    latestAccessRequest = null,
    latestAccessProof = null,
  ) => {
    if (directoryServiceData.organizationName) {
      this.organizationName = directoryServiceData.organizationName
    }

    if (directoryServiceData.serviceName) {
      this.serviceName = directoryServiceData.serviceName
    }

    if (directoryServiceData.state) {
      this.state = directoryServiceData.state
    }

    if (directoryServiceData.apiSpecificationType) {
      this.apiSpecificationType = directoryServiceData.apiSpecificationType
    }

    throwErrorWhenNotInstanceOf(latestAccessRequest, OutgoingAccessRequestModel)
    throwErrorWhenNotInstanceOf(latestAccessProof, AccessProofModel)

    this.latestAccessRequest = latestAccessRequest
    this.latestAccessProof = latestAccessProof
  }

  fetch = async () => {
    await this.directoryServicesStore.fetch(this)
  }

  requestAccess = flow(function* requestAccess() {
    try {
      this.latestAccessRequest = yield this.directoryServicesStore.requestAccess(
        this,
      )
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
}

export default DirectoryServiceModel
