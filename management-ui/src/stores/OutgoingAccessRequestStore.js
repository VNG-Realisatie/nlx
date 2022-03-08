// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import OutgoingAccessRequestModel from './models/OutgoingAccessRequestModel'

class OutgoingAccessRequestStore {
  outgoingAccessRequests = observable.map()

  constructor({ rootStore, managementApiClient }) {
    makeAutoObservable(this)

    this.rootStore = rootStore
    this._managementApiClient = managementApiClient
  }

  updateFromServer = (outgoingAccessRequestData) => {
    if (!outgoingAccessRequestData) return null

    const cachedOutgoingAccessRequest = this.outgoingAccessRequests.get(
      outgoingAccessRequestData.id,
    )

    if (cachedOutgoingAccessRequest) {
      cachedOutgoingAccessRequest.update(outgoingAccessRequestData)
      return cachedOutgoingAccessRequest
    }

    const outgoingAccessRequest = new OutgoingAccessRequestModel({
      accessRequestData: outgoingAccessRequestData,
      outgoingAccessRequestStore: this,
    })

    this.outgoingAccessRequests.set(
      outgoingAccessRequest.id,
      outgoingAccessRequest,
    )

    return outgoingAccessRequest
  }

  create = flow(function* create(
    organizationSerialNumber,
    serviceName,
    publicKeyPEM,
  ) {
    const response =
      yield this._managementApiClient.managementCreateAccessRequest({
        body: {
          organizationSerialNumber,
          serviceName,
          publicKeyPEM,
        },
      })

    return new OutgoingAccessRequestModel({
      accessRequestData: response,
    })
  }).bind(this)

  retry = flow(function* retry(outgoingAccessRequestModel) {
    const response =
      yield this._managementApiClient.managementSendAccessRequest({
        accessRequestID: outgoingAccessRequestModel.id,
      })
    yield this.updateFromServer(response)
  }).bind(this)
}

export default OutgoingAccessRequestStore
