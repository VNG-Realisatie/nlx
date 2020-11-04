// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable } from 'mobx'
import AccessRequestRepository from '../domain/access-request-repository'
import IncomingAccessRequestModel from '../models/IncomingAccessRequestModel'

class IncomingAccessRequestsStore {
  incomingAccessRequests = observable.map()

  constructor({ accessRequestRepository = AccessRequestRepository }) {
    makeAutoObservable(this)

    this.accessRequestRepository = accessRequestRepository
  }

  updateFromServer = (incomingAccessRequestData) => {
    if (!incomingAccessRequestData) return null

    const cachedIncomingAccessRequest = this.incomingAccessRequests.get(
      incomingAccessRequestData.id,
    )

    if (cachedIncomingAccessRequest) {
      cachedIncomingAccessRequest.update(incomingAccessRequestData)
      return cachedIncomingAccessRequest
    }

    const incomingAccessRequest = new IncomingAccessRequestModel({
      incomingAccessRequestStore: this,
      accessRequestData: incomingAccessRequestData,
    })

    this.incomingAccessRequests.set(
      incomingAccessRequest.id,
      incomingAccessRequest,
    )

    return incomingAccessRequest
  }

  fetchForService = async ({ name }) => {
    const accessRequests = await this.accessRequestRepository.fetchByServiceName(
      name,
    )
    accessRequests.map((accessRequest) => this.updateFromServer(accessRequest))
  }

  getForService = (serviceModel) => {
    const arrayOfModels = [...this.incomingAccessRequests.values()]

    return arrayOfModels.filter(
      (incomingAccessRequestModel) =>
        incomingAccessRequestModel.serviceName === serviceModel.name,
    )
  }

  approveAccessRequest = async ({ serviceName, id }) => {
    await this.accessRequestRepository.approveIncomingAccessRequest({
      serviceName,
      id,
    })

    this.fetchForService({ name: serviceName })
  }

  rejectAccessRequest = async ({ serviceName, id }) => {
    await this.accessRequestRepository.rejectIncomingAccessRequest({
      serviceName,
      id,
    })

    this.fetchForService({ name: serviceName })
  }
}

export default IncomingAccessRequestsStore
