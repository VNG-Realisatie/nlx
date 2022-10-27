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

  terminate = flow(function* terminate(outgoingAccessRequest) {
    yield this._managementApiClient.managementServiceTerminateAccessProof({
      organizationSerialNumber: outgoingAccessRequest.organization.serialNumber,
      serviceName: outgoingAccessRequest.serviceName,
      publicKeyFingerprint: outgoingAccessRequest.publicKeyFingerprint,
    })
  }).bind(this)

  withdraw = flow(function* withdraw(outgoingAccessRequest) {
    yield this._managementApiClient.managementServiceWithdrawOutgoingAccessRequest(
      {
        organizationSerialNumber:
          outgoingAccessRequest.organization.serialNumber,
        serviceName: outgoingAccessRequest.serviceName,
        publicKeyFingerprint: outgoingAccessRequest.publicKeyFingerprint,
      },
    )
  }).bind(this)

  send = flow(function* create(
    organizationSerialNumber,
    serviceName,
    publicKeyPem,
  ) {
    const response =
      yield this._managementApiClient.managementServiceSendAccessRequest({
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
