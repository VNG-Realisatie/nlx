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

  constructor({ accessGrantStore, accessGrantData }) {
    makeAutoObservable(this)

    this.accessGrantStore = accessGrantStore
    this.update(accessGrantData)
  }

  update = (accessProofData) => {
    if (!accessProofData) {
      throw Error('Data required to update accessProof')
    }

    if (accessProofData.id) {
      this.id = accessProofData.id
    }

    if (accessProofData.organization) {
      this.organization.name = accessProofData.organization.name
      this.organization.serialNumber = accessProofData.organization.serialNumber
    }

    if (accessProofData.serviceName) {
      this.serviceName = accessProofData.serviceName
    }

    if (accessProofData.publicKeyFingerprint) {
      this.publicKeyFingerprint = accessProofData.publicKeyFingerprint
    }

    if (accessProofData.createdAt) {
      this.createdAt = new Date(accessProofData.createdAt)
    }

    if (accessProofData.revokedAt) {
      this.revokedAt = new Date(accessProofData.revokedAt)
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
