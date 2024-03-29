// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'

export const AUTH_OIDC = 'auth-oidc'
export const AUTH_BASIC_AUTH = 'auth-basic-auth'

class ApplicationStore {
  authStrategy = AUTH_OIDC
  isOrganizationInwaySet = null
  isOrganizationEmailAddressSet = null
  error = ''

  constructor({ managementApiClient, directoryApiClient }) {
    makeAutoObservable(this)

    this.error = ''
    this._managementApiClient = managementApiClient
    this._directoryApiClient = directoryApiClient
  }

  setBasicAuthStrategy() {
    this.authStrategy = AUTH_BASIC_AUTH
  }

  updateOrganizationInway(entries) {
    if (
      Object.prototype.hasOwnProperty.call(entries, 'isOrganizationInwaySet')
    ) {
      this.isOrganizationInwaySet = !!entries.isOrganizationInwaySet
    }
  }

  updateOrganizationEmailAddress(entries) {
    if (
      Object.prototype.hasOwnProperty.call(
        entries,
        'isOrganizationEmailAddressSet',
      )
    ) {
      this.isOrganizationEmailAddressSet =
        !!entries.isOrganizationEmailAddressSet
    }
  }

  getGeneralSettings = flow(function* getGeneralSettings() {
    try {
      const response =
        yield this._managementApiClient.managementServiceGetSettings()
      return response.settings
    } catch (e) {
      this.error = e
      throw new Error(e)
    }
  }).bind(this)

  async getTermsOfService() {
    return this._directoryApiClient.directoryServiceGetTermsOfService()
  }

  async getTermsOfServiceStatus() {
    return this._managementApiClient.managementServiceGetTermsOfServiceStatus()
  }

  async acceptTermsOfService() {
    return this._managementApiClient.managementServiceAcceptTermsOfService()
  }

  async updateGeneralSettings(settings) {
    return await this._managementApiClient.managementServiceUpdateSettings({
      body: settings,
    })
  }
}

export default ApplicationStore
