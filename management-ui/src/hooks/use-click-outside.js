// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useCallback, useEffect } from 'react'

const useClickOutside = (ref, handler) => {
  const handleClickOutside = useCallback(
    (event) => {
      if (ref.current && !ref.current.contains(event.target)) {
        handler({ event, ref })
      }
    },
    [ref, handler],
  )
  useEffect(() => {
    /**
     * Alert if clicked on outside of element
     */
    // Bind the event listener
    document.addEventListener('mousedown', handleClickOutside)
    return () => {
      // Unbind the event listener on clean up
      document.removeEventListener('mousedown', handleClickOutside)
    }
  }, [handleClickOutside, ref])
}

export default useClickOutside
