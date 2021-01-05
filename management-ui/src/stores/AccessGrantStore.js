// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable, flow } from 'mobx'

import AccessGrantModel from '../models/AccessGrantModel'

class AccessGrantStore {
  accessGrants = observable.map()

  constructor({ managementApiClient, accessGrantRepository }) {
    makeAutoObservable(this)

    this._managementApiClient = managementApiClient
    this.accessGrantRepository = accessGrantRepository
  }

  updateFromServer = (accessGrantData) => {
    if (!accessGrantData) return null

    const cachedAccessGrant = this.accessGrants.get(accessGrantData.id)

    if (cachedAccessGrant) {
      cachedAccessGrant.update(accessGrantData)
      return cachedAccessGrant
    }

    const accessGrant = new AccessGrantModel({
      accessGrantStore: this,
      accessGrantData,
    })

    this.accessGrants.set(accessGrant.id, accessGrant)

    return accessGrant
  }

  fetchForService = flow(function* fetchForService({ name }) {
    const response = yield this._managementApiClient.managementListAccessGrantsForService(
      {
        serviceName: name,
      },
    )

    response.accessGrants.map((accessGrantData) =>
      this.updateFromServer(accessGrantData),
    )
  }).bind(this)

  getForService = ({ name }) => {
    const arrayOfModels = [...this.accessGrants.values()]

    return arrayOfModels.filter(
      (accessGrantModel) => accessGrantModel.serviceName === name,
    )
  }

  revokeAccessGrant = async ({ organizationName, serviceName, id }) => {
    await this._managementApiClient.managementRevokeAccessGrant({
      serviceName,
      organizationName,
      accessGrantID: id,
    })
    this.fetchForService({ name: serviceName })
  }
}

export default AccessGrantStore
