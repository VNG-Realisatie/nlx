// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

class Organization {
  serialNumber = ''
  name = ''

  constructor(serialNumber, name) {
    this.serialNumber = serialNumber
    this.name = name
  }
}

class IncomingOrderModel {
  _description = null
  _delegator = null
  _reference = null
  _services = null
  _revokedAt = null
  _validFrom = null
  _validUntil = null

  constructor({ orderStore, orderData }) {
    makeAutoObservable(this)

    this.orderStore = orderStore
    this.update(orderData)
  }

  get description() {
    return this._description
  }

  get delegator() {
    return this._delegator
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

  get services() {
    return this._services
  }

  update = (orderData) => {
    if (!orderData) {
      throw Error('Data required to update incoming order')
    }

    if (orderData.reference) {
      this._reference = orderData.reference
    }

    if (orderData.description) {
      this._description = orderData.description
    }

    if (orderData.delegator) {
      this._delegator = new Organization(
        orderData.delegator.serialNumber,
        orderData.delegator.name || orderData.delegator.serialNumber,
      )
    }

    if (orderData.services) {
      this._services = orderData.services
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
  }
}

export default IncomingOrderModel
