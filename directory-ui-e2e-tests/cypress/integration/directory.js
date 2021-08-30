// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
describe('Directory', () => {
  beforeEach(() => {
    cy.visit('/')
  })

  it('should display the support email address in the detail pane', () => {
    cy.get('input[placeholder="Zoeken…"]').type(
      'basisregister-fictieve-kentekens',
    )

    cy.findByTestId('directory-services').get('tbody tr').first().click()

    cy.findByText('support@nlx.io').should('be.visible')
  })
})
