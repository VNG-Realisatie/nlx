// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'

export const STATES = {
  CREATED: 'CREATED',
  FAILED: 'FAILED',
  RECEIVED: 'RECEIVED',
  CANCELLED: 'CANCELLED',
  REJECTED: 'REJECTED',
  APPROVED: 'APPROVED',
}

class IncomingAccessRequestModel {
  id = ''
  organization = {
    serialNumber: '',
    name: '',
  }

  serviceName = ''
  state = ''
  createdAt = ''
  updatedAt = ''
  publicKeyFingerprint = ''

  constructor({ incomingAccessRequestStore, accessRequestData }) {
    makeAutoObservable(this)

    this.incomingAccessRequestStore = incomingAccessRequestStore
    this.update(accessRequestData)
  }

  get isResolved() {
    return !(this.state === STATES.CREATED || this.state === STATES.RECEIVED)
  }

  update(accessRequestData) {
    if (!accessRequestData) {
      return
    }

    if (accessRequestData.id) {
      this.id = accessRequestData.id
    }

    if (accessRequestData.organization) {
      this.organization.serialNumber =
        accessRequestData.organization.serialNumber
      this.organization.name =
        accessRequestData.organization.name ||
        accessRequestData.organization.serialNumber
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

    if (accessRequestData.publicKeyFingerprint) {
      this.publicKeyFingerprint = accessRequestData.publicKeyFingerprint
    }
  }

  approve = flow(function* approve() {
    try {
      yield this.incomingAccessRequestStore.approveAccessRequest(this)
    } catch (error) {
      console.error('Failed to approve access request: ', error)
      throw error
    }
  }).bind(this)

  reject = flow(function* reject() {
    try {
      yield this.incomingAccessRequestStore.rejectAccessRequest(this)
    } catch (error) {
      console.error('Failed to reject access request: ', error.message)
      throw error
    }
  }).bind(this)
}

export default IncomingAccessRequestModel
