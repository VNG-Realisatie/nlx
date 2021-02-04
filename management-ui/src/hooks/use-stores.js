// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useContext } from 'react'
import { storesContext } from '../stores'

const useStores = () => useContext(storesContext)

export const useApplicationStore = () => {
  const { applicationStore } = useStores()
  return applicationStore
}

export const useDirectoryServiceStore = () => {
  const { directoryServicesStore } = useStores()

  if (!directoryServicesStore.isInitiallyFetched) {
    directoryServicesStore.fetchAll()
  }
  return directoryServicesStore
}

export const useServiceStore = () => {
  const { servicesStore } = useStores()
  if (!servicesStore.isInitiallyFetched) {
    servicesStore.fetchAll()
  }
  return servicesStore
}

export const useInwayStore = () => {
  const { inwayStore } = useStores()
  if (!inwayStore.isInitiallyFetched) {
    inwayStore.fetchInways()
  }
  return inwayStore
}

export default useStores
