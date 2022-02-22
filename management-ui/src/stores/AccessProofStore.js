// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable } from 'mobx'

import AccessProofModel from './models/AccessProofModel'

class AccessProofStore {
  _accessProofs = observable.map()

  constructor() {
    makeAutoObservable(this)
  }

  get accessProofs() {
    return [...this._accessProofs.values()]
  }

  updateFromServer = (accessProofData) => {
    if (!accessProofData) return null

    const cachedAccessProof = this._accessProofs.get(accessProofData.id)

    if (cachedAccessProof) {
      cachedAccessProof.update(accessProofData)
      return cachedAccessProof
    }

    const accessProof = new AccessProofModel({ accessProofData })

    this._accessProofs.set(accessProof.id, accessProof)

    return accessProof
  }
}

export default AccessProofStore
