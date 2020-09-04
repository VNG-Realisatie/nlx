// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useContext } from 'react'
import { storesContext } from '../stores'

// Tested in `src/stores/index.test.js`
const useStores = () => useContext(storesContext)

export const useDirectoryStore = () => {
  const { directoryStore } = useStores()
  if (!directoryStore.isInitiallyFetched) {
    directoryStore.fetchServices()
  }
  return directoryStore
}

export const useServicesStore = () => {
  const { servicesStore } = useStores()
  if (!servicesStore.isInitiallyFetched) {
    servicesStore.fetchServices()
  }
  return servicesStore
}

export const useInwaysStore = () => {
  const { inwaysStore } = useStores()
  if (!inwaysStore.isInitiallyFetched) {
    inwaysStore.fetchInways()
  }
  return inwaysStore
}

export default useStores
