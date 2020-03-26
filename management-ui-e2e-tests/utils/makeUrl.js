import getBaseUrl from './getBaseUrl'

module.exports = (path = '') => {
  const separator = path.includes('?') ? '&' : '?'
  return `${getBaseUrl()}/${path}${separator}isTest`
}
