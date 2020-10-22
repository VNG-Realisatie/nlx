// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import AccessRequestRepository from '../domain/access-request-repository'
import IncomingAccessRequestModel from '../models/IncomingAccessRequestModel'

class IncomingAccessRequestsStore {
  incomingAccessRequests = observable.map()

  constructor({
    rootStore,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this.accessRequestRepository = accessRequestRepository
  }

  updateFromServer = (incomingAccessRequestData) => {
    const cachedIncomingAccessRequest = this.incomingAccessRequests.get(
      incomingAccessRequestData.id,
    )

    if (cachedIncomingAccessRequest) {
      cachedIncomingAccessRequest.update(incomingAccessRequestData)
      return cachedIncomingAccessRequest
    }

    const incomingAccessRequest = new IncomingAccessRequestModel({
      accessRequestData: incomingAccessRequestData,
    })

    this.incomingAccessRequests.set(
      incomingAccessRequest.id,
      incomingAccessRequest,
    )

    return incomingAccessRequest
  }

  fetchForService = flow(function* fetchForService(serviceModel) {
    const response = yield this.accessRequestRepository.listIncomingAccessRequests(
      serviceModel.name,
    )

    const incomingAccessRequestModels = response.map((accessRequest) =>
      this.updateFromServer(accessRequest),
    )

    yield incomingAccessRequestModels
  }).bind(this)

  getForService = (serviceModel) => {
    const arrayOfModels = [...this.incomingAccessRequests.values()]

    return arrayOfModels.filter(
      (incomingAccessRequestModel) =>
        incomingAccessRequestModel.serviceName === serviceModel.name,
    )
  }
}

export default IncomingAccessRequestsStore
