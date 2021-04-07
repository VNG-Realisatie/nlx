// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'

class OrderStore {
  constructor({ managementApiClient }) {
    makeAutoObservable(this)

    this._managementApiClient = managementApiClient
  }

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
