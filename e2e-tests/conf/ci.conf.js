/**
 * Copyright © VNG Realisatie 2022
 * Licensed under the EUPL
 */

exports.config = {
  server: "http://localhost:4444/wd/hub",
  username: "",
  accessKey: "",

  capabilities: [
    {
      browserName: "chrome",
    },
  ],
};

// Code to support common capabilities
exports.config.capabilities.forEach(function (caps) {
  for (var i in exports.config.commonCapabilities)
    caps[i] = caps[i] || exports.config.commonCapabilities[i];
});
