// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import '@testing-library/cypress/add-commands'

Cypress.Commands.add('loginUsingDex', () => {
  Cypress.log({ name: 'login using dex' })

  cy.visit('login')
  cy.get('[data-testid="login"]').click()
  cy.get('#login').type(Cypress.env('LOGIN_USERNAME'))
  cy.get('#password').type(Cypress.env('LOGIN_PASSWORD'))
  cy.findByText('Login').click()
  cy.get('button[type="submit"]').findByText('Grant Access').click()
})
