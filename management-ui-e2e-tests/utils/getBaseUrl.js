const getBaseUrl = () => {
  const url = process.env.URL

  if (!url) {
    console.warn(`environment variable 'URL' not set. using the default 'http://localhost:3000'`)
  }

  return url || 'http://localhost:3000'
}


module.exports = getBaseUrl
