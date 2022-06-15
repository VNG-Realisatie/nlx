// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import AccessProofModel from './AccessProofModel'

function throwErrorWhenNotInstanceOf(object, model) {
  if (object && !(object instanceof model)) {
    throw new Error(`Object should be an instance of ${model}`)
  }
}

class Organization {
  serialNumber = ''
  name = ''

  constructor(serialNumber, name) {
    this.serialNumber = serialNumber
    this.name = name
  }
}

class OutgoingOrderModel {
  _description = null
  _delegatee = null
  _reference = null
  _publicKeyPem = null
  _revokedAt = null
  _validFrom = null
  _validUntil = null
  _accessProofs = []

  constructor({ orderStore, orderData, accessProofs }) {
    makeAutoObservable(this)

    this.orderStore = orderStore
    this.update({ orderData, accessProofs })
  }

  get description() {
    return this._description
  }

  get publicKeyPem() {
    return this._publicKeyPem
  }

  get delegatee() {
    return this._delegatee
  }

  get reference() {
    return this._reference
  }

  get revokedAt() {
    return this._revokedAt
  }

  get validFrom() {
    return this._validFrom
  }

  get validUntil() {
    return this._validUntil
  }

  get accessProofs() {
    return this._accessProofs
  }

  get publicKeyPEM() {
    return this._publicKeyPem
  }

  revoke = flow(function* revoke() {
    yield this.orderStore.revokeOutgoing(this)
  }).bind(this)

  update = ({ orderData, accessProofs }) => {
    if (!orderData) {
      throw Error('Data required to update outgoing order')
    }

    if (orderData.reference) {
      this._reference = orderData.reference
    }

    if (orderData.description) {
      this._description = orderData.description
    }

    if (orderData.delegatee) {
      this._delegatee = new Organization(
        orderData.delegatee.serialNumber,
        orderData.delegatee.name,
      )
    }

    if (orderData.publicKeyPem) {
      this._publicKeyPem = orderData.publicKeyPem
    }

    if (orderData.revokedAt) {
      this._revokedAt = new Date(orderData.revokedAt)
    }

    if (orderData.validFrom) {
      this._validFrom = new Date(orderData.validFrom)
    }

    if (orderData.validUntil) {
      this._validUntil = new Date(orderData.validUntil)
    }

    if (accessProofs) {
      this._accessProofs = []

      accessProofs.forEach((accessProof) => {
        throwErrorWhenNotInstanceOf(accessProof, AccessProofModel)
        this._accessProofs.push(accessProof)
      })
    }
  }
}

export default OutgoingOrderModel
