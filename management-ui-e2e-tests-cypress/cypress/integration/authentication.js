// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import logViolations from '../axe-utilities/log-violations'

describe('Login page', () => {
  beforeEach(() => {
    cy.visit('/login')
    cy.injectAxe()
    cy.get('h1').should('contain', 'Welkom')
  })

  it('Page is accessible', () => {
    cy.checkA11y(null, null, logViolations)
  })

  it('Displays the organization name', () => {
    cy.get('[data-testid=organizationName]').should(
      'contain',
      Cypress.env('LOGIN_ORGANIZATION_NAME'),
    )
  })
})

describe('Authenticating', () => {
  it('Redirects to the Inways page after successful login', () => {
    cy.loginUsingDex()
    cy.url().should('contain', '/inways')
  })

  it('Displays error on login with invalid credentials', () => {
    cy.visit('/login')
    cy.get('[data-testid="login"]').click()
    cy.get('#login').type('foo')
    cy.get('#password').type('bar')
    cy.findByText('Login').click()

    cy.findByText('Invalid Email Address and password.').should('be.visible')
  })
})
