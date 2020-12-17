// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { makeAutoObservable } from 'mobx'

class ApplicationStore {
  isOrganizationInwaySet = null

  constructor() {
    makeAutoObservable(this)
  }

  update(entries) {
    if (
      Object.prototype.hasOwnProperty.call(entries, 'isOrganizationInwaySet')
    ) {
      this.isOrganizationInwaySet = entries.isOrganizationInwaySet
    }
  }
}

export default ApplicationStore
