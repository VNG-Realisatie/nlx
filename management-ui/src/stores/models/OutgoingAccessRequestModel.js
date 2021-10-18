// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'

export const ACCESS_REQUEST_STATES = {
  CREATED: 'CREATED',
  FAILED: 'FAILED',
  RECEIVED: 'RECEIVED',
  CANCELLED: 'CANCELLED',
  REJECTED: 'REJECTED',
  APPROVED: 'APPROVED',
}

class OutgoingAccessRequestModel {
  id = ''
  organization = {
    serialNumber: '',
    name: '',
  }

  serviceName = ''
  state = ''
  createdAt = null
  updatedAt = null
  errorDetails = null

  constructor({ accessRequestData, outgoingAccessRequestStore }) {
    makeAutoObservable(this)

    this.outgoingAccessRequestStore = outgoingAccessRequestStore

    this.update(accessRequestData)
  }

  update = (accessRequestData) => {
    if (!accessRequestData) {
      throw Error('Data required to update outgoingAccessRequest')
    }

    if (accessRequestData.id) {
      this.id = accessRequestData.id
    }

    if (accessRequestData.organization) {
      this.organization.serialNumber =
        accessRequestData.organization.serialNumber
      this.organization.name = accessRequestData.organization.name
    }

    if (accessRequestData.serviceName) {
      this.serviceName = accessRequestData.serviceName
    }

    if (accessRequestData.state) {
      this.state = accessRequestData.state
    }

    if (accessRequestData.createdAt) {
      this.createdAt = new Date(accessRequestData.createdAt)
    }

    if (accessRequestData.updatedAt) {
      this.updatedAt = new Date(accessRequestData.updatedAt)
    }

    if (accessRequestData.errorDetails) {
      this.errorDetails = accessRequestData.errorDetails
    }
  }

  retry = flow(function* retry() {
    yield this.outgoingAccessRequestStore.retry(this)
  }).bind(this)
}

export const createAccessRequestInstance = (requestData) => {
  return new OutgoingAccessRequestModel({ accessRequestData: requestData })
}

export default OutgoingAccessRequestModel
