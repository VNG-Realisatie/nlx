// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable } from 'mobx'

export const AUTH_OIDC = 'auth-oidc'
export const AUTH_BASIC_AUTH = 'auth-basic-auth'

class ApplicationStore {
  authStrategy = AUTH_OIDC
  isOrganizationInwaySet = null
  error = ''

  constructor({ managementApiClient }) {
    makeAutoObservable(this)

    this.error = ''
    this._managementApiClient = managementApiClient
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

  getGeneralSettings = flow(function* getGeneralSettings() {
    try {
      const response = yield this._managementApiClient.managementGetSettings()
      return response
    } catch (e) {
      this.error = e
      throw new Error(e)
    }
  }).bind(this)

  async updateGeneralSettings(settings) {
    try {
      const response = await this._managementApiClient.managementUpdateSettings(
        {
          body: settings,
        },
      )
      return response
    } catch (e) {
      this.error = e
      throw new Error(e)
    }
  }
}

export default ApplicationStore
