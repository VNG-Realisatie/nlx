// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'

class OutwayModel {
  _name = ''
  _ipAddress = ''
  _publicKeyPEM = ''
  _version = ''

  get name() {
    return this._name
  }

  get ipAddress() {
    return this._ipAddress
  }

  get publicKeyPEM() {
    return this._publicKeyPEM
  }

  get version() {
    return this._version
  }

  constructor({ store, outwayData }) {
    makeAutoObservable(this)

    this.outwayStore = store

    this.update(outwayData)
  }

  fetch = flow(function* fetch() {
    const outway = yield this.outwayStore.fetch({ name: this.name })
    this.with(outway)
  }).bind(this)

  update = function (outway) {
    this._name = outway.name
    this._ipAddress = outway.ipAddress
    this._publicKeyPEM = outway.publicKeyPEM
    this._version = outway.version
  }
}

export default OutwayModel
