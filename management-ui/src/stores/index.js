// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext } from 'react'
import { configure } from 'mobx'
import { node, object } from 'prop-types'
import AccessRequestRepository from '../domain/access-request-repository'
import DirectoryRepository from '../domain/directory-repository'
import InwaysStore from './InwaysStore'
import ServicesStore from './ServicesStore'
import DirectoryStore from './DirectoryStore'
import OutgoingAccessRequestStore from './OutgoingAccessRequestStore'

if (process.env.NODE_ENV !== 'test') {
  // `setupTests` has 'never' set. But some tests include this file,
  // so make sure we don't override the test setup.
  configure({ enforceActions: 'observed' })
}

export const storesContext = createContext(null)

export class RootStore {
  constructor({
    accessRequestRepository = AccessRequestRepository,
    directoryRepository = DirectoryRepository,
  } = {}) {
    this.outgoingAccessRequestsStore = new OutgoingAccessRequestStore({
      rootStore: this,
      accessRequestRepository,
    })
    this.directoryStore = new DirectoryStore({
      rootStore: this,
      directoryRepository,
    })
    this.servicesStore = new ServicesStore({ rootStore: this })
    this.inwaysStore = new InwaysStore({ rootStore: this })
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
