// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { makeAutoObservable, observable } from 'mobx'
import OutgoingAccessRequestSyncErrorModel, {
  SYNC_ERROR_INTERNAL,
  SYNC_ERROR_SERVICE_PROVIDER_NO_ORGANIZATION_INWAY_SPECIFIED,
  SYNC_ERROR_SERVICE_PROVIDER_ORGANIZATION_INWAY_UNREACHABLE,
} from './models/OutgoingAccessRequestSyncErrorModel'

class OutgoingAccessRequestSyncErrorStore {
  _errorsByOrganization = observable.map()

  constructor() {
    makeAutoObservable(this)
  }

  loadFromSyncResponse(organizationSerialNumber, serviceName, jsonResponse) {
    if (!jsonResponse) {
      throw new Error('please provide the JSON response')
    }

    const syncError = new OutgoingAccessRequestSyncErrorModel({
      syncErrorData: {
        error: convertMessageToError(jsonResponse.message),
        organizationSerialNumber: organizationSerialNumber,
        serviceName: serviceName,
      },
    })

    let errorsForOrganization = this._errorsByOrganization.get(
      organizationSerialNumber,
    )

    if (typeof errorsForOrganization === 'undefined') {
      const newMap = observable.map()
      this._errorsByOrganization.set(organizationSerialNumber, newMap)
      errorsForOrganization = newMap
    }

    errorsForOrganization.set(serviceName, syncError)
  }

  loadFromSyncAllResponse(jsonResponse) {
    if (!jsonResponse) {
      throw new Error('please provide the JSON response')
    }

    if (
      !jsonResponse.details ||
      !jsonResponse.details[0] ||
      !jsonResponse.details[0].metadata
    ) {
      return
    }

    Object.keys(jsonResponse.details[0].metadata).forEach(
      (organizationSerialNumber) => {
        const message =
          jsonResponse.details[0].metadata[`${organizationSerialNumber}`]

        const syncError = new OutgoingAccessRequestSyncErrorModel({
          syncErrorData: {
            error: convertMessageToError(message),
            organizationSerialNumber: organizationSerialNumber,
          },
        })

        let errorsForOrganization = this._errorsByOrganization.get(
          organizationSerialNumber,
        )

        if (typeof errorsForOrganization === 'undefined') {
          const newMap = observable.map()
          this._errorsByOrganization.set(organizationSerialNumber, newMap)
          errorsForOrganization = newMap
        }

        errorsForOrganization.set('', syncError)
      },
    )
  }

  clearAll() {
    this._errorsByOrganization.clear()
  }

  clearForService(organizationSerialNumber, serviceName) {
    const errorsByOrg = this._errorsByOrganization.get(organizationSerialNumber)

    if (!errorsByOrg) {
      return
    }

    errorsByOrg.delete(serviceName)
  }

  getForService = (organizationSerialNumber, serviceName) => {
    let errorsForOrganization = this._errorsByOrganization.get(
      organizationSerialNumber,
    )

    if (typeof errorsForOrganization === 'undefined') {
      this._errorsByOrganization.set(organizationSerialNumber, observable.map())

      errorsForOrganization = this._errorsByOrganization.get(
        organizationSerialNumber,
      )
    }

    const errorForService = errorsForOrganization.get(serviceName)

    if (errorForService) {
      return errorForService
    }

    const organizationWideError = errorsForOrganization.get('')

    if (organizationWideError) {
      return organizationWideError
    }
  }
}

const convertMessageToError = (message) => {
  switch (message) {
    case 'service_provider_no_organization_inway_specified':
      return SYNC_ERROR_SERVICE_PROVIDER_NO_ORGANIZATION_INWAY_SPECIFIED

    case 'service_provider_organization_inway_unreachable':
      return SYNC_ERROR_SERVICE_PROVIDER_ORGANIZATION_INWAY_UNREACHABLE

    case 'internal_error':
      return SYNC_ERROR_INTERNAL

    default:
      throw new Error(`unable to convert message '${message}' to error`)
  }
}

export default OutgoingAccessRequestSyncErrorStore
