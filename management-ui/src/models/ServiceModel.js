// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, decorate, flow, observable } from 'mobx'
import { arrayOf, bool, func, string } from 'prop-types'
import { createModelSchema, list, primitive, serialize } from 'serializr'
import {
  ACCESS_REQUEST_STATES,
  createIncomingAccessRequest,
} from './IncomingAccessRequestModel'

export const serviceModelPropTypes = {
  name: string.isRequired,
  endpointURL: string.isRequired,
  documentationURL: string.isRequired,
  apiSpecificationURL: string.isRequired,
  internal: bool.isRequired,
  techSupportContact: string.isRequired,
  publicSupportContact: string.isRequired,
  inways: arrayOf(string),
  fetch: func.isRequired,
  update: func.isRequired,
}

// TODO test
class ServiceModel {
  name = ''
  endpointURL = ''
  documentationURL = ''
  apiSpecificationURL = ''
  internal = false
  techSupportContact = ''
  publicSupportContact = ''
  inways = []
  incomingAccessRequests = []
  accessGrants = []

  constructor({ store, service }) {
    this.store = store
    this.name = service.name
    this.with(service)
  }

  fetch = flow(function* fetch() {
    const service = yield this.store.serviceRepository.getByName(this.name)
    this.with(service)
  })

  with = function (service) {
    this.endpointURL = service.endpointURL || ''
    this.documentationURL = service.documentationURL || ''
    this.apiSpecificationURL = service.apiSpecificationURL || ''
    this.internal = service.internal || false
    this.techSupportContact = service.techSupportContact || ''
    this.publicSupportContact = service.publicSupportContact || ''
    this.inways = service.inways || []
  }

  update = flow(function* update(values) {
    this.with(values)
    yield this.store.serviceRepository.update(this.name, serialize(this))
    return this
  })

  fetchIncomingAccessRequests = flow(function* fetchIncomingAccessRequests() {
    const accessRequests = yield this.store.accessRequestRepository.getIncomingAccessRequests(
      this.name,
    )

    this.incomingAccessRequests = accessRequests
      .filter(
        (accessRequest) =>
          accessRequest.state === ACCESS_REQUEST_STATES.RECEIVED,
      )
      .map((accessRequest) =>
        createIncomingAccessRequest({
          store: this,
          accessRequestData: accessRequest,
        }),
      )
  })

  removeIncomingAccessRequest = function (removeWithId) {
    this.incomingAccessRequests = this.incomingAccessRequests.filter(
      ({ id }) => id !== removeWithId,
    )
  }

  fetchAccessGrants = flow(function* fetchAccessGrants() {
    this.accessGrants = yield this.store.accessGrantRepository.getByService(
      this.name,
    )
  })
}

createModelSchema(ServiceModel, {
  name: primitive(),
  endpointURL: primitive(),
  documentationURL: primitive(),
  apiSpecificationURL: primitive(),
  internal: primitive(),
  techSupportContact: primitive(),
  publicSupportContact: primitive(),
  inways: list(primitive()),
})

decorate(ServiceModel, {
  name: observable,
  endpointURL: observable,
  documentationURL: observable,
  apiSpecificationURL: observable,
  internal: observable,
  techSupportContact: observable,
  publicSupportContact: observable,
  inways: observable,
  incomingAccessRequests: observable,
  accessGrants: observable,
  fetch: action.bound,
  update: action.bound,
  fetchIncomingAccessRequests: action.bound,
  fetchAccessGrants: action.bound,
  tempRemoveFromList: action,
})

export const createService = (...args) => new ServiceModel(...args)

export default ServiceModel
