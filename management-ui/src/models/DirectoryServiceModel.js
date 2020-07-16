// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, flow } from 'mobx'

import { createAccessRequestInstance } from './OutgoingAccessRequestModel'

class DirectoryServiceModel {
  id = ''
  organizationName = ''
  serviceName = ''
  status = ''
  apiSpecificationType = ''
  latestAccessRequest = null

  constructor({ store, service }) {
    this.store = store

    this.id = `${service.organizationName}/${service.serviceName}`
    this.organizationName = service.organizationName
    this.serviceName = service.serviceName
    this.status = service.status
    this.apiSpecificationType = service.apiSpecificationType
    this.latestAccessRequest = service.latestAccessRequest
      ? createAccessRequestInstance(service.latestAccessRequest)
      : null
  }

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
  latestAccessRequest: observable,
})

export default DirectoryServiceModel
