// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, flow, action } from 'mobx'
import { string, object, func, bool } from 'prop-types'

import { createAccessRequestInstance } from './OutgoingAccessRequestModel'

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

  isLoading: bool.isRequired,
}

class DirectoryServiceModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  apiSpecificationType = ''
  latestAccessRequest = null

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

    this.isLoading = false
  }

  fetch = flow(function* fetch() {
    this.isLoading = true

    try {
      const service = yield this.store.domain.getByName(
        this.organizationName,
        this.serviceName,
      )

      // state and latestAccessRequest are the only ones that are likely to be changed
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
    if (this.latestAccessRequest && this.latestAccessRequest.isOpen)
      return false

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
  })
}

decorate(DirectoryServiceModel, {
  state: observable,
  latestAccessRequest: observable,
  isLoading: observable,
  requestAccess: action.bound,
  fetch: action.bound,
})

export const createDirectoryService = (...args) =>
  new DirectoryServiceModel(...args)

export default DirectoryServiceModel
