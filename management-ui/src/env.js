// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

// in order to be able to specify base URLs of external services at run time,
// we load an `env.js` file which populates the global window._env variable.

// If you want to consume any of these values, please do so by importing this
// file, so we can group the values and provide sensible defaults in one place.

window._env = window._env || {}

export default {
  oidcBaseUrl: window._env.OIDC_BASE_URL || '/oidc',
  managementApiBaseUrl: window._env.MANAGEMENT_API_BASE_URL || '/api',
}
