import i18n from 'i18next'
import XHR from 'i18next-xhr-backend'
import { initReactI18next } from 'react-i18next'

const isTest = !!(window && window.location.search.match(/[?&]isTest/))

i18n
  .use(XHR)
  .use(initReactI18next) // passes i18n down to react-i18next
  .init({
    lng: isTest ? 'cimode' : 'nl',
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
  })

export default i18n
