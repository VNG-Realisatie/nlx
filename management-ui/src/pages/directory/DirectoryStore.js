// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, flow, action } from 'mobx'

import DirectoryRepository from '../../domain/directory-repository'
import { createDirectoryService } from '../../models/DirectoryServiceModel'

class DirectoryStore {
  services = []
  error = ''
  // This is set to true after the first call has been made. Regardless of success.
  isInitiallyFetched = false
  // This is internal state to prevent concurrent fetchServices calls being in flight.
  isFetching = false

  constructor({ rootStore, domain = DirectoryRepository }) {
    this.rootStore = rootStore
    this.domain = domain

    this.services = []
    // @TODO:
    // All stores that do async stuff have isInitiallyFetched and error.
    // Consider using a reqres model / state machine
    // See: https://benmccormick.org/2018/05/14/mobx-state-machines-and-flags/
    this.error = ''
    this.isInitiallyFetched = false
    this.isFetching = false
  }

  fetchServices = flow(function* fetchServices() {
    if (this.isFetching) {
      return
    }
    this.isFetching = true
    this.error = ''

    try {
      const services = yield this.domain.getAll()
      this.services = services.map((service) =>
        createDirectoryService({ store: this, service }),
      )
    } catch (e) {
      this.error = e
    } finally {
      this.isInitiallyFetched = true
      this.isFetching = false
    }
  })

  selectService = ({ organizationName, serviceName }) => {
    const directoryServiceModel = this.services.find(
      (service) =>
        service.organizationName === organizationName &&
        service.serviceName === serviceName,
    )
    if (directoryServiceModel) {
      directoryServiceModel.fetch()
    }
    return directoryServiceModel
  }
}

decorate(DirectoryStore, {
  services: observable,
  isInitiallyFetched: observable,
  error: observable,
  fetchServices: action.bound,
})

export const createDirectoryStore = (...args) => new DirectoryStore(...args)

export default DirectoryStore
