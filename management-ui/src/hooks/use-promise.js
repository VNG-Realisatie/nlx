// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { useState, useEffect } from 'react'

const usePromise = (getPromise, ...args) => {
  const [result, setResult] = useState(null)
  const [isReady, setIsReady] = useState(false)
  const [error, setError] = useState(null)

  const [reloadCounter, setReloadCounter] = useState(0)

  const reload = () => {
    setReloadCounter(reloadCounter + 1)
  }

  useEffect(() => {
    const resolvePromise = async () => {
      try {
        const res = await getPromise(...args)
        setResult(res)
      } catch (error) {
        setError(error)
      }
      setIsReady(true)
    }

    resolvePromise()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [...(args || []), reloadCounter])
  return { isReady, error, result, reload }
}

export default usePromise
