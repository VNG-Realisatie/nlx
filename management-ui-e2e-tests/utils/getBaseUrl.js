// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

const constructBaseUrl = () => {
  const url = process.env.URL

  if (!url) {
    console.warn(`Warning: Environment variable 'URL' not set. using the default 'http://localhost:3002'.`)
  }

  return url || 'http://localhost:3002'
}

let baseUrl = ''

module.exports = function () {
  // Cache baseUrl, se we don't get many "URL env var" warnings
  if (!baseUrl) baseUrl = constructBaseUrl()
  return baseUrl
}
