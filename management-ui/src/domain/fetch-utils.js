// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import asyncMemoize from './async-memoize'

const responseClone = (response) => response.clone()

const memoizedFetch = asyncMemoize(fetch)

export const fetchWithCaching = (...args) =>
  memoizedFetch(...args).then(responseClone)

// Expose the memoize instance for tests to be able to clear it
fetchWithCaching.memo = memoizedFetch

const statusErrorMessage = {
  400: 'invalid user input',
  401: 'no user is authenticated',
  403: 'forbidden',
  404: 'not found',
}
const genericErrorMessage = 'unable to handle the request'

export const throwOnError = (response, customStatusErrorMessage = {}) => {
  if (response.ok) return response
  throw new Error(
    { ...statusErrorMessage, ...customStatusErrorMessage }[response.status] ||
      genericErrorMessage,
  )
}
