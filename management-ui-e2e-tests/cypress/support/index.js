// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import 'cypress-axe'
import './commands'

Cypress.on('uncaught:exception', () => {
  return false
})
