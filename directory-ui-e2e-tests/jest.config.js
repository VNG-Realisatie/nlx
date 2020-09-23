// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

// Setup and teardown of the Jest environment is based on
// an example from the Jest docs. See:
// https://jestjs.io/docs/en/puppeteer#custom-example-without-jest-puppeteer-preset

module.exports = {
  globalSetup: './setup.js',
  globalTeardown: './teardown.js',
  testEnvironment: './puppeteer_environment.js',
};
