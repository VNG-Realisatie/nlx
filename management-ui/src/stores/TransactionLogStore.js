// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import TransactionLogModel from '../stores/models/TransactionLogModel'

class TransactionLogStore {
  _isLoading = false
  _enabled = null
  _transactionLogs = []
  _isInitiallyFetched = false

  constructor({ txLogApiClient, managementApiClient }) {
    makeAutoObservable(this)
    this._txLogApiClient = txLogApiClient
    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  get isEnabled() {
    return this._enabled
  }

  get transactionLogs() {
    return this._transactionLogs
  }

  get isInitiallyFetched() {
    return this._isInitiallyFetched
  }

  fetchAll = flow(function* fetchAll() {
    try {
      this._isLoading = true
      if (this._enabled === null) {
        const result = yield this._managementApiClient.managementIsTXLogEnabled(
          {},
        )
        this._enabled = result.enabled
      }

      if (this._enabled) {
        const transactionLogsData = yield this._txLogApiClient.tXLogListRecords(
          {},
        )
        this._transactionLogs = transactionLogsData.records.map(
          (transactionLogData) =>
            new TransactionLogModel({
              transactionLogStore: this,
              transactionLogData: transactionLogData,
            }),
        )
      }

      this._isLoading = false
    } catch (e) {
      throw Error(e.message)
    } finally {
      this._isLoading = false
      this._isInitiallyFetched = true
    }
  }).bind(this)
}

export default TransactionLogStore
