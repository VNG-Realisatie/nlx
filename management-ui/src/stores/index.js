// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext } from 'react'
import { configure } from 'mobx'
import { node, object } from 'prop-types'

import AccessRequestRepository from '../domain/access-request-repository'
import AccessGrantRepository from '../domain/access-grant-repository'
import DirectoryRepository from '../domain/directory-repository'
import InwayRepository from '../domain/inway-repository'
import ApplicationStore from './ApplicationStore'
import DirectoryServicesStore from './DirectoryServicesStore'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'
import AccessGrantStore from './AccessGrantStore'
import AccessProofStore from './AccessProofStore'
import ServicesStore from './ServicesStore'
import IncomingAccessRequestsStore from './IncomingAccessRequestsStore'
import InwaysStore from './InwaysStore'

if (process.env.NODE_ENV !== 'test') {
  // `setupTests` has 'never' set. But some tests include this file,
  // so make sure we don't override the test setup.
  configure({ enforceActions: 'observed' })
}

export const storesContext = createContext(null)

export class RootStore {
  constructor({
    accessRequestRepository = AccessRequestRepository,
    accessGrantRepository = AccessGrantRepository,
    directoryRepository = DirectoryRepository,
    inwayRepository = InwayRepository,
  } = {}) {
    this.applicationStore = new ApplicationStore({ rootStore: this })
    this.directoryServicesStore = new DirectoryServicesStore({
      rootStore: this,
      directoryRepository,
    })
    this.outgoingAccessRequestStore = new OutgoingAccessRequestStore({
      rootStore: this,
      accessRequestRepository,
    })
    this.accessGrantStore = new AccessGrantStore({ accessGrantRepository })
    this.accessProofStore = new AccessProofStore()
    this.servicesStore = new ServicesStore({ rootStore: this })
    this.incomingAccessRequestsStore = new IncomingAccessRequestsStore({
      accessRequestRepository,
    })
    this.inwaysStore = new InwaysStore({ rootStore: this, inwayRepository })
  }
}

export const StoreProvider = ({ children, store = new RootStore() }) => {
  return (
    <storesContext.Provider value={store}>{children}</storesContext.Provider>
  )
}

StoreProvider.propTypes = {
  children: node,
  store: object,
}
