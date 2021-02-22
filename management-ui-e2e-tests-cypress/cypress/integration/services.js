// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import logViolations from '../axe-utilities/log-violations'

describe('Services', () => {
  beforeEach(() => {
    cy.loginUsingDex()
    cy.visit('/services')
    cy.injectAxe()
    cy.get('h1').should('contain', 'Services')
  })

  it('Overview list is accessible', () => {
    cy.checkA11y(null, null, logViolations)
  })
})
