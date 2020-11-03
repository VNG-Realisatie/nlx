// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import { func, object, string, instanceOf } from 'prop-types'

export const ACCESS_REQUEST_STATES = {
  CREATED: 'CREATED',
  FAILED: 'FAILED',
  RECEIVED: 'RECEIVED',
  CANCELLED: 'CANCELLED',
  REJECTED: 'REJECTED',
  APPROVED: 'APPROVED',
}

export const outgoingAccessRequestPropTypes = {
  id: string,
  organizationName: string.isRequired,
  serviceName: string.isRequired,
  state: string,
  createdAt: instanceOf(Date),
  updatedAt: instanceOf(Date),
  error: string,
  errorDetails: object,

  retry: func,
}

class OutgoingAccessRequestModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  createdAt = null
  updatedAt = null

  error = ''

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

    if (accessRequestData.organizationName) {
      this.organizationName = accessRequestData.organizationName
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
