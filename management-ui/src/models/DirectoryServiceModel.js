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

  constructor({ directoryServicesStore, service }) {
    makeAutoObservable(this)

    this.directoryServicesStore = directoryServicesStore

    this.update(service)
  }

  update = (directoryServiceData) => {
    this.organizationName = directoryServiceData.organizationName
    this.serviceName = directoryServiceData.serviceName
    this.state = directoryServiceData.state
    this.apiSpecificationType = directoryServiceData.apiSpecificationType

    throwErrorWhenNotInstanceOf(latestAccessRequest, OutgoingAccessRequestModel)
    throwErrorWhenNotInstanceOf(latestAccessProof, AccessProofModel)

    this.latestAccessRequest = directoryServiceData.latestAccessRequest || null
    this.latestAccessProof = directoryServiceData.latestAccessProof || null
  }

  fetch = flow(function* fetch() {
    yield this.directoryServicesStore.fetch(this)
  }).bind(this)

  requestAccess = flow(function* requestAccess() {
    if (
      this.latestAccessRequest &&
      !this.latestAccessRequest.isCancelledOrRejected
    ) {
      return false
    }

    try {
      this.latestAccessRequest = yield this.directoryServicesStore.requestAccess(
        this,
      )
    } catch (e) {
      console.error(e)
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
