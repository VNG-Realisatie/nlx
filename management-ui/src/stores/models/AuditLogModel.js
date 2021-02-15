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
  id = null
  action = null
  user = null
  createdAt = null
  organization = null
  service = null
  operatingSystem = null
  browser = null
  client = null

  constructor({ auditLogData }) {
    makeAutoObservable(this)
    this.update(auditLogData)
  }

  update = (auditLogData) => {
    if (!auditLogData) {
      throw Error('Data required to update audit log')
    }

    if (auditLogData.id) {
      this.id = auditLogData.id
    }

    if (auditLogData.action) {
      this.action = auditLogData.action
    }

    if (auditLogData.user) {
      this.user = auditLogData.user
    }

    if (auditLogData.createdAt) {
      this.createdAt = new Date(auditLogData.createdAt)
    }

    if (auditLogData.organization) {
      this.organization = auditLogData.organization
    }

    if (auditLogData.service) {
      this.service = auditLogData.service
    }

    if (auditLogData.operatingSystem) {
      this.operatingSystem = auditLogData.operatingSystem
    }

    if (auditLogData.browser) {
      this.browser = auditLogData.browser
    }

    if (auditLogData.client) {
      this.client = auditLogData.client
    }
  }
}

export default AuditLogModel
