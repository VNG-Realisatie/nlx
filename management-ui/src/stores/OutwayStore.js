// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import OutwayModel from './models/OutwayModel'

class OutwayStore {
  _isLoading = false
  _outways = []

  constructor({ rootStore, managementApiClient }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  get outways() {
    return this._outways
  }

  fetchAll = flow(function* fetchAll() {
    try {
      this._isLoading = true
      const result = yield this._managementApiClient.managementListOutways()

      this._outways = result.outways.map(
        (outway) => new OutwayModel({ store: this, outwayData: outway }),
      )

      this._isLoading = false
    } catch (err) {
      this._isLoading = false
      throw new Error(err.message)
    }
  }).bind(this)
}

export default OutwayStore
