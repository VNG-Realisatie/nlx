import i18n from 'i18next'
import { initReactI18next } from 'react-i18next'

i18n
  .use(initReactI18next) // passes i18n down to react-i18next
  .init(
    {
      lng: 'nl',
      fallbackLng: 'en',
      ns: ['common'],
      defaultNS: 'common',

      resources: {},

      // We do not use keys in form messages.welcome
      keySeparator: false,
      nsSeparator: false,

      interpolation: {
        escapeValue: false, // react already safes from xss
      },
    }
  )

export default i18n
