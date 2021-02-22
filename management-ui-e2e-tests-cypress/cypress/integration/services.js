// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import logViolations from '../axe-utilities/log-violations'

function randomNumberString() {
  return `${Math.floor(Math.random() * 10000)}`
}

function getTestName() {
  const testName = Cypress.mocha.getRunner().suite.ctx.test.title
  return testName.replaceAll(' ', '-')
}

function generateServiceName() {
  return `service-e2e-${getTestName()}-${randomNumberString()}`
}

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

  it('Add and remove service', () => {
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
    cy.findByLabelText('Publiceren in de centrale directory').uncheck()

    cy.findByText('Service toevoegen').click()
    cy.dismissToaster('De service is toegevoegd')

    cy.visit('/services')
    cy.findByText(serviceName).click()
    cy.findByText('Verwijderen').click()
    cy.clickModalButton('Verwijderen')

    cy.dismissToaster('De service is verwijderd')
  })
})
