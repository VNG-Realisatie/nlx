// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow, action } from 'mobx'
import InwayRepository from '../domain/inway-repository'
import { createInway } from '../models/InwayModel'

class InwaysStore {
  inways = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchInways calls being in flight.
  isFetching = false

  constructor({ rootStore, inwayRepository = InwayRepository }) {
    makeAutoObservable(this, {
      selectInway: action.bound,
    })

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
      this.inways = inways.map((inway) => createInway({ store: this, inway }))
    } catch (e) {
      this.error = e
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  }).bind(this)

  selectInway(inwayName) {
    const inwayModel = this.inways.find((inway) => inway.name === inwayName)

    if (inwayModel) {
      inwayModel.fetch()
    }

    return inwayModel
  }
}

export const createInwaysStore = (...args) => new InwaysStore(...args)

export default InwaysStore
