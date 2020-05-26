// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import pMemoize from 'p-memoize'

// Wrap function with a memoizer.
// In case this fails on an older browser fall back to the unwrapped function.
export default function asyncMemoize(fn, ...args) {
  try {
    return pMemoize(fn, ...args)
  } catch {
    return fn
  }
}

export const clear = pMemoize.clear
