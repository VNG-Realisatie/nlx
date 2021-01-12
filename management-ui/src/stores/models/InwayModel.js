// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'
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
    makeAutoObservable(this)

    this.inwayStore = store
    this.name = inway.name
    this.with(inway)
  }

  fetch = flow(function* fetch() {
    const inway = yield this.inwayStore.fetch({ name: this.name })
    this.with(inway)
  }).bind(this)

  with = function (inway) {
    this.name = inway.name || ''
    this.ipAddress = inway.ipAddress || ''
    this.hostname = inway.hostname || ''
    this.selfAddress = inway.selfAddress || ''
    this.services = inway.services || []
    this.version = inway.version || ''
  }
}

export default InwayModel
