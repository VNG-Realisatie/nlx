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
        let tos = {}
        try {
          tos = await stores.applicationStore.getTermsOfService()
        } catch {
          if (componentIsMounted) {
            setTos({ enabled: false })
            setReady(true)
          }
          return
        }

        try {
          const tosAccepted =
            await stores.applicationStore.getTermsOfServiceStatus()

          if (componentIsMounted) {
            setTos({ ...tos, accepted: tosAccepted.accepted })
          }
        } catch (err) {
          if (componentIsMounted) {
            setTos(null)
          }
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

  const accept = async () =>
    await stores.applicationStore.acceptTermsOfService()

  return (
    <ToSContext.Provider
      value={{
        tos: tos,
        isReady: isReady,
        accept: accept,
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
    accepted: bool,
  }),
}

export default ToSContext
export { ToSContextProvider }
