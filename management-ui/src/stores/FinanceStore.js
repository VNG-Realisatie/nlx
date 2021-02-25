// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable, flow } from 'mobx'

class FinanceStore {
  _isLoading = false
  enabled = null
  isInitiallyFetched = false

  constructor({ managementApiClient }) {
    makeAutoObservable(this)
    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  fetch = flow(function* fetch() {
    try {
      this._isLoading = true
      const result = yield this._managementApiClient.managementIsBillingEnabled()

      this.enabled = result.enabled
      this._isLoading = false
    } catch (err) {
      this._isLoading = false
      throw new Error(err.message)
    } finally {
      this._isLoading = false
      this.isInitiallyFetched = true
    }
  }).bind(this)

  async downloadExport() {
    try {
      const result = await this._managementApiClient.managementDownloadBillingExport()

      return result
    } catch (err) {
      throw new Error(err.message)
    }
  }
}

export default FinanceStore
