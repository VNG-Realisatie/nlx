// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useContext } from 'react'
import { storesContext } from '../stores'

export const useStores = () => useContext(storesContext)

export const useApplicationStore = () => {
  const { applicationStore } = useStores()
  return applicationStore
}

export const useDirectoryServiceStore = () => {
  const { directoryServicesStore } = useStores()
  return directoryServicesStore
}

export const useFinanceStore = () => {
  const { financeStore } = useStores()
  if (!financeStore.isInitiallyFetched) {
    financeStore.fetch()
  }
  return financeStore
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

export const useTransactionLogStore = () => {
  const { transactionLogStore } = useStores()
  if (!transactionLogStore.isInitiallyFetched) {
    transactionLogStore.fetchAll()
  }
  return transactionLogStore
}

export const useOutwayStore = () => {
  const { outwayStore } = useStores()
  return outwayStore
}

export const useOrderStore = () => {
  const { orderStore } = useStores()
  return orderStore
}

export const useAccessProofStore = () => {
  const { accessProofStore } = useStores()
  return accessProofStore
}

export const useOutgoingAccessRequestStore = () => {
  const { outgoingAccessRequestStore } = useStores()
  return outgoingAccessRequestStore
}

export default useStores
