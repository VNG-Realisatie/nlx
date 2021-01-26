// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { flow, makeAutoObservable, observable } from 'mobx'
import AccessGrantModel from './models/AccessGrantModel'

class AccessGrantStore {
  accessGrants = observable.map()

  constructor({ managementApiClient }) {
    makeAutoObservable(this)

    this._managementApiClient = managementApiClient
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
    const result = yield this.returnForService({ name })

    // delete access grants which do not exist anymore
    const newIds = result.map((ar) => ar.id)
    this.getForService({ name }).forEach((ar) => {
      if (newIds.includes(ar.id)) {
        return
      }

      this.accessGrants.delete(ar.id)
    })

    // recreate models in-memory
    result.map((accessGrantData) => this.updateFromServer(accessGrantData))
  }).bind(this)

  returnForService = async ({ name }) => {
    const result = await this._managementApiClient.managementListAccessGrantsForService(
      {
        serviceName: name,
      },
    )

    return result.accessGrants
  }

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

  haveChangedForService = flow(function* fetch(service) {
    let latestAccessGrants = yield this.returnForService(service)

    // we are only interested in access grants which have not been revoked
    latestAccessGrants = latestAccessGrants.filter(
      (ag) => ag.revokedAt === null || ag.revokedAt === undefined,
    )

    if (latestAccessGrants.length !== service.accessGrants.length) {
      return true
    }

    // we will compare a list of sorted IDs to determine if the
    // latest access grants have changed
    const accessGrantIds = service.accessGrants.map((ag) => ag.id)
    accessGrantIds.sort()

    const latestAccessGrantIds = latestAccessGrants.map((ar) => ar.id)
    latestAccessGrantIds.sort()

    return (
      JSON.stringify(accessGrantIds) !== JSON.stringify(latestAccessGrantIds)
    )
  }).bind(this)
}

export default AccessGrantStore
