// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const DIRECTION_IN = 'TX_LOG_DIRECTION_IN'
export const DIRECTION_OUT = 'TX_LOG_DIRECTION_OUT'

class Organization {
  serialNumber = ''
  name = ''

  constructor(serialNumber, name) {
    this.serialNumber = serialNumber
    this.name = name
  }
}

class Order {
  delegator = null
  reference = ''

  constructor({ delegator, reference }) {
    this.reference = reference
    this.delegator = new Organization(
      delegator.serialNumber,
      delegator.name || delegator.serialNumber,
    )
  }
}

class TransactionLogModel {
  transactionId = null
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
    if (transactionLog.transactionId) {
      this.transactionId = transactionLog.transactionId
    }

    if (transactionLog.direction) {
      this.direction = transactionLog.direction
    }

    if (transactionLog.source) {
      this.source = new Organization(
        transactionLog.source.serialNumber,
        transactionLog.source.name || transactionLog.source.serialNumber,
      )
    }

    if (transactionLog.destination) {
      this.destination = new Organization(
        transactionLog.destination.serialNumber,
        transactionLog.destination.name ||
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
