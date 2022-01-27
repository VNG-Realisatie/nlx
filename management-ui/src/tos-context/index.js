// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState, useEffect, createContext, useRef } from 'react'
import { string, node, shape, bool } from 'prop-types'
import useStores from '../hooks/use-stores'

const ToSContext = createContext()

const ToSContextProvider = ({ children, tos: defaultTos }) => {
  const [tos, setTos] = useState(defaultTos || null)
  const [isReady, setReady] = useState(false)
  const stores = useStores()

  const componentIsMounted = useRef(true)

  useEffect(
    () => {
      const fetchTos = async () => {
        try {
          const tos = await stores.applicationStore.getTermsOfService()

          if (componentIsMounted) {
            setTos(tos)
          }
        } catch (err) {
          setTos(null)
        }

        if (componentIsMounted) {
          setReady(true)
        }
      }

      if (defaultTos) {
        setReady(true)
        return
      }

      fetchTos()

      return () => {
        componentIsMounted.current = false
      }
    },
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [],
  )

  return (
    <ToSContext.Provider
      value={{
        tos: tos,
        isReady: isReady,
      }}
    >
      {children}
    </ToSContext.Provider>
  )
}

ToSContextProvider.propTypes = {
  children: node,
  tos: shape({
    url: string,
    enabled: bool,
  }),
}

export default ToSContext
export { ToSContextProvider }
