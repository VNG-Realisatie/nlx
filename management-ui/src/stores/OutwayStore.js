// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import OutwayModel from './models/OutwayModel'

class OutwayStore {
  _isLoading = false
  _outways = []

  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false

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

  fetchAll = flow(function* fetchInways() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true

    try {
      const result = yield this._managementApiClient.managementListOutways()
      this._outways = result.outways.map(
        (outway) => new OutwayModel({ store: this, outwayData: outway }),
      )
    } catch (err) {
      throw new Error(err.message)
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)
}

export default OutwayStore
