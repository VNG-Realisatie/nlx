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
  ipAddress: string,
  hostname: string,
  selfAddress: string,
  services: arrayOf(
    shape({
      name: string.isRequired,
    }),
  ),
  version: string,
}

class InwayModel {
  name = ''
  ipAddress = ''
  hostname = ''
  selfAddress = ''
  services = []
  version = ''

  constructor({ store, inway }) {
    this.store = store
    this.name = inway.name
    this.with(inway)
  }

  fetch = flow(function* fetch() {
    const inway = yield this.store.inwayRepository.getByName(this.name)
    this.with(inway)
  })

  with = function (inway) {
    this.name = inway.name || ''
    this.ipAddress = inway.ipAddress || ''
    this.hostname = inway.hostname || ''
    this.selfAddress = inway.selfAddress || ''
    this.services = inway.services || []
    this.version = inway.version || ''
  }
}

createModelSchema(InwayModel, {
  name: primitive(),
  ipAddress: primitive(),
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
  ipAddress: observable,
  hostname: observable,
  selfAddress: observable,
  services: observable,
  version: observable,
  fetch: action.bound,
})

export const createInway = (...args) => new InwayModel(...args)

export default InwayModel
