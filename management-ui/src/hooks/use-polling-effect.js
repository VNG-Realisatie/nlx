// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//

import { useRef, useEffect } from 'react'

// hook from https://archive.ph/wip/dsFYc
function usePollingEffect(
  asyncCallback,
  dependencies = [],
  { interval = 3000, onCleanUp = () => {} } = {},
) {
  const timeoutIdRef = useRef(null)

  useEffect(() => {
    let _stopped = false
    // Side note: preceding semicolon needed for IIFEs.
    ;(async function pollingCallback() {
      try {
        await asyncCallback()
      } finally {
        // Set timeout after it finished, unless stopped
        timeoutIdRef.current =
          !_stopped && setTimeout(pollingCallback, interval)
      }
    })()

    // Clean up if dependencies change
    return () => {
      _stopped = true // prevent racing conditions
      clearTimeout(timeoutIdRef.current)
      onCleanUp()
    }
  }, [...dependencies, interval]) // eslint-disable-line react-hooks/exhaustive-deps
}

export default usePollingEffect
