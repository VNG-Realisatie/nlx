// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'

class AccessGrantModel {
  id = ''
  organization = {
    name: '',
    serialNumber: '',
  }

  serviceName = ''
  publicKeyFingerprint = ''
  createdAt = null
  revokedAt = null
  terminatedAt = null

  constructor({ accessGrantStore, accessGrantData }) {
    makeAutoObservable(this)

    this.accessGrantStore = accessGrantStore
    this.update(accessGrantData)
  }

  update = (accessGrantData) => {
    if (!accessGrantData) {
      throw Error('Data required to update accessProof')
    }

    if (accessGrantData.id) {
      this.id = accessGrantData.id
    }

    if (accessGrantData.organization) {
      this.organization.name =
        accessGrantData.organization.name ||
        accessGrantData.organization.serialNumber
      this.organization.serialNumber = accessGrantData.organization.serialNumber
    }

    if (accessGrantData.serviceName) {
      this.serviceName = accessGrantData.serviceName
    }

    if (accessGrantData.publicKeyFingerprint) {
      this.publicKeyFingerprint = accessGrantData.publicKeyFingerprint
    }

    if (accessGrantData.createdAt) {
      this.createdAt = new Date(accessGrantData.createdAt)
    }

    if (accessGrantData.revokedAt) {
      this.revokedAt = new Date(accessGrantData.revokedAt)
    }

    if (accessGrantData.terminatedAt) {
      this.terminatedAt = new Date(accessGrantData.terminatedAt)
    }
  }

  revoke = flow(function* revoke() {
    try {
      yield this.accessGrantStore.revokeAccessGrant(this)
    } catch (err) {
      console.error(err)
      throw err
    }
  }).bind(this)
}

export default AccessGrantModel
