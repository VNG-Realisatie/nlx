// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, flow, observable } from 'mobx'

import { createModelSchema, primitive } from 'serializr'
import { string } from 'prop-types'

export const inwayModelPropTypes = {
  name: string.isRequired,
}

class InwayModel {
  name = ''

  constructor({ store, inway }) {
    this.store = store
    this.name = inway.name
    this.with(inway)
  }

  fetch = flow(function* fetch() {
    const inway = yield this.store.domain.getByName(this.name)
    this.with(inway)
  })

  with = function (inway) {
    this.name = inway.name || ''
  }
}

createModelSchema(InwayModel, {
  name: primitive(),
})

decorate(InwayModel, {
  name: observable,
})

export const createInway = (...args) => new InwayModel(...args)

export default InwayModel
