// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, flow, action } from 'mobx'

import DirectoryRepository from '../../domain/directory-repository'
import DirectoryServiceModel from '../../models/DirectoryServiceModel'

class DirectoryStore {
  services = []

  constructor({ rootStore, domain = DirectoryRepository }) {
    this.rootStore = rootStore
    this.domain = domain

    // @TODO:
    // All stores/models that do async stuff now have isLoading and error.
    // Consider using a reqres model / state machine
    // See: https://benmccormick.org/2018/05/14/mobx-state-machines-and-flags/
    this.isLoading = false
    this.error = ''
  }

  fetchServices = flow(function* fetchServices() {
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

  selectService = ({ organizationName, serviceName }) => {
    return this.services.find(
      (service) =>
        service.organizationName === organizationName &&
        service.serviceName === serviceName,
    )
  }
}

decorate(DirectoryStore, {
  services: observable,
  isLoading: observable,
  error: observable,
  fetchServices: action.bound,
})

export const createDirectoryStore = (...args) => new DirectoryStore(...args)

export default DirectoryStore
