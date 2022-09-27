// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import OutgoingOrderModel from './models/OutgoingOrderModel'
import IncomingOrderModel from './models/IncomingOrderModel'

class OrderStore {
  _isLoading = false
  _outgoingOrders = observable.map()
  _incomingOrders = observable.map()

  constructor({ rootStore, managementApiClient }) {
    makeAutoObservable(this)

    this._rootStore = rootStore
    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  get outgoingOrders() {
    return [...this._outgoingOrders.values()]
  }

  get incomingOrders() {
    return [...this._incomingOrders.values()]
  }

  create = flow(function* create(formData) {
    yield this._managementApiClient.managementCreateOutgoingOrder({
      body: {
        reference: formData.reference,
        description: formData.description,
        publicKeyPem: formData.publicKeyPem,
        delegatee: formData.delegatee,
        validFrom: formData.validFrom,
        validUntil: formData.validUntil,
        accessProofIds: formData.accessProofIds,
      },
    })
  }).bind(this)

  fetchOutgoing = flow(function* fetchOutgoing() {
    this._isLoading = true

    try {
      const result =
        yield this._managementApiClient.managementListOutgoingOrders()

      result.orders.forEach((order) => {
        order.accessProofs = order.accessProofs || []

        const accessProofModels = order.accessProofs.map((accessProof) =>
          this._rootStore.accessProofStore.updateFromServer(accessProof),
        )

        const orderModel = new OutgoingOrderModel({
          orderStore: this,
          orderData: order,
          accessProofs: accessProofModels,
        })

        this._outgoingOrders.set(
          getOutgoingKey(order.delegatee.serialNumber, order.reference),
          orderModel,
        )
      })
    } catch (error) {
      this._isLoading = false
      throw new Error(error.message)
    }

    this._isLoading = false
  }).bind(this)

  updateOutgoing = flow(function* updateOutgoing(order) {
    try {
      yield this._managementApiClient.managementUpdateOutgoingOrder({
        body: order,
      })
    } catch (err) {
      this._isLoading = false
      throw err
    }

    this._isLoading = false
  })

  getOutgoing = (delegateeSerialNumber, reference) => {
    return this._outgoingOrders.get(
      getOutgoingKey(delegateeSerialNumber, reference),
    )
  }

  revokeOutgoing = flow(function* revokeOutgoing(order) {
    yield this._managementApiClient.managementRevokeOutgoingOrder({
      delegatee: order.delegatee.serialNumber,
      reference: order.reference,
    })

    order.update({ orderData: { revokedAt: new Date() } })
  }).bind(this)

  fetchIncoming = flow(function* fetchIncoming() {
    this._isLoading = true

    try {
      const result =
        yield this._managementApiClient.managementListIncomingOrders()

      result.orders.forEach((order) => {
        const orderModel = new IncomingOrderModel({
          orderStore: this,
          orderData: order,
        })

        this._incomingOrders.set(
          getIncomingKey(order.delegator.serialNumber, order.reference),
          orderModel,
        )
      })
    } catch (error) {
      this._isLoading = false
      throw new Error(error.message)
    }

    this._isLoading = false
  }).bind(this)

  getIncoming = (delegatorSerialNumber, reference) => {
    return this._incomingOrders.get(
      getIncomingKey(delegatorSerialNumber, reference),
    )
  }

  updateIncoming = flow(function* updateIncoming() {
    this._isLoading = true

    try {
      const result =
        yield this._managementApiClient.managementSynchronizeOrders()

      result.orders.forEach((order) => {
        const orderModel = new IncomingOrderModel({
          orderStore: this,
          orderData: order,
        })

        this._incomingOrders.set(
          getIncomingKey(order.delegator.serialNumber, order.reference),
          orderModel,
        )
      })
    } catch (error) {
      this._isLoading = false
      throw error
    }

    this._isLoading = false
  }).bind(this)
}

const getOutgoingKey = (delegateeSerialNumber, reference) => {
  return `${delegateeSerialNumber}_${reference}`
}

const getIncomingKey = (delegatorSerialNumber, reference) => {
  return `${delegatorSerialNumber}_${reference}`
}

export default OrderStore
