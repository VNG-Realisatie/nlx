// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import AuditLogModel from './models/AuditLogModel'

class AuditLogStore {
  _isLoading = false
  _auditLogs = observable.map()

  constructor({ managementApiClient }) {
    makeAutoObservable(this)
    this._managementApiClient = managementApiClient
  }

  get isLoading() {
    return this._isLoading
  }

  get auditLogs() {
    return [...this._auditLogs.values()]
  }

  fetchAll = flow(function* fetchAll() {
    try {
      this._isLoading = true
      const result = yield this._managementApiClient.managementListAuditLogs()
      const auditLogsData = result.auditLogs

      // delete audit logs which do not exist anymore
      const newIds = auditLogsData.map((al) => al.id)
      this.auditLogs.forEach((al) => {
        if (newIds.includes(al.id)) {
          return
        }

        this._auditLogs.delete(al.id)
      })

      // recreate models in-memory
      auditLogsData.map((auditLogData) => this.updateFromServer(auditLogData))

      this._isLoading = false
    } catch (err) {
      this._isLoading = false
      throw new Error(err.message)
    }
  }).bind(this)

  updateFromServer = (auditLogData) => {
    if (!auditLogData) return null

    const cachedAuditLog = this._auditLogs.get(auditLogData.id)

    if (cachedAuditLog) {
      cachedAuditLog.update(auditLogData)
      return cachedAuditLog
    }

    const auditLog = new AuditLogModel({ auditLogData })
    this._auditLogs.set(auditLog.id, auditLog)

    return auditLog
  }
}

export default AuditLogStore
