// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import logViolations from '../axe-utilities/log-violations'

describe('Directory page', () => {
  beforeEach(() => {
    cy.loginUsingDex()
    cy.visit('/directory')
    cy.injectAxe()
    cy.get('h1').should('contain', 'Directory')
  })

  it('Page is accessible', () => {
    cy.checkA11y(null, null, logViolations)
  })
})
