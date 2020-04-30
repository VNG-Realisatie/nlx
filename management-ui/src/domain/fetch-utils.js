// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import fetchDefaults from 'fetch-defaults'

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
