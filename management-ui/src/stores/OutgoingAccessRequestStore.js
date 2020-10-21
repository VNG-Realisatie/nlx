// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import AccessRequestRepository from '../domain/access-request-repository'
import OutgoingAccessRequestModel from '../models/OutgoingAccessRequestModel'

class OutgoingAccessRequestStore {
  outgoingAccessRequests = observable.map()

  constructor({
    rootStore,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this.accessRequestRepository = accessRequestRepository
  }

  // TODO: reconsider this method name
  // the method is created so we can create / update OutgoingAccessRequest instances
  // with data that has been loaded from other resources (eg. a directoryService resource
  // which includes the data for an OutgoingAccessRequest).
  loadOutgoingAccessRequest = flow(function* (outgoingAccessRequestData) {
    const cachedOutgoingAccessRequest = this.outgoingAccessRequests.get(
      outgoingAccessRequestData.id,
    )

    if (cachedOutgoingAccessRequest) {
      cachedOutgoingAccessRequest.update(outgoingAccessRequestData)
      return yield cachedOutgoingAccessRequest
    } else {
      const outgoingAccessRequest = new OutgoingAccessRequestModel({
        accessRequestData: outgoingAccessRequestData,
        outgoingAccessRequestStore: this,
      })

      this.outgoingAccessRequests.set(
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
      accessRequestData: response,
    })
  }).bind(this)

  retry = flow(function* (outgoingAccessRequestModel) {
    const response = yield this.accessRequestRepository.sendAccessRequest({
      organizationName: outgoingAccessRequestModel.organizationName,
      serviceName: outgoingAccessRequestModel.serviceName,
      id: outgoingAccessRequestModel.id,
    })

    yield this.loadOutgoingAccessRequest(response)
  }).bind(this)
}

export default OutgoingAccessRequestStore
