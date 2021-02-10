// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const ACTION_LOGIN_SUCCESS = 'login_success'
export const ACTION_LOGIN_FAIL = 'login_fail'
export const ACTION_LOGOUT_SUCCESS = 'logout_success'
export const ACTION_INCOMING_ACCESS_REQUEST_ACCEPT =
  'incoming_access_request_accept'
export const ACTION_INCOMING_ACCESS_REQUEST_REJECT =
  'incoming_access_request_reject'
export const ACTION_ACCESS_GRANT_REVOKE = 'access_grant_revoke'
export const ACTION_OUTGOING_ACCESS_REQUEST_CREATE =
  'outgoing_access_request_create'
export const ACTION_SERVICE_CREATE = 'service_create'
export const ACTION_SERVICE_UPDATE = 'service_update'
export const ACTION_SERVICE_DELETE = 'service_delete'

//            access_grant.go    │ 32     OR action_type = 'organization_settings_update'
//            access_proof.go    │ 33     OR action_type = 'organization_insight_configuration_update'

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
