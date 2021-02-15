// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'
import AuditLogModel from './models/AuditLogModel'

class AuditLogStore {
  _isLoading = false
  _auditLogs = []

  constructor({ managementApiClient }) {
    makeAutoObservable(this)
    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  get auditLogs() {
    return this._auditLogs
  }

  fetchAll = flow(function* fetchAll() {
    try {
      this._isLoading = true
      const result = yield this._managementApiClient.managementListAuditLogs()
      const auditLogsData = result.auditLogs

      // delete audit logs which do not exist anymore
      const newIds = auditLogsData.map((al) => al.id)
      this._auditLogs.forEach((al, index) => {
        if (newIds.includes(al.id)) {
          return
        }

        this._auditLogs.splice(index, 1)
      })

      // recreate models in-memory
      this.updateFromServer(auditLogsData)

      this._isLoading = false
    } catch (err) {
      this._isLoading = false
      throw new Error(err.message)
    }
  }).bind(this)

  _getById(id) {
    return this.auditLogs.find((auditLog) => auditLog.id === id)
  }

  updateFromServer = (auditLogsData) => {
    if (!auditLogsData) return null

    auditLogsData.forEach((auditLogData) => {
      const cachedAuditLog = this._getById(auditLogData.id)

      if (cachedAuditLog) {
        cachedAuditLog.update(auditLogData)
        return cachedAuditLog
      }

      const auditLog = new AuditLogModel({ auditLogData })
      this._auditLogs.push(auditLog)
    })

    // sort items according to order from server
    const sortedIds = auditLogsData.map((auditLog) => auditLog.id)
    this._auditLogs.sort((firstEl, secondEl) =>
      sortedIds.indexOf(firstEl.id) > sortedIds.indexOf(secondEl.id) ? 1 : -1,
    )
  }
}

export default AuditLogStore
