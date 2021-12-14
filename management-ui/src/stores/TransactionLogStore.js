// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import TransactionLogModel from '../stores/models/TransactionLogModel'

class TransactionLogStore {
  _isLoading = false
  _transactionLogs = []

  constructor({ txLogApiClient }) {
    makeAutoObservable(this)
    this._txLogApiClient = txLogApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  get transactionLogs() {
    return this._transactionLogs
  }

  fetchAll = flow(function* fetchAll() {
    try {
      this._isLoading = true
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

      this._isLoading = false
    } catch (e) {
      throw Error(e.message)
    } finally {
      this._isLoading = false
    }
  }).bind(this)
}

export default TransactionLogStore
