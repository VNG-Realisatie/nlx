// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
const DEFAULT_BASE_URL = 'http://management.organization-a.nlx.local:3011'
const constructBaseUrl = () => {
  const url = process.env.URL

  if (!url) {
    console.warn(
      `Warning: Environment variable 'URL' not set. using the default '${DEFAULT_BASE_URL}'.`,
    )
  }

  return url || DEFAULT_BASE_URL
}

let baseUrl = ''

module.exports = function () {
  // Cache baseUrl, se we don't get many "URL env var" warnings
  if (!baseUrl) baseUrl = constructBaseUrl()
  return baseUrl
}
