// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

class FinanceStore {
  constructor({ managementApiClient }) {
    makeAutoObservable(this)
    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

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
