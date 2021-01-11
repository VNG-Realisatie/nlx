// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import InwayModel from './models/InwayModel'

class InwayStore {
  inways = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchInways calls being in flight.
  isFetching = false

  constructor({ rootStore, managementApiClient }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this._managementApiClient = managementApiClient
  }

  fetchInways = flow(function* fetchInways() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true
    this.error = ''

    try {
      const result = yield this._managementApiClient.managementListInways()
      this.inways = result.inways.map(
        (inway) => new InwayModel({ store: this, inway }),
      )
    } catch (e) {
      this.error = e
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

  fetch = flow(function* fetch({ name }) {
    const inwayData = yield this._managementApiClient.managementGetInway({
      name,
    })
    let inway = this.getInway({ name })

    if (!inway) {
      inway = this._updateFromServer(inwayData)
      this.inways.push(inway)
      return inway
    }

    return this._updateFromServer(inwayData)
  }).bind(this)

  getInway = ({ name }) => {
    return this.inways.find((inway) => inway.name === name)
  }

  _updateFromServer(inwayData) {
    const cachedInway = this.getInway({
      name: inwayData.name,
    })

    if (cachedInway) {
      cachedInway.with(inwayData)
      return cachedInway
    }

    return new InwayModel({
      store: this,
      inway: inwayData,
    })
  }
}

export default InwayStore
