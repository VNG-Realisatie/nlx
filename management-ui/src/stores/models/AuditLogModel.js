// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const ACTION_LOGIN_SUCCESS = 'login_success'
export const ACTION_LOGIN_FAIL = 'login_fail'
export const ACTION_LOGOUT_SUCCESS = 'logout_success'
export const ACTION_INCOMING_ACCESS_REQUEST_ACCEPT =
  'incoming_access_request_accept'

//            scheduler_test.go  │ 25     OR action_type = 'incoming_access_request_accept'
//     ▾   database/            │ 26     OR action_type = 'incoming_access_request_reject'
//       ▸   mock/              │ 27     OR action_type = 'access_grant_revoke'
//            127.0.0.1:210010401│ 28     OR action_type = 'outgoing_access_request_create'
//            127.0.0.1:210020401│ 29     OR action_type = 'service_create'
//            127.0.0.1:210250448│ 30     OR action_type = 'service_update'
//            127.0.0.1:210260448│ 31     OR action_type = 'service_delete'
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
