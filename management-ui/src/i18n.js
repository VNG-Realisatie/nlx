// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import i18n from 'i18next'
import HttpApi from 'i18next-http-backend'
import { initReactI18next } from 'react-i18next'
import dayjs from 'dayjs'

import options from './i18n.options'

i18n
  .use(HttpApi)
  .use(initReactI18next) // passes i18n down to react-i18next
  .on('languageChanged', async (lng) => {
    try {
      const localeData = await import(`dayjs/locale/${lng}.js`)
      dayjs.locale(localeData.default)
    } catch (e) {
      console.error(`Can't find locale for language ${lng}`)
    }
  })
  .init(options)

export default i18n
