// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
export function randomNumberString() {
  return `${Math.floor(Math.random() * 10000)}`
}

export function getTestName() {
  const testName = Cypress.mocha.getRunner().suite.ctx.test.title
  return testName.replaceAll(' ', '-')
}

export function generateServiceName() {
  return `service-e2e-${getTestName()}-${randomNumberString()}`
}
