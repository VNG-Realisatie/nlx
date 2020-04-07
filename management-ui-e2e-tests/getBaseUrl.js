// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

const getBaseUrl = () => {
  const url = process.env.URL

  if (!url) {
    console.warn(`environment variable 'URL' not set. using the default 'http://localhost:3002'`)
  }

  return url || 'http://localhost:3002'
}


module.exports = getBaseUrl
