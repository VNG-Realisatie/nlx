// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const ACTION_LOGIN_SUCCESS = 'loginSuccess'
export const ACTION_LOGIN_FAIL = 'loginFail'
export const ACTION_LOGOUT = 'logout'
export const ACTION_INCOMING_ACCESS_REQUEST_ACCEPT =
  'incomingAccessRequestAccept'
export const ACTION_INCOMING_ACCESS_REQUEST_REJECT =
  'incomingAccessRequestReject'
export const ACTION_ACCESS_GRANT_REVOKE = 'accessGrantRevoke'
export const ACTION_OUTGOING_ACCESS_REQUEST_CREATE =
  'outgoingAccessRequestCreate'
export const ACTION_OUTGOING_ACCESS_REQUEST_FAIL = 'outgoingAccessRequestFail'
export const ACTION_SERVICE_CREATE = 'serviceCreate'
export const ACTION_SERVICE_UPDATE = 'serviceUpdate'
export const ACTION_SERVICE_DELETE = 'serviceDelete'
export const ACTION_ORGANIZATION_SETTINGS_UPDATE = 'organizationSettingsUpdate'
export const ACTION_INSIGHT_CONFIGURATION_UPDATE =
  'organizationInsightConfigurationUpdate'
export const ACTION_ORDER_CREATE = 'orderCreate'
export const ACTION_ORDER_OUTGOING_REVOKE = 'orderOutgoingRevoke'
export const ACTION_INWAY_DELETE = 'inwayDelete'
export const ACTION_ORDER_OUTGOING_UPDATE = 'orderOutgoingUpdate'
export const ACTION_ACCEPT_TERMS_OF_SERVICE = 'acceptTermsOfService'

class Organization {
  serialNumber = ''
  name = ''

  constructor(serialNumber, name) {
    this.serialNumber = serialNumber
    this.name = name
  }
}

class AuditLogModel {
  id = null
  action = null
  user = null
  createdAt = null
  delegatee = null
  services = null
  operatingSystem = null
  browser = null
  client = null
  data = {
    delegatee: null,
    reference: '',
    inwayName: '',
  }

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

    if (auditLogData.services) {
      this.services = auditLogData.services
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

    if (auditLogData.data) {
      const { delegatee, reference, inwayName } = auditLogData.data

      if (delegatee) {
        this.data.delegatee = new Organization(
          delegatee.serialNumber,
          delegatee.name || delegatee.serialNumber,
        )
      }

      this.data.reference = reference
      this.data.inwayName = inwayName
    }
  }
}

export default AuditLogModel
