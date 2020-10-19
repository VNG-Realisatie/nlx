// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'
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
    makeAutoObservable(this)

    this.store = store
    this.name = service.name
    this.with(service)
  }

  fetch = flow(function* fetch() {
    const service = yield this.store.serviceRepository.getByName(this.name)
    this.with(service)
  }).bind(this)

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
  }).bind(this)

  fetchIncomingAccessRequests = flow(function* fetchIncomingAccessRequests() {
    const accessRequests = yield this.store.accessRequestRepository.listIncomingAccessRequests(
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
  }).bind(this)

  removeIncomingAccessRequest = function (removeWithId) {
    this.incomingAccessRequests = this.incomingAccessRequests.filter(
      ({ id }) => id !== removeWithId,
    )
  }

  fetchAccessGrants = flow(function* fetchAccessGrants() {
    this.accessGrants = yield this.store.accessGrantRepository.getByServiceName(
      this.name,
    )
  }).bind(this)
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

export const createService = (...args) => new ServiceModel(...args)

export default ServiceModel
