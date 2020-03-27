// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

// This file is in ES5 on purpose, to be able to be shared between `i18n.js` for use in the application and in
// `../i18next.parser.config.js` for use in build validation and generation of translations.

module.exports = {
  lng: 'nl',
  fallbackLng: 'en',
  ns: ['common'],
  defaultNS: 'common',

  backend: {
    loadPath: `${process.env.PUBLIC_URL}/i18n/{{lng}}/{{ns}}.json`,
  },

  // We do not use keys in form messages.welcome
  keySeparator: false,
  nsSeparator: false,

  interpolation: {
    escapeValue: false, // react already safes from xss
  },
}
