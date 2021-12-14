// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const DIRECTION_IN = 'IN'
export const DIRECTION_OUT = 'OUT'

class Organization {
  serialNumber = ''

  constructor(serialNumber) {
    this.serialNumber = serialNumber
  }
}

class Order {
  delegator = ''
  reference = ''

  constructor({ delegator, reference }) {
    this.delegator = delegator
    this.reference = reference
  }
}

class TransactionLogModel {
  direction = null
  source = null
  destination = null
  serviceName = null
  order = null
  createdAt = null

  constructor({ transactionLogStore, transactionLogData }) {
    makeAutoObservable(this)

    this.transactionLogStore = transactionLogStore
    this.update(transactionLogData)
  }

  fetch = async () => {
    await this.transactionLogStore.fetch(this)
  }

  update = (transactionLog) => {
    if (transactionLog.direction) {
      this.direction = transactionLog.direction
    }

    if (transactionLog.source) {
      this.source = new Organization(transactionLog.source.serialNumber)
    }

    if (transactionLog.destination) {
      this.destination = new Organization(
        transactionLog.destination.serialNumber,
      )
    }

    if (transactionLog.service && transactionLog.service.name) {
      this.serviceName = transactionLog.service.name
    }

    if (transactionLog.order) {
      this.order = new Order({
        delegator: transactionLog.order.delegator,
        reference: transactionLog.order.reference,
      })
    }

    if (transactionLog.createdAt) {
      this.createdAt = new Date(transactionLog.createdAt)
    }
  }
}

export default TransactionLogModel
