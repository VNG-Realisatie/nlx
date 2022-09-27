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

  send = flow(function* create(
    organizationSerialNumber,
    serviceName,
    publicKeyPem,
  ) {
    const response =
      yield this._managementApiClient.managementSendAccessRequest({
        body: {
          organizationSerialNumber,
          serviceName,
          publicKeyPem,
        },
      })

    return new OutgoingAccessRequestModel({
      accessRequestData: response.outgoingAccessRequest,
    })
  }).bind(this)
}

export default OutgoingAccessRequestStore
