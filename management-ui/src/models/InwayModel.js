// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, decorate, flow, observable } from 'mobx'

import {
  createModelSchema,
  createSimpleSchema,
  list,
  object,
  primitive,
} from 'serializr'
import { arrayOf, shape, string } from 'prop-types'

export const inwayModelPropTypes = {
  name: string.isRequired,
  hostname: string.isRequired,
  selfAddress: string.isRequired,
  services: arrayOf(
    shape({
      name: string.isRequired,
    }),
  ).isRequired,
  version: string.isRequired,
}

class InwayModel {
  name = ''

  constructor({ store, inway }) {
    this.store = store

    this.name = inway.name
    this.hostname = inway.hostname
    this.selfAddress = inway.selfAddress
    this.services = inway.services
    this.version = inway.version

    this.with(inway)
  }

  fetch = flow(function* fetch() {
    const inway = yield this.store.domain.getByName(this.name)
    this.with(inway)
  })

  with = function (inway) {
    this.name = inway.name || ''
    this.hostname = inway.hostname || ''
    this.selfAddress = inway.selfAddress || ''
    this.services = inway.services || []
    this.version = inway.version || ''
  }
}

createModelSchema(InwayModel, {
  name: primitive(),
  hostname: primitive(),
  selfAddress: primitive(),
  services: list(
    object(
      createSimpleSchema({
        name: primitive(),
      }),
    ),
  ),
  version: primitive(),
})

decorate(InwayModel, {
  name: observable,
  hostname: observable,
  selfAddress: observable,
  services: observable,
  version: observable,
  fetch: action.bound,
})

export const createInway = (...args) => new InwayModel(...args)

export default InwayModel
