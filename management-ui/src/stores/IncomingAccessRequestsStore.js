// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable } from 'mobx'
import IncomingAccessRequestModel from './models/IncomingAccessRequestModel'

class IncomingAccessRequestsStore {
  incomingAccessRequests = observable.map()

  constructor({ managementApiClient }) {
    makeAutoObservable(this)

    this._managementApiClient = managementApiClient
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
    const result = await this._managementApiClient.managementListIncomingAccessRequest(
      {
        serviceName: name,
      },
    )
    result.accessRequests.map((accessRequest) =>
      this.updateFromServer(accessRequest),
    )
  }

  getForService = (serviceModel) => {
    const arrayOfModels = [...this.incomingAccessRequests.values()]

    return arrayOfModels.filter(
      (incomingAccessRequestModel) =>
        incomingAccessRequestModel.serviceName === serviceModel.name,
    )
  }

  approveAccessRequest = async ({ serviceName, id }) => {
    await this._managementApiClient.managementApproveIncomingAccessRequest({
      serviceName,
      accessRequestID: id,
    })
    this.fetchForService({ name: serviceName })
  }

  rejectAccessRequest = async ({ serviceName, id }) => {
    await this._managementApiClient.managementRejectIncomingAccessRequest({
      serviceName,
      accessRequestID: id,
    })
    this.fetchForService({ name: serviceName })
  }
}

export default IncomingAccessRequestsStore
