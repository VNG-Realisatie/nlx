// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, flow, action } from 'mobx'

import DirectoryRepository from '../../domain/directory-repository'
import { createDirectoryService } from '../../models/DirectoryServiceModel'

class DirectoryStore {
  services = []
  isLoading = false
  isReady = false

  constructor({ rootStore, domain = DirectoryRepository }) {
    this.rootStore = rootStore
    this.domain = domain

    // @TODO:
    // All stores that do async stuff have isReady and error.
    // Consider using a reqres model / state machine
    // See: https://benmccormick.org/2018/05/14/mobx-state-machines-and-flags/
    this.isLoading = false
    this.isReady = false
    this.error = ''
  }

  fetchServices = flow(function* fetchServices() {
    // This prevents making concurrent fetch calls being triggered by rerendering
    if (this.isLoading) {
      return
    }
    this.isLoading = true
    this.error = ''

    try {
      const services = yield this.domain.getAll()
      this.services = services.map((service) =>
        createDirectoryService({ store: this, service }),
      )
    } catch (e) {
      this.error = e
    } finally {
      this.isReady = true
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
  isReady: observable,
  error: observable,
  fetchServices: action.bound,
})

export const createDirectoryStore = (...args) => new DirectoryStore(...args)

export default DirectoryStore
