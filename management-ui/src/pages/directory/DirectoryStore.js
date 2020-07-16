// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, flow, action } from 'mobx'

import DirectoryServiceModel from '../../models/DirectoryServiceModel'

class DirectoryStore {
  services = []

  constructor(rootStore, domain) {
    this.rootStore = rootStore
    this.domain = domain

    // @TODO:
    // All stores that do async stuff now have isLoading and error.
    // Consider using a reqres model / state machine
    // See: https://benmccormick.org/2018/05/14/mobx-state-machines-and-flags/
    this.isLoading = false
    this.error = ''
  }

  getServices = flow(function* getServices() {
    this.isLoading = true
    this.error = ''

    try {
      const services = yield this.domain.getAll()
      this.services = services.map(
        (service) => new DirectoryServiceModel({ store: this, service }),
      )
    } catch (e) {
      this.error = e
    } finally {
      this.isLoading = false
    }
  })
}

decorate(DirectoryStore, {
  services: observable,
  isLoading: observable,
  error: observable,
  // flow'ed functions don't need action. This is so we can destructure it in component:
  getServices: action.bound,
})

export const createDirectoryStore = (...args) => new DirectoryStore(...args)

export default DirectoryStore
