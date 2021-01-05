// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext } from 'react'
import { configure } from 'mobx'
import { node, object } from 'prop-types'
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
    directoryApiClient,
    managementApiClient,
    inwayRepository = InwayRepository,
  } = {}) {
    this.applicationStore = new ApplicationStore()
    this.directoryServicesStore = new DirectoryServicesStore({
      rootStore: this,
      directoryApiClient,
    })
    this.outgoingAccessRequestStore = new OutgoingAccessRequestStore({
      rootStore: this,
      managementApiClient,
    })
    this.accessGrantStore = new AccessGrantStore({
      managementApiClient,
    })
    this.accessProofStore = new AccessProofStore()
    this.servicesStore = new ServicesStore({
      rootStore: this,
      managementApiClient,
    })
    this.incomingAccessRequestsStore = new IncomingAccessRequestsStore({
      managementApiClient,
    })
    this.inwaysStore = new InwaysStore({ rootStore: this, inwayRepository })
  }
}

export const StoreProvider = ({ children, rootStore }) => {
  return (
    <storesContext.Provider value={rootStore}>
      {children}
    </storesContext.Provider>
  )
}

StoreProvider.propTypes = {
  children: node,
  rootStore: object.isRequired,
}
