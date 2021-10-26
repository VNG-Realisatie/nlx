// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import OutwayModel from './models/OutwayModel'

class OutwayStore {
  _isLoading = false
  _outways = []

  // This is set to true after the first call has been made. Regardless of success.
  _isInitiallyFetched = false

  // This is internal state to prevent concurrent fetchInways calls being in flight.
  _isFetching = false

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

  fetch = flow(function* fetch({ name }) {
    const outwayData = yield this._managementApiClient.managementGetOutway({
      name,
    })
    let outway = this.getByName({ name })

    if (!outway) {
      outway = this._updateFromServer(outwayData)
      this._outways.push(outway)
      return outway
    }

    return this._updateFromServer(outwayData)
  }).bind(this)

  getByName = (name) => {
    return this._outways.find((outway) => outway.name === name)
  }

  _updateFromServer(outwayData) {
    const cachedOutway = this.getByName(outwayData.name)

    if (cachedOutway) {
      cachedOutway.update(outwayData)
      return cachedOutway
    }

    return new OutwayModel({
      store: this,
      outwayData: outwayData,
    })
  }
}

export default OutwayStore
