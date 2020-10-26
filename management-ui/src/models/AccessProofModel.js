// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'
import { string, instanceOf } from 'prop-types'

export const accessProofPropTypes = {
  id: string.isRequired,
  organizationName: string,
  serviceName: string,
  createdAt: instanceOf(Date),
  updatedAt: instanceOf(Date),
}

class AccessProofModel {
  id = ''
  organizationName = ''
  serviceName = ''
  createdAt = null
  revokedAt = null

  static verifyInstance(object, objectName = 'given object') {
    if (object && !(object instanceof AccessProofModel)) {
      throw new Error(
        `The ${objectName} should be an instance of the AccessProofModel`,
      )
    }
  }

  constructor({ accessProofData, accessProofStore }) {
    makeAutoObservable(this)

    this.accessProofStore = accessProofStore

    this.update(accessProofData)
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

    if (accessProofData.createdAt) {
      this.createdAt = new Date(accessProofData.createdAt)
    }

    if (accessProofData.revokedAt) {
      this.revokedAt = new Date(accessProofData.revokedAt)
    }
  }
}

export default AccessProofModel
