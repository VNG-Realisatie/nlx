// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'
export const SYNC_ERROR_INTERNAL = 'internal_error'
export const SYNC_ERROR_SERVICE_PROVIDER_NO_ORGANIZATION_INWAY_SPECIFIED =
  'service_provider_no_organization_inway_specified'
export const SYNC_ERROR_SERVICE_PROVIDER_ORGANIZATION_INWAY_UNREACHABLE =
  'service_provider_organization_inway_unreachable'

const allowedSyncErrors = [
  SYNC_ERROR_INTERNAL,
  SYNC_ERROR_SERVICE_PROVIDER_NO_ORGANIZATION_INWAY_SPECIFIED,
  SYNC_ERROR_SERVICE_PROVIDER_ORGANIZATION_INWAY_UNREACHABLE,
]

class OutgoingAccessRequestSyncErrorModel {
  _error = null
  _organizationSerialNumber = null
  _serviceName = null

  constructor({ syncErrorData }) {
    makeAutoObservable(this)

    this.update(syncErrorData)
  }

  get error() {
    return this._error
  }

  get message() {
    switch (this._error) {
      case SYNC_ERROR_INTERNAL:
        return 'Internal error while trying to retrieve the current state of your access request. Please consult your system administrator.'

      case SYNC_ERROR_SERVICE_PROVIDER_NO_ORGANIZATION_INWAY_SPECIFIED:
        return 'The organization has not specified an organization Inway. We are unable to retrieve the current state of your access requests.'

      case SYNC_ERROR_SERVICE_PROVIDER_ORGANIZATION_INWAY_UNREACHABLE:
        return 'The organization Inway of this organization is unreachable. We are unable to retrieve the current state of your access requests.'

      default:
        return `unknown error: ${this._error}`
    }
  }

  update = (syncErrorData) => {
    if (!syncErrorData) {
      throw Error('Data required to update sync error')
    }

    if (!syncErrorData.organizationSerialNumber) {
      throw new Error('please provide an organization serial number')
    }

    if (!allowedSyncErrors.includes(syncErrorData.error)) {
      throw new Error(
        `unknown sync error '${
          syncErrorData.error
        }', must be one of [${allowedSyncErrors.join(', ')}]`,
      )
    }

    this._organizationSerialNumber = syncErrorData.organizationSerialNumber
    this._serviceName = syncErrorData.serviceName
    this._error = syncErrorData.error
  }
}

export default OutgoingAccessRequestSyncErrorModel
