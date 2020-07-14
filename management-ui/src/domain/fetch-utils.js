// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import fetchDefaults from 'fetch-defaults'
import asyncMemoize from './async-memoize'

// needed to prevent caching on IE 11
// https://stackoverflow.com/questions/37755782/prevent-ie11-caching-get-call-in-angular-2/44561162#44561162
export const PREVENT_CACHING_HEADERS = {
  'Cache-Control': 'no-cache',
  Pragma: 'no-cache',
  Expires: 'Sat, 01 Jan 2000 00:00:00 GMT',
}

export const fetchWithoutCaching = fetchDefaults(global.fetch, {
  headers: {
    ...PREVENT_CACHING_HEADERS,
  },
})

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

export const throwOnError = (response) => {
  if (response.ok) return response
  throw new Error(statusErrorMessage[response.status] || genericErrorMessage)
}
