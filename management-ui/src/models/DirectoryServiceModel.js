// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, flow, action } from 'mobx'

import { createAccessRequestInstance } from './OutgoingAccessRequestModel'

class DirectoryServiceModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  apiSpecificationType = ''
  latestAccessRequest = null

  isLoading = false

  constructor({ store, service }) {
    this.store = store

    this.id = `${service.organizationName}/${service.serviceName}`
    this.organizationName = service.organizationName
    this.serviceName = service.serviceName
    this.state = service.state
    this.apiSpecificationType = service.apiSpecificationType
    this.latestAccessRequest = service.latestAccessRequest
      ? createAccessRequestInstance(service.latestAccessRequest)
      : null
  }

  fetch = flow(function* fetch() {
    this.isLoading = true

    try {
      const service = yield this.store.domain.getByName(
        this.organizationName,
        this.serviceName,
      )

      this.state = service.state
      this.latestAccessRequest = service.latestAccessRequest
        ? createAccessRequestInstance(service.latestAccessRequest)
        : null
    } catch (e) {
    } finally {
      this.isLoading = false
    }
  })

  requestAccess = flow(function* requestAccess() {
    if (
      !this.latestAccessRequest ||
      this.latestAccessRequest.hasUnsuccessfulEndstate
    ) {
      this.latestAccessRequest = yield createAccessRequestInstance({
        organizationName: this.organizationName,
        serviceName: this.serviceName,
      })

      try {
        yield this.latestAccessRequest.send()
      } catch (e) {
        console.error(e)
        this.latestAccessRequest = null
      }
    }
  })
}

decorate(DirectoryServiceModel, {
  state: observable,
  latestAccessRequest: observable,
  isLoading: observable,
  requestAccess: action.bound,
  fetch: action.bound,
})

export default DirectoryServiceModel
