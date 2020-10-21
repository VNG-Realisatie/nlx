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

export const incomingAccessRequestPropTypes = {
  id: string,
  organizationName: string.isRequired,
  serviceName: string.isRequired,
  state: string,
  createdAt: string,
  updatedAt: string,

  approve: func,
  reject: func,
  error: string,
}

class IncomingAccessRequestModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  createdAt = ''
  updatedAt = ''
  error = ''

  constructor({
    store,
    accessRequestData,
    accessRequestRepository = AccessRequestRepository,
  }) {
    makeAutoObservable(this)

    this.store = store || undefined
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

  approve = flow(function* send() {
    try {
      this.error = ''

      const { serviceName, id } = this
      yield this.accessRequestRepository.approveIncomingAccessRequest({
        serviceName,
        id,
      })

      this.store.fetchIncomingAccessRequests()
    } catch (e) {
      this.error = e.message
    }
  }).bind(this)

  reject = flow(function* flow() {
    try {
      this.error = ''

      const { serviceName, id } = this
      yield this.accessRequestRepository.rejectIncomingAccessRequest({
        serviceName,
        id,
      })

      this.store.fetchIncomingAccessRequests()
    } catch (e) {
      this.error = e.message
    }
  }).bind(this)
}

export const createIncomingAccessRequest = (...args) =>
  new IncomingAccessRequestModel(...args)

export default IncomingAccessRequestModel
