// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { decorate, observable, computed, action, flow } from 'mobx'

import AccessRequestRepository from '../domain/access-request-repository'

export const accessRequestStates = {
  CREATED: 'CREATED',
  FAILED: 'FAILED',
  CANCELLED: 'CANCELLED',
  REJECTED: 'REJECTED',
  ACCEPTED: 'ACCEPTED',
}

class OutgoingAccessRequestModel {
  id = ''
  organizationName = ''
  serviceName = ''
  state = ''
  createdAt = ''
  updatedAt = ''

  get hasUnsuccessfulEndstate() {
    const { CANCELLED, REJECTED } = accessRequestStates
    return [CANCELLED, REJECTED].includes(this.state)
  }

  constructor({ json, domain = AccessRequestRepository }) {
    this.domain = domain
    this.update(json)

    // Currently not used, but part of pattern
    this.isLoading = false
    this.error = ''
  }

  update(json) {
    if (json) {
      Object.keys(json).forEach((key) => {
        this[key] = json[key]
      })
    }
  }

  send = flow(function* send() {
    if (this.id) {
      console.error('Request was already sent, ignoring request')
      return
    }

    try {
      this.isLoading = true
      this.error = ''

      yield this.update({ state: accessRequestStates.CREATED })

      const result = yield this.domain.requestAccess({
        organizationName: this.organizationName,
        serviceName: this.serviceName,
      })

      // Hydrate the object with response from server
      yield this.update(result)
    } catch (e) {
      this.error = e
      throw e
    } finally {
      this.isLoading = false
    }
  })
}

decorate(OutgoingAccessRequestModel, {
  state: observable,
  hasUnsuccessfulEndstate: computed,
  update: action.bound,
  send: action.bound,
})

export const createAccessRequestInstance = (requestData) => {
  return new OutgoingAccessRequestModel({ json: requestData })
}

export default OutgoingAccessRequestModel
