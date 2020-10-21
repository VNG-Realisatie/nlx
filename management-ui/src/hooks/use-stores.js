// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useContext } from 'react'
import { storesContext } from '../stores'

const useStores = () => useContext(storesContext)

export const useDirectoryServicesStore = () => {
  const { directoryServicesStore } = useStores()
  if (!directoryServicesStore.isInitiallyFetched) {
    directoryServicesStore.fetchServices()
  }
  return directoryServicesStore
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
