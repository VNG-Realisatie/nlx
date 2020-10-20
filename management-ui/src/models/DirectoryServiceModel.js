// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, flow, makeAutoObservable } from 'mobx'
import { func, object, string } from 'prop-types'

import AccessRequestRepository from '../domain/access-request-repository'
import OutgoingAccessRequestModel from './OutgoingAccessRequestModel'

export const directoryServicePropTypes = {
  id: string.isRequired,
  organizationName: string.isRequired,
  serviceName: string.isRequired,
  state: string.isRequired,
  apiSpecificationType: string,
  latestAccessRequest: object,
  fetch: func.isRequired,
  requestAccess: func.isRequired,
}

class DirectoryServiceModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  apiSpecificationType = ''
  latestAccessRequest = null

  constructor({
    directoryServiceStore,
    service,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this, {
      update: action,
    })

    this.directoryServiceStore = directoryServiceStore
    this.accessRequestRepository = accessRequestRepository

    this.id = `${service.organizationName}/${service.serviceName}`
    this.organizationName = service.organizationName
    this.serviceName = service.serviceName
    this.state = service.state
    this.apiSpecificationType = service.apiSpecificationType

    if (
      service.latestAccessRequest &&
      !(service.latestAccessRequest instanceof OutgoingAccessRequestModel)
    ) {
      throw new Error(
        'the latestAccessRequest should be an instance of the OutgoingAccessRequestModel',
      )
    }

    this.latestAccessRequest = service.latestAccessRequest || null
  }

  update(directoryServiceData) {
    this.state = directoryServiceData.state
    this.apiSpecificationType = directoryServiceData.apiSpecificationType

    if (
      directoryServiceData.latestAccessRequest &&
      !(
        directoryServiceData.latestAccessRequest instanceof
        OutgoingAccessRequestModel
      )
    ) {
      throw new Error(
        'the latestAccessRequest should be an instance of the OutgoingAccessRequestModel',
      )
    }

    this.latestAccessRequest = directoryServiceData.latestAccessRequest || null
  }

  fetch = flow(function* fetch() {
    yield this.directoryServiceStore.fetch(this)
  }).bind(this)

  requestAccess = flow(function* requestAccess() {
    if (
      this.latestAccessRequest &&
      !this.latestAccessRequest.isCancelledOrRejected
    ) {
      return false
    }

    try {
      this.latestAccessRequest = yield this.directoryServiceStore.requestAccess(
        this,
      )
    } catch (e) {
      console.error(e)
      this.latestAccessRequest = null
    }
  }).bind(this)
}

export default DirectoryServiceModel
