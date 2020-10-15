// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'
import { string, func } from 'prop-types'

import AccessRequestRepository from '../domain/access-request-repository'

export const ACCESS_REQUEST_STATES = {
  CREATED: 'CREATED',
  FAILED: 'FAILED',
  RECEIVED: 'RECEIVED',
  CANCELLED: 'CANCELLED',
  REJECTED: 'REJECTED',
  APPROVED: 'APPROVED',
}

export const UNSUCCESSFUL_ACCESS_REQUEST_STATES = [
  ACCESS_REQUEST_STATES.CANCELLED,
  ACCESS_REQUEST_STATES.REJECTED,
]

export const outgoingAccessRequestPropTypes = {
  id: string,
  organizationName: string.isRequired,
  serviceName: string.isRequired,
  state: string,
  createdAt: string,
  updatedAt: string,

  send: func,
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

  constructor({
    accessRequestData,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this)

    this.accessRequestRepository = accessRequestRepository
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

  send = flow(function* send() {
    if (this.id) {
      console.error('Request was already sent, ignoring request')
      return
    }

    try {
      this.error = ''

      this.update({ state: ACCESS_REQUEST_STATES.CREATED })

      const result = yield this.accessRequestRepository.createAccessRequest({
        organizationName: this.organizationName,
        serviceName: this.serviceName,
      })

      // Hydrate the object with response from server
      this.update(result)
    } catch (e) {
      this.error = e
      console.error(e)
    }
  })
}

export const createAccessRequestInstance = (requestData) => {
  return new OutgoingAccessRequestModel({ accessRequestData: requestData })
}

export default OutgoingAccessRequestModel
