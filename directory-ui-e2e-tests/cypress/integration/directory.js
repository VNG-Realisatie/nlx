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
      'kentekenregister',
    )

    cy.findByText(/kentekenregister/).click()

    cy.get('[data-test="service-detail-pane"]')
      .findByText('public-support@nlx.io')
      .should('be.visible')
  })
})
