// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import { bool, func, object, string } from 'prop-types'

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
  isOpen: bool,
}

class DirectoryServiceModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  apiSpecificationType = ''
  latestAccessRequest = null

  constructor({
    store,
    service,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this)

    this.store = store
    this.accessRequestRepository = accessRequestRepository

    this.id = `${service.organizationName}/${service.serviceName}`
    this.organizationName = service.organizationName
    this.serviceName = service.serviceName
    this.state = service.state
    this.apiSpecificationType = service.apiSpecificationType
    this.latestAccessRequest = service.latestAccessRequest
      ? new OutgoingAccessRequestModel({
          accessRequestData: service.latestAccessRequest,
          accessRequestRepository: accessRequestRepository,
        })
      : null
  }

  fetch = flow(function* fetch() {
    const service = yield this.store.directoryRepository.getByName(
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

    this.latestAccessRequest = yield new OutgoingAccessRequestModel({
      accessRequestData: {
        organizationName: this.organizationName,
        serviceName: this.serviceName,
      },
      accessRequestRepository: this.accessRequestRepository,
    })

    try {
      yield this.latestAccessRequest.send()
    } catch (e) {
      console.error(e)
      this.latestAccessRequest = null
    }
  }).bind(this)
}

export const createDirectoryService = (...args) =>
  new DirectoryServiceModel(...args)

export default DirectoryServiceModel
