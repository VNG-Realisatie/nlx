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
    document.addEventListener('mousedown', handleClickOutside)

    return () => {
      document.removeEventListener('mousedown', handleClickOutside)
    }
  }, [handleClickOutside, ref])
}

export default useClickOutside
