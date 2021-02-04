// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const AUDIT_LOG_ACTION_LOGIN = 'login'

class AuditLogModel {
  _id = null
  _action = null
  _user = null
  _createdAt = null
  _organization = null

  constructor({ auditLogData }) {
    makeAutoObservable(this)
    this.update(auditLogData)
  }

  get id() {
    return this._id
  }

  get action() {
    return this._action
  }

  get user() {
    return this._user
  }

  get createdAt() {
    return this._createdAt
  }

  get organization() {
    return this._organization
  }

  update = (auditLogData) => {
    if (!auditLogData) {
      throw Error('Data required to update audit log')
    }

    if (auditLogData.id) {
      this._id = auditLogData.id
    }

    if (auditLogData.action) {
      this._action = auditLogData.action
    }

    if (auditLogData.user) {
      this._user = auditLogData.user
    }

    if (auditLogData.createdAt) {
      this._createdAt = new Date(auditLogData.createdAt)
    }

    if (auditLogData.organization) {
      this._organization = auditLogData.organization
    }
  }
}

export default AuditLogModel
