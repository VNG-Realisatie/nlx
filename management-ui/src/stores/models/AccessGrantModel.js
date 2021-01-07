// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'

// TODO: move `models` directory inside `stores`

class AccessGrantModel {
  id = ''
  organizationName = ''
  serviceName = ''
  publicKeyFingerprint = ''
  createdAt = null
  revokedAt = null

  error = ''

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

    if (accessProofData.organizationName) {
      this.organizationName = accessProofData.organizationName
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
      this.error = ''
      yield this.accessGrantStore.revokeAccessGrant(this)
    } catch (e) {
      this.error = e.message
    }
  }).bind(this)
}

export default AccessGrantModel
