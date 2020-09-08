// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext } from 'react'
import { node, object } from 'prop-types'
import { configure } from 'mobx'

import { createDirectoryStore } from '../pages/directory/DirectoryStore'
import { createServicesStore } from '../pages/services/ServicesStore'

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
