// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const ACTION_LOGIN_SUCCESS = 'loginSuccess'
export const ACTION_LOGIN_FAIL = 'loginFail'
export const ACTION_LOGOUT_SUCCESS = 'logoutSuccess'
export const ACTION_INCOMING_ACCESS_REQUEST_ACCEPT =
  'incomingAccessRequestAccept'
export const ACTION_INCOMING_ACCESS_REQUEST_REJECT =
  'incomingAccessRequestReject'
export const ACTION_ACCESS_GRANT_REVOKE = 'access_grant_revoke'
export const ACTION_OUTGOING_ACCESS_REQUEST_CREATE =
  'outgoingAccessRequestCreate'
export const ACTION_SERVICE_CREATE = 'serviceCreate'
export const ACTION_SERVICE_UPDATE = 'serviceUpdate'
export const ACTION_SERVICE_DELETE = 'serviceDelete'
export const ACTION_ORGANIZATION_SETTINGS_UPDATE = 'organizationSettingsUpdate'
export const ACTION_INSIGHT_CONFIGURATION_UPDATE =
  'organizationInsightConfiguration_update'

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
