// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

// This file is in ES5 on purpose, to be able to be shared between `i18n.js` for use in the application and in
// `../i18next.parser.config.js` for use in build validation and generation of translations.
const dayjs = require('dayjs')
const localizedFormat = require('dayjs/plugin/localizedFormat')
dayjs.extend(localizedFormat)

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

    format: (value, format, lng) => {
      /**
       * Date
       *
       * In JS, pass: { date: new Date('<date string>') }
       * In language file use:  {{date, DD MMMM YYYY}} for custom formatting
       * For default value use: {{date,}}
       *
       * Default shows localized date: https://day.js.org/docs/en/display/format#localized-formats
       */
      if (value instanceof Date) {
        return dayjs(value).format(format || 'LL')
      }
    },
  },
}
