// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useContext } from 'react'
import { storesContext } from '../stores'

const useStores = () => useContext(storesContext)

export const useDirectoryStore = () => {
  const { directoryStore } = useStores()
  return directoryStore
}

export default useStores
