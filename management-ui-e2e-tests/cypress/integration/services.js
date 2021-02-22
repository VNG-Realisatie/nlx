// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import logViolations from '../axe-utilities/log-violations'
import { generateServiceName } from './helpers/services'

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

  it('Add edit and remove service', () => {
    const serviceName = generateServiceName()

    cy.visit('/services/add-service')
    cy.findByLabelText('Servicenaam').type(serviceName)
    cy.findByLabelText('API endpoint URL').type('my-service.test:8000')
    cy.findByLabelText('API documentatie URL').type('my-service.test:8000/docs')
    cy.findByLabelText('API specificatie URL').type(
      'my-service.test:8000/openapi.json',
    )
    cy.findByLabelText('Technisch support e-mailadres').type(
      'tech@organization.test',
    )
    cy.findByLabelText('Publiek support e-mailadres').type(
      'public@organization.test',
    )

    cy.findByText('Service toevoegen').click()
    cy.dismissToaster('De service is toegevoegd')

    // Test detail page
    cy.visit('/services')
    cy.injectAxe()
    cy.findByText(serviceName).click()

    // disable 'tabindex' because the 'focus-lock' dependency
    // creates an element with tabindex="1"
    // https://github.com/theKashey/react-focus-lock/blob/2b6ae70f0b15046ee3ac3227c53bb7c21f551ff4/src/Lock.js#L127
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

    // Edit via detail view
    cy.findByText('Aanpassen').click()
    cy.findByText('Service nog niet benaderbaar').should('exist')
    cy.checkA11y(null, null, logViolations)
    cy.findByLabelText(Cypress.env('INWAY_NAME')).check()
    cy.findByText('Service bijwerken').click()
    cy.dismissToaster('De service is bijgewerkt')

    // Edit via deeplink
    cy.visit(`/services/${serviceName}/edit-service`)
    cy.findByText('Terug').click()

    // Remove
    cy.findByText('Verwijderen').click()
    cy.findByText('Weet je het zeker?').should('exist')
    cy.clickModalButton('Verwijderen')

    cy.dismissToaster('De service is verwijderd')
  })
})
