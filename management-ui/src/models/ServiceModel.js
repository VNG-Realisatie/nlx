// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, decorate, flow, observable } from 'mobx'
import { arrayOf, bool, func, string } from 'prop-types'
import { createModelSchema, list, primitive, serialize } from 'serializr'

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

  constructor({ store, service }) {
    this.store = store
    this.name = service.name
    this.with(service)
  }

  fetch = flow(function* fetch() {
    const service = yield this.store.domain.getByName(this.name)
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
    yield this.store.domain.update(this.name, serialize(this))
    return this
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
  fetch: action.bound,
  update: action.bound,
})

export const createService = (...args) => new ServiceModel(...args)

export default ServiceModel
