// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable, observable, flow } from 'mobx'

import AccessGrantModel from '../models/AccessGrantModel'

class AccessGrantStore {
  accessGrants = observable.map()

  constructor({ accessGrantRepository }) {
    makeAutoObservable(this)

    this.accessGrantRepository = accessGrantRepository
  }

  updateFromServer = (accessGrantData) => {
    if (!accessGrantData) return null

    const cachedAccessGrant = this.accessGrants.get(accessGrantData.id)

    if (cachedAccessGrant) {
      cachedAccessGrant.update(accessGrantData)
      return cachedAccessGrant
    }

    const accessGrant = new AccessGrantModel({ accessGrantData })

    this.accessGrants.set(accessGrant.id, accessGrant)

    return accessGrant
  }

  fetchForService = flow(function* fetchForService(serviceModel) {
    const response = yield this.accessGrantRepository.fetchByServiceName(
      serviceModel.name,
    )

    response.map((accessGrantData) => this.updateFromServer(accessGrantData))
  }).bind(this)

  getForService = (serviceModel) => {
    const arrayOfModels = [...this.accessGrants.values()]

    return arrayOfModels.filter(
      (accessGrantModel) => accessGrantModel.serviceName === serviceModel.name,
    )
  }
}

export default AccessGrantStore
