// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'

class OrderStore {
  _isLoading = false
  _orders = observable.map()

  constructor({ managementApiClient }) {
    makeAutoObservable(this)

    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  get orders() {
    return [...this._orders.values()]
  }

  fetchAll = flow(function* fetchAll() {
    this._isLoading = true

    try {
      const result =
        yield this._managementApiClient.managementListIssuedOrders()

      result.orders.forEach((order) => {
        this._orders.set(order.reference, order)
      })
    } catch (error) {
      this._isLoading = false
      throw new Error(error.message)
    }

    this._isLoading = false
  }).bind(this)

  create = flow(function* create(formData) {
    try {
      const orderData = yield this._managementApiClient.managementCreateOrder({
        body: formData,
      })

      return orderData.id
    } catch (err) {
      throw new Error(err.message)
    }
  }).bind(this)
}

export default OrderStore
