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

  updateFromServer = flow(function* updateFromServer(
    outgoingAccessRequestData,
  ) {
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

  create = flow(function* create({ organizationName, serviceName }) {
    const response = yield this.accessRequestRepository.createAccessRequest({
      organizationName,
      serviceName,
    })

    return new OutgoingAccessRequestModel({
      accessRequestData: response,
    })
  }).bind(this)

  retry = flow(function* retry(outgoingAccessRequestModel) {
    const response = yield this.accessRequestRepository.sendAccessRequest({
      organizationName: outgoingAccessRequestModel.organizationName,
      serviceName: outgoingAccessRequestModel.serviceName,
      id: outgoingAccessRequestModel.id,
    })

    yield this.updateFromServer(response)
  }).bind(this)
}

export default OutgoingAccessRequestStore
