// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

class AccessProofModel {
  id = ''
  organization = {
    serialNumber: '',
    name: '',
  }

  serviceName = ''
  createdAt = null
  revokedAt = null
  accessRequestId = null
  publicKeyFingerprint = ''

  constructor({ accessProofData }) {
    makeAutoObservable(this)

    this.update(accessProofData)
  }

  update = (accessProofData) => {
    if (!accessProofData) {
      throw Error('Data required to update accessProof')
    }

    if (accessProofData.id) {
      this.id = accessProofData.id
    }

    if (accessProofData.organization) {
      this.organization.serialNumber = accessProofData.organization.serialNumber
      this.organization.name =
        accessProofData.organization.name ||
        accessProofData.organization.serialNumber
    }

    if (accessProofData.serviceName) {
      this.serviceName = accessProofData.serviceName
    }

    if (accessProofData.createdAt) {
      this.createdAt = new Date(accessProofData.createdAt)
    }

    if (accessProofData.revokedAt) {
      this.revokedAt = new Date(accessProofData.revokedAt)
    }

    if (accessProofData.accessRequestId) {
      this.accessRequestId = accessProofData.accessRequestId
    }

    if (accessProofData.publicKeyFingerprint) {
      this.publicKeyFingerprint = accessProofData.publicKeyFingerprint
    }
  }
}

export default AccessProofModel
