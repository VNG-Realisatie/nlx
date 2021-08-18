// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
describe('Directory', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('should display the support email address in the detail pane', () => {
    cy.get('[data-testid=directory-service-row]').within(() => {
      cy.findAllByText(/basisregister-fictieve-kentekens/)
        .last()
        .click()
    })

    cy.findByText('support@nlx.io').should('be.visible')
  })
})
