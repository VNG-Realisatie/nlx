// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

export const ACTION_LOGIN_SUCCESS = 'ACTION_TYPE_LOGIN_SUCCESS'
export const ACTION_LOGIN_FAIL = 'ACTION_TYPE_LOGIN_FAIL'
export const ACTION_LOGOUT = 'ACTION_TYPE_LOGOUT'
export const ACTION_INCOMING_ACCESS_REQUEST_ACCEPT =
  'ACTION_TYPE_INCOMING_ACCESS_REQUEST_ACCEPT'
export const ACTION_INCOMING_ACCESS_REQUEST_REJECT =
  'ACTION_TYPE_INCOMING_ACCESS_REQUEST_REJECT'
export const ACTION_ACCESS_GRANT_REVOKE = 'ACTION_TYPE_ACCESS_GRANT_REVOKE'
export const ACTION_OUTGOING_ACCESS_REQUEST_CREATE =
  'ACTION_TYPE_OUTGOING_ACCESS_REQUEST_CREATE'
export const ACTION_OUTGOING_ACCESS_REQUEST_FAIL =
  'ACTION_TYPE_OUTGOING_ACCESS_REQUEST_FAIL'
export const ACTION_OUTGOING_ACCESS_REQUEST_WITHDRAW =
  'ACTION_TYPE_OUTGOING_ACCESS_REQUEST_WITHDRAW'
export const ACTION_ACCESS_TERMINATE = 'ACTION_TYPE_ACCESS_TERMINATE'
export const ACTION_SERVICE_CREATE = 'ACTION_TYPE_SERVICE_CREATE'
export const ACTION_SERVICE_UPDATE = 'ACTION_TYPE_SERVICE_UPDATE'
export const ACTION_SERVICE_DELETE = 'ACTION_TYPE_SERVICE_DELETE'
export const ACTION_ORGANIZATION_SETTINGS_UPDATE =
  'ACTION_TYPE_ORGANIZATION_SETTINGS_UPDATE'
export const ACTION_ORDER_CREATE = 'ACTION_TYPE_ORDER_CREATE'
export const ACTION_ORDER_OUTGOING_REVOKE = 'ACTION_TYPE_ORDER_OUTGOING_REVOKE'
export const ACTION_ORDER_INCOMING_REVOKE = 'ACTION_TYPE_ORDER_INCOMING_REVOKE'
export const ACTION_INWAY_DELETE = 'ACTION_TYPE_INWAY_DELETE'
export const ACTION_OUTWAY_DELETE = 'ACTION_TYPE_OUTWAY_DELETE'
export const ACTION_ORDER_OUTGOING_UPDATE = 'ACTION_TYPE_ORDER_OUTGOING_UPDATE'
export const ACTION_ACCEPT_TERMS_OF_SERVICE =
  'ACTION_TYPE_ACCEPT_TERMS_OF_SERVICE'

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
    outwayName: '',
    publicKeyFingerprint: '',
  }
  hasSucceeded = null

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

    if (typeof auditLogData.hasSucceeded != 'undefined') {
      this.hasSucceeded = auditLogData.hasSucceeded
    }

    if (auditLogData.data) {
      const {
        delegatee,
        reference,
        inwayName,
        outwayName,
        publicKeyFingerprint,
      } = auditLogData.data

      if (delegatee) {
        this.data.delegatee = new Organization(
          delegatee.serialNumber,
          delegatee.name || delegatee.serialNumber,
        )
      }

      this.data.reference = reference
      this.data.inwayName = inwayName
      this.data.outwayName = outwayName
      this.data.publicKeyFingerprint = publicKeyFingerprint
    }
  }
}

export default AuditLogModel
