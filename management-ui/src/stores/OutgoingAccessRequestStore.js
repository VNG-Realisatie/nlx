// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, flow, makeAutoObservable } from 'mobx'
import AccessRequestRepository from '../domain/access-request-repository'
import OutgoingAccessRequestModel from '../models/OutgoingAccessRequestModel'

class OutgoingAccessRequestStore {
  constructor({
    rootStore,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this, {
      create: action.bound,
    })

    this.rootStore = rootStore
    this.accessRequestRepository = accessRequestRepository
  }

  create = flow(function* ({ organizationName, serviceName }) {
    const response = yield this.accessRequestRepository.createAccessRequest({
      organizationName,
      serviceName,
    })

    return new OutgoingAccessRequestModel({
      accessRequestData: {
        id: response.id,
        organizationName: response.organizationName,
        serviceName: response.serviceName,
        state: response.state,
        createdAt: response.createdAt,
        updatedAt: null,
      },
    })
  }).bind(this)
}

export default OutgoingAccessRequestStore
