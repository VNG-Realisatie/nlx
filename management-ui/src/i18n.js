// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import i18n from 'i18next'
import XHR from 'i18next-xhr-backend'
import { initReactI18next } from 'react-i18next'
import options from './i18n.options'

i18n
  .use(XHR)
  .use(initReactI18next) // passes i18n down to react-i18next
  .init(options)

export default i18n
