// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable } from 'mobx'

import AccessProofModel from '../models/AccessProofModel'

class AccessProofStore {
  accessProofs = observable.map()

  constructor() {
    makeAutoObservable(this)
  }

  updateFromServer = (accessProofData) => {
    if (!accessProofData) return null

    const cachedAccessProof = this.accessProofs.get(accessProofData.id)

    if (cachedAccessProof) {
      cachedAccessProof.update(accessProofData)
      return cachedAccessProof
    } else {
      const accessProof = new AccessProofModel({
        accessProofStore: this,
        accessProofData,
      })

      this.accessProofs.set(accessProof.id, accessProof)

      return accessProof
    }
  }
}

export default AccessProofStore
