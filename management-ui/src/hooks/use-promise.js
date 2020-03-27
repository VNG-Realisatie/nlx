// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { useState, useEffect } from 'react'

const usePromise = (getPromise, ...args) => {
  const [result, setResult] = useState(null)
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState(null)

  useEffect(() => {
    const resolvePromise = async () => {
      setLoading(true)
      try {
        const res = await getPromise(...args)
        setResult(res)
        setLoading(false)
      } catch (error) {
        setError(error)
        setLoading(false)
      }
    }

    resolvePromise()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])
  return { loading, error, result }
}

export default usePromise
