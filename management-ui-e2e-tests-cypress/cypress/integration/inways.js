// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import logViolations from '../axe-utilities/log-violations'

describe('Inways', () => {
  beforeEach(() => {
    cy.loginUsingDex()
    cy.visit('/inways')
    cy.injectAxe()
    cy.get('h1').should('contain', 'Inways')
  })

  it('Overview list is accessible', () => {
    cy.checkA11y(null, null, logViolations)
  })

  it('Inway details are displayed and can be closed', () => {
    cy.findByText(Cypress.env('INWAY_NAME')).click()
    cy.checkA11y(
      null,
      {
        rules: {
          tabindex: {
            enabled: false,
          },
        },
      },
      logViolations,
    )
    cy.closeTopDrawer()
    cy.location().should('match', /\/inways$/)
  })
})
