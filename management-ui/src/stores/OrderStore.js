// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import OutgoingOrderModel from './models/OutgoingOrderModel'

class OrderStore {
  _isLoading = false
  _outgoingOrders = observable.map()
  _incomingOrders = observable.map()

  constructor({ managementApiClient }) {
    makeAutoObservable(this)
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
    try {
      yield this._managementApiClient.managementCreateOutgoingOrder({
        body: formData,
      })

      const order = new OutgoingOrderModel({
        orderStore: this,
        orderData: formData,
      })

      this._outgoingOrders.set(
        getOutgoingKey(order.delegatee, order.reference),
        order,
      )

      return order
    } catch (response) {
      const err = yield response.json()
      throw new Error(err.message)
    }
  }).bind(this)

  fetchOutgoing = flow(function* fetchOutgoing() {
    this._isLoading = true

    try {
      const result =
        yield this._managementApiClient.managementListOutgoingOrders()

      result.orders.forEach((order) => {
        const orderModel = new OutgoingOrderModel({
          orderStore: this,
          orderData: order,
        })

        this._outgoingOrders.set(
          getOutgoingKey(order.delegatee, order.reference),
          orderModel,
        )
      })
    } catch (error) {
      this._isLoading = false
      throw new Error(error.message)
    }

    this._isLoading = false
  }).bind(this)

  getOutgoing = (delegatee, reference) => {
    return this._outgoingOrders.get(getOutgoingKey(delegatee, reference))
  }

  fetchIncoming = flow(function* fetchIncoming() {
    this._isLoading = true

    try {
      const result =
        yield this._managementApiClient.managementListIncomingOrders()

      result.orders.forEach((order) => {
        this._incomingOrders.set(order.reference, order)
      })
    } catch (error) {
      this._isLoading = false
      throw new Error(error.message)
    }

    this._isLoading = false
  }).bind(this)

  updateIncoming = flow(function* updateIncoming() {
    this._isLoading = true

    try {
      const result =
        yield this._managementApiClient.managementSynchronizeOrders()

      result.orders.forEach((order) => {
        this._incomingOrders.set(order.reference, order)
      })
    } catch (error) {
      this._isLoading = false
      throw new Error(error.message)
    }

    this._isLoading = false
  }).bind(this)
}

const getOutgoingKey = (delegatee, reference) => {
  return `${delegatee}_${reference}`
}

export default OrderStore
