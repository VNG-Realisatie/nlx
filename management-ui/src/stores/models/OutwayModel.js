// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

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

  constructor({ outwayData }) {
    makeAutoObservable(this)
    this.update(outwayData)
  }

  update = function (outway) {
    this._name = outway.name
    this._ipAddress = outway.ipAddress
    this._publicKeyPEM = outway.publicKeyPEM
    this._version = outway.version
  }
}

export default OutwayModel
