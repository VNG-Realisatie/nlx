// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

describe('Settings', () => {
  beforeEach(() => {
    cy.loginUsingDex()
    cy.visit('/settings')
    cy.get('h1').should('contain', 'Instellingen')
  })

  it('Not specifying the organization inway', () => {
    cy.get('#organizationInway').click()
    cy.get('.ReactSelect__menu-list').findByText(/Geen/).click()
    cy.findByText('Instellingen opslaan').click()

    cy.findByText('Weet je het zeker?').should('exist')
    cy.clickModalButton('Opslaan')
    cy.dismissToaster('De instellingen zijn bijgewerkt')

    cy.findByText('Organisatie inway').should('exist')

    cy.get('#organizationInway').click()
    cy.get('.ReactSelect__menu-list')
      .findByText(new RegExp(Cypress.env('INWAY_NAME'), 'gm'))
      .click()

    cy.findByText('Instellingen opslaan').click()
    cy.findByText(
      'Er kunnen geen toegangsverzoeken ontvangen worden. Stel in welke inway toegangsverzoeken moet afhandelen.',
    ).should('not.exist')
  })
})
