// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext } from 'react'
import { configure } from 'mobx'
import { node, object } from 'prop-types'
import { createServicesStore } from '../pages/services/ServicesStore'
import { createInwaysStore } from '../pages/inways/InwaysStore'
import { createDirectoryStore } from './DirectoryStore'

if (process.env.NODE_ENV !== 'test') {
  // `setupTests` has 'never' set. But some tests include this file,
  // so make sure we don't override the test setup.
  configure({ enforceActions: 'observed' })
}

export const storesContext = createContext(null)

class RootStore {
  constructor() {
    this.directoryStore = createDirectoryStore({ rootStore: this })
    this.servicesStore = createServicesStore({ rootStore: this })
    this.inwaysStore = createInwaysStore({ rootStore: this })
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
