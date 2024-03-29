// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import OutwayModel from './models/OutwayModel'

class OutwayStore {
  _isLoading = false
  _outways = []
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

  get publicKeyFingerprints() {
    return this._outways
      .map((outway) => outway.publicKeyFingerprint)
      .filter(
        (publicKeyFingerprint, i, array) =>
          array.indexOf(publicKeyFingerprint) === i,
      )
  }

  fetchAll = flow(function* fetchAll() {
    if (this._isFetching) {
      return
    }

    this._isFetching = true

    try {
      const result =
        yield this._managementApiClient.managementServiceListOutways()
      this._outways = result.outways.map(
        (outway) => new OutwayModel({ store: this, outwayData: outway }),
      )
      this._isFetching = false
    } catch (err) {
      this._isFetching = false
      throw new Error(err.message)
    }
  }).bind(this)

  fetch = flow(function* fetch({ name }) {
    const outwayData =
      yield this._managementApiClient.managementServiceGetOutway({
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

  getByPublicKeyFingerprint = (publicKeyFingerprint) => {
    return this._outways.filter(
      (outway) => outway.publicKeyFingerprint === publicKeyFingerprint,
    )
  }

  removeOutway = flow(function* removeOutway(name) {
    yield this._managementApiClient.managementServiceDeleteOutway({
      name,
    })

    yield this.fetchAll()
  }).bind(this)

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
