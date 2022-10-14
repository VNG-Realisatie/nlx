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
  documentationUrl = ''
  publicSupportContact = ''
  oneTimeCosts = 0
  monthlyCosts = 0
  requestCosts = 0
  _accessStates = observable.map()

  constructor({ directoryServicesStore, serviceData, accessStates }) {
    makeAutoObservable(this)

    this.directoryServiceServicesStore = directoryServicesStore

    this.update({ serviceData, accessStates })
  }

  getAccessStateFor(publicKeyFingerprint) {
    return this._accessStates.get(publicKeyFingerprint) || {}
  }

  getFailingAccessStates() {
    return [...this._accessStates.values()].filter((accessState) => {
      return accessState.accessRequest.state === ACCESS_REQUEST_STATES.FAILED
    })
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

    if (serviceData.documentationUrl) {
      this.documentationUrl = serviceData.documentationUrl
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
      this.organization.serialNumber = serviceData.organization.serialNumber
      this.organization.name =
        serviceData.organization.name || serviceData.organization.serialNumber
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
    await this.directoryServiceServicesStore.fetch(
      this.organization.serialNumber,
      this.serviceName,
    )
  }

  syncOutgoingAccessRequests = flow(function* syncOutgoingAccessRequests() {
    yield this.directoryServiceServicesStore.syncOutgoingAccessRequests(
      this.organization.serialNumber,
      this.serviceName,
    )
  }).bind(this)

  requestAccess = flow(function* requestAccess(publicKeyPem) {
    yield this.directoryServiceServicesStore.requestAccess(
      this.organization.serialNumber,
      this.serviceName,
      publicKeyPem,
    )

    yield this.fetch()
  }).bind(this)

  _accessStateHasAccess(accessState) {
    const { accessRequest, accessProof } = accessState
    return !!(
      accessRequest.state === ACCESS_REQUEST_STATES.APPROVED &&
      accessProof &&
      !accessProof.revokedAt
    )
  }

  hasAccess(publicKeyFingerprint) {
    const accessStateForFingerprint =
      this._accessStates.get(publicKeyFingerprint)

    if (!accessStateForFingerprint) {
      return false
    }

    return this._accessStateHasAccess(accessStateForFingerprint)
  }

  get accessStatesWithAccess() {
    return [...this._accessStates.values()].filter((accessState) => {
      return this._accessStateHasAccess(accessState)
    })
  }
}

export default DirectoryServiceModel
