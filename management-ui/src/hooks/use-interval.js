// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { useEffect, useRef } from 'react'

export const DEFAULT_INTERVAL = 3000

function useInterval(callback, delay) {
  const savedCallback = useRef()
  delay = delay || DEFAULT_INTERVAL

  useEffect(() => {
    savedCallback.current = callback
  }, [callback])

  useEffect(() => {
    function tick() {
      savedCallback.current()
    }

    const id = setInterval(tick, delay)
    return () => clearInterval(id)
  }, [delay])
}

export default useInterval
