// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { action, flow, makeAutoObservable, observable } from 'mobx'
import AccessRequestRepository from '../domain/access-request-repository'
import OutgoingAccessRequestModel from '../models/OutgoingAccessRequestModel'

class OutgoingAccessRequestStore {
  outgoingAccessRequests = observable.map()

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

  getOutgoingAccessRequest(id) {
    return this.outgoingAccessRequests.get(id)
  }

  setOutgoingAccessRequest(id, model) {
    return this.outgoingAccessRequests.set(id, model)
  }

  // TODO: reconsider this method name
  // the method is created so we can create / update OutgoingAccessRequest instances
  // with data that has been loaded from other resources (eg. a directoryService resource
  // which includes the data for an OutgoingAccessRequest).
  loadOutgoingAccessRequest = flow(function* (outgoingAccessRequestData) {
    const cachedOutgoingAccessRequest = this.getOutgoingAccessRequest(
      outgoingAccessRequestData.id,
    )

    if (cachedOutgoingAccessRequest) {
      cachedOutgoingAccessRequest.update(outgoingAccessRequestData)
      return yield cachedOutgoingAccessRequest
    } else {
      const outgoingAccessRequest = new OutgoingAccessRequestModel({
        accessRequestData: {
          id: outgoingAccessRequestData.id,
          organizationName: outgoingAccessRequestData.organizationName,
          serviceName: outgoingAccessRequestData.serviceName,
          state: outgoingAccessRequestData.state,
          createdAt: outgoingAccessRequestData.createdAt,
          updatedAt: outgoingAccessRequestData.updatedAt,
        },
        outgoingAccessRequestStore: this,
      })

      this.setOutgoingAccessRequest(
        outgoingAccessRequest.id,
        outgoingAccessRequest,
      )

      return yield outgoingAccessRequest
    }
  }).bind(this)

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
