// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { createContext } from 'react'
import { node } from 'prop-types'
import { configure } from 'mobx'
import 'mobx-react-lite/batchingForReactDom'

import DirectoryRepository from '../domain/directory-repository'
import { createDirectoryStore } from '../pages/directory/DirectoryStore'

configure({ enforceActions: 'observed' })

export const storesContext = createContext(null)

class RootStore {
  constructor() {
    this.directoryStore = createDirectoryStore(this, DirectoryRepository)
  }
}

export const StoreProvider = ({ children }) => {
  return (
    <storesContext.Provider value={new RootStore()}>
      {children}
    </storesContext.Provider>
  )
}

StoreProvider.propTypes = {
  children: node,
}
