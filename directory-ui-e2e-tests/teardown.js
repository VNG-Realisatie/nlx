// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

const rimraf = require('rimraf');

module.exports = async function () {
  // close the browser instance
  await global.__BROWSER_GLOBAL__.close();

  const DIR = global.__JEST_TEMP_DIR__

  // clean-up the wsEndpoint file
  rimraf.sync(DIR);
};
