// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable } from 'mobx'
import AccessRequestRepository from '../domain/access-request-repository'

class IncomingAccessRequestStore {
  incomingAccessRequests = observable.map()

  constructor({
    rootStore,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this.accessRequestRepository = accessRequestRepository
  }
}

export default IncomingAccessRequestStore
