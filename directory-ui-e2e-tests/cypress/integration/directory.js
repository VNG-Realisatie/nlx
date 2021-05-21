// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

describe('Directory', () => {
  beforeEach(() => {
    cy.visit('/')
    cy.get('.nav-item.active a').should('contain', 'Directory')
  })

  it('should display the support email address in the detail pane', () => {
    cy.get('input[placeholder="Search for an organization or service…"]').type(
      'basisregister-fictieve-kentekens',
    )

    cy.findByText(/basisregister-fictieve-kentekens/).click()

    cy.get('[data-test="service-detail-pane"]')
      .findByText('support@nlx.io')
      .should('be.visible')
  })
})
