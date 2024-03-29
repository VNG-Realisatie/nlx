// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'

export const ACCESS_REQUEST_STATES = {
  FAILED: 'ACCESS_REQUEST_STATE_FAILED',
  RECEIVED: 'ACCESS_REQUEST_STATE_RECEIVED',
  REJECTED: 'ACCESS_REQUEST_STATE_REJECTED',
  APPROVED: 'ACCESS_REQUEST_STATE_APPROVED',
  WITHDRAWN: 'ACCESS_REQUEST_STATE_WITHDRAWN',
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
  publicKeyFingerprint = ''
  errorDetails = {
    cause: null,
  }

  constructor({ accessRequestData, outgoingAccessRequestStore }) {
    makeAutoObservable(this)

    this.outgoingAccessRequestStore = outgoingAccessRequestStore

    this.update(accessRequestData)
  }

  update = (accessRequestData) => {
    if (!accessRequestData) {
      throw Error('Data required to update OutgoingAccessRequest')
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

    if (accessRequestData.errorDetails) {
      this.errorDetails.cause = accessRequestData.errorDetails.cause
    }
  }

  terminate = flow(function* reject() {
    try {
      yield this.outgoingAccessRequestStore.terminate(this)
    } catch (error) {
      console.error('Failed to terminate access: ', error.message)
      throw error
    }
  }).bind(this)

  withdraw = flow(function* reject() {
    try {
      yield this.outgoingAccessRequestStore.withdraw(this)
    } catch (error) {
      console.error('Failed to cancel access: ', error.message)
      throw error
    }
  }).bind(this)
}

export default OutgoingAccessRequestModel
