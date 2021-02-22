// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import logViolations from '../axe-utilities/log-violations'
import { generateServiceName } from './helpers/services'

describe('Inways', () => {
  beforeEach(() => {
    cy.loginUsingDex()
    cy.visit('/inways')
    cy.injectAxe()
    cy.get('h1').should('contain', 'Inways')
  })

  it('Screens are accessible and details can be closed', () => {
    cy.checkA11y(null, null, logViolations)

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

  it('Deeplink to inway details and go to connected service page', () => {
    const serviceName = generateServiceName()

    // Create service with inway
    cy.visit('/services/add-service')
    cy.findByLabelText('Servicenaam').type(serviceName)
    cy.findByLabelText('API endpoint URL').type('my-service.test:8000')
    cy.findByLabelText(Cypress.env('INWAY_NAME')).check()
    cy.findByText('Service toevoegen').click()
    cy.dismissToaster('De service is toegevoegd')

    // Go to service via inway
    cy.visit(`/inways/${Cypress.env('INWAY_NAME')}`)
    cy.findByText('Gekoppelde services').click()
    cy.findByText(serviceName).click()

    // Clean up
    cy.findByText('Verwijderen').click()
    cy.findByText('Weet je het zeker?').should('exist')
    cy.clickModalButton('Verwijderen')
  })
})
