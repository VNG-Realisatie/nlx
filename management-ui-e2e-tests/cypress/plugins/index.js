// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

// This function is called when a project is opened or re-opened (e.g. due to
// the project's config changing)

/**
 * @type {Cypress.PluginConfig}
 */
module.exports = (on, config) => {
  on('task', {
    log(message) {
      // eslint-disable-next-line no-console
      console.log(message)

      return null
    },
    table(message) {
      // eslint-disable-next-line no-console
      console.table(message)

      return null
    },
  })
}
