// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, computed, action, flow } from 'mobx'
import { string, func, bool } from 'prop-types'

import AccessRequestRepository from '../domain/access-request-repository'

export const ACCESS_REQUEST_STATES = {
  CREATED: 'CREATED',
  FAILED: 'FAILED',
  CANCELLED: 'CANCELLED',
  REJECTED: 'REJECTED',
  ACCEPTED: 'ACCEPTED',
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
  isOpen: bool,

  error: string,
}

class OutgoingAccessRequestModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  createdAt = ''
  updatedAt = ''

  get isOpen() {
    return !UNSUCCESSFUL_ACCESS_REQUEST_STATES.includes(this.state)
  }

  constructor({ accessRequestData, domain = AccessRequestRepository }) {
    this.domain = domain
    this.update(accessRequestData)

    this.error = ''
  }

  update(accessRequestData) {
    if (accessRequestData) {
      Object.keys(accessRequestData)
        .filter((key) => key in this)
        .forEach((key) => {
          this[key] = accessRequestData[key]
        })
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

      const result = yield this.domain.requestAccess({
        organizationName: this.organizationName,
        serviceName: this.serviceName,
      })

      // Hydrate the object with response from server
      this.update(result)
    } catch (e) {
      this.error = e
      throw e
    }
  })
}

decorate(OutgoingAccessRequestModel, {
  state: observable,
  isOpen: computed,
  update: action.bound,
  send: action.bound,
})

export const createAccessRequestInstance = (requestData) => {
  return new OutgoingAccessRequestModel({ accessRequestData: requestData })
}

export default OutgoingAccessRequestModel
