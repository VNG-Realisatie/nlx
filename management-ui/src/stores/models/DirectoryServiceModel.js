// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import OutgoingAccessRequestModel, {
  ACCESS_REQUEST_STATES,
} from './OutgoingAccessRequestModel'
import AccessProofModel from './AccessProofModel'

function throwErrorWhenNotInstanceOf(object, model) {
  if (object && !(object instanceof model)) {
    throw new Error(`Object should be an instance of ${model}`)
  }
}

class DirectoryServiceModel {
  organization = {
    name: '',
    serialNumber: '',
  }

  serviceName = ''
  state = ''
  apiSpecificationType = ''
  documentationURL = ''
  publicSupportContact = ''
  oneTimeCosts = 0
  monthlyCosts = 0
  requestCosts = 0
  _accessStates = observable.map()

  constructor({ directoryServicesStore, serviceData, accessStates }) {
    makeAutoObservable(this)

    this.directoryServicesStore = directoryServicesStore

    this.update({ serviceData, accessStates })
  }

  // TODO: add test
  getAccessStateFor(publicKeyFingerprint) {
    return this._accessStates.get(publicKeyFingerprint) || {}
  }

  update({ serviceData, accessStates }) {
    if (serviceData.serviceName) {
      this.serviceName = serviceData.serviceName
    }

    if (serviceData.state) {
      this.state = serviceData.state
    }

    if (serviceData.apiSpecificationType) {
      this.apiSpecificationType = serviceData.apiSpecificationType
    }

    if (serviceData.documentationURL) {
      this.documentationURL = serviceData.documentationURL
    }

    if (serviceData.publicSupportContact) {
      this.publicSupportContact = serviceData.publicSupportContact
    }

    if (serviceData.oneTimeCosts) {
      this.oneTimeCosts = serviceData.oneTimeCosts / 100
    }

    if (serviceData.monthlyCosts) {
      this.monthlyCosts = serviceData.monthlyCosts / 100
    }

    if (serviceData.requestCosts) {
      this.requestCosts = serviceData.requestCosts / 100
    }

    if (serviceData.organization) {
      this.organization.name = serviceData.organization.name
      this.organization.serialNumber = serviceData.organization.serialNumber
    }

    if (accessStates) {
      if (!Array.isArray(accessStates)) {
        throw new Error('invalid accessStates provided. expected array.')
      }

      this._accessStates.clear()

      accessStates.forEach((accessState) => {
        const accessRequest = accessState.accessRequest
        const accessProof = accessState.accessProof

        throwErrorWhenNotInstanceOf(accessRequest, OutgoingAccessRequestModel)
        throwErrorWhenNotInstanceOf(accessProof, AccessProofModel)

        this._accessStates.set(accessRequest.publicKeyFingerprint, {
          accessRequest,
          accessProof,
        })
      })
    }

    return this
  }

  fetch = async () => {
    await this.directoryServicesStore.fetch(
      this.organization.serialNumber,
      this.serviceName,
    )
  }

  requestAccess = flow(function* requestAccess(publicKeyFingerprint) {
    yield this.directoryServicesStore.requestAccess(
      this.organization.serialNumber,
      this.serviceName,
      publicKeyFingerprint,
    )

    yield this.fetch()
  }).bind(this)

  retryRequestAccess = flow(function* retryRequestAccess(publicKeyFingerprint) {
    const accessStateForFingerprint =
      this._accessStates.get(publicKeyFingerprint)

    if (
      !accessStateForFingerprint ||
      !accessStateForFingerprint.accessRequest
    ) {
      return false
    }

    yield accessStateForFingerprint.accessRequest.retry()
    yield this.fetch()
  }).bind(this)

  hasAccess(publicKeyFingerprint) {
    const accessStateForFingerprint =
      this._accessStates.get(publicKeyFingerprint)

    if (!accessStateForFingerprint) {
      return false
    }

    const accessRequest = accessStateForFingerprint.accessRequest
    const accessProof = accessStateForFingerprint.accessProof

    return !!(
      accessRequest &&
      accessRequest.state === ACCESS_REQUEST_STATES.APPROVED &&
      accessProof &&
      !accessProof.revokedAt
    )
  }
}

export default DirectoryServiceModel
