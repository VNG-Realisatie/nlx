// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, action, flow } from 'mobx'
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

export const incomingAccessRequestPropTypes = {
  id: string,
  organizationName: string.isRequired,
  serviceName: string.isRequired,
  state: string,
  createdAt: string,
  updatedAt: string,

  approve: func,
  error: string,
}

class IncomingAccessRequestModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  createdAt = ''
  updatedAt = ''

  constructor({
    store,
    accessRequestData,
    accessRequestRepository = AccessRequestRepository,
  }) {
    this.store = store || undefined
    this.accessRequestRepository = accessRequestRepository
    this.error = ''

    this.update(accessRequestData)
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

  approve = flow(function* send() {
    try {
      this.error = ''

      const { serviceName, id } = this
      yield this.accessRequestRepository.approveAccessRequest({
        serviceName,
        id,
      })

      this.store.fetchIncomingAccessRequests()
    } catch (e) {
      this.error = e.message
    }
  })
}

decorate(IncomingAccessRequestModel, {
  state: observable,
  update: action.bound,
  approve: action.bound,
  error: observable,
})

export const createIncomingAccessRequest = (...args) =>
  new IncomingAccessRequestModel(...args)

export default IncomingAccessRequestModel
