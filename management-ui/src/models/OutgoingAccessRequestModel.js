// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'
import { string } from 'prop-types'

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
  createdAt: string,
  updatedAt: string,
  error: string,
}

class OutgoingAccessRequestModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  createdAt = ''
  updatedAt = ''
  error = ''

  constructor({ accessRequestData }) {
    makeAutoObservable(this)
    this.update(accessRequestData)
  }

  update(accessRequestData) {
    if (!accessRequestData) {
      return
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
      this.createdAt = accessRequestData.createdAt
    }

    if (accessRequestData.updatedAt) {
      this.updatedAt = accessRequestData.updatedAt
    }
  }

  get isCancelledOrRejected() {
    return (
      this.state === ACCESS_REQUEST_STATES.CANCELLED ||
      this.state === ACCESS_REQUEST_STATES.REJECTED
    )
  }
}

export const createAccessRequestInstance = (requestData) => {
  return new OutgoingAccessRequestModel({ accessRequestData: requestData })
}

export default OutgoingAccessRequestModel
