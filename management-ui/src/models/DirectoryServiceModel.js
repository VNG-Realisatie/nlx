// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
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
    makeAutoObservable(this)

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

  fetch = flow(function* fetch() {
    const service = yield this.directoryServiceStore.directoryRepository.getByName(
      this.organizationName,
      this.serviceName,
    )

    this.state = service.state
    this.latestAccessRequest = service.latestAccessRequest
      ? new OutgoingAccessRequestModel({
          accessRequestData: service.latestAccessRequest,
        })
      : null
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

export const createDirectoryService = (...args) =>
  new DirectoryServiceModel(...args)

export default DirectoryServiceModel
