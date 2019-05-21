// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

const appendUrlPrefix = (url, prefix) => (prefix ? `${prefix}${url}` : url)

export const relativeToFullUrl = (url) =>
    appendUrlPrefix(url, process.env.REACT_APP_API_BASE_URL || '/api')
