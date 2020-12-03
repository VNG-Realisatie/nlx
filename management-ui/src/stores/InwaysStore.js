// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'
import InwayRepository from '../domain/inway-repository'
import InwayModel from '../models/InwayModel'

class InwaysStore {
  inways = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchInways calls being in flight.
  isFetching = false

  constructor({ rootStore, inwayRepository = InwayRepository }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this.inwayRepository = inwayRepository

    this.inways = []
    this.error = ''
    this.isInitiallyFetched = false
    this.isFetching = false
  }

  fetchInways = flow(function* fetchInways() {
    if (this.isFetching) {
      return
    }

    this.isFetching = true
    this.error = ''

    try {
      const inways = yield this.inwayRepository.getAll()
      this.inways = inways.map(
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
    const inwayData = yield this.inwayRepository.getByName(name)
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

export default InwaysStore
