// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

class OutwayModel {
  _name = ''
  _version = ''

  get name() {
    return this._name
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
    this._version = outway.version
  }
}

export default OutwayModel
