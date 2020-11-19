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
    if (serviceData.organizationName) {
      this.organizationName = serviceData.organizationName
    }

    if (serviceData.serviceName) {
      this.serviceName = serviceData.serviceName
    }

    if (serviceData.state) {
      this.state = serviceData.state
    }

    if (serviceData.apiSpecificationType) {
      this.apiSpecificationType = serviceData.apiSpecificationType
    }

    throwErrorWhenNotInstanceOf(latestAccessRequest, OutgoingAccessRequestModel)
    throwErrorWhenNotInstanceOf(latestAccessProof, AccessProofModel)

    this.latestAccessRequest = latestAccessRequest
    this.latestAccessProof = latestAccessProof

    return this
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
