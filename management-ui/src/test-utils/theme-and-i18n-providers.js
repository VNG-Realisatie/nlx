// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import { ThemeProvider } from 'styled-components'
import { I18nextProvider } from 'react-i18next'
import { node } from 'prop-types'
import React from 'react'
import theme from '../theme'
import i18n from './i18nTestConfig'

const ThemeAndI18nProviders = ({ children }) => (
  <ThemeProvider theme={theme}>
    <I18nextProvider i18n={i18n}>{children}</I18nextProvider>
  </ThemeProvider>
)

ThemeAndI18nProviders.propTypes = {
  children: node,
}

export default ThemeAndI18nProviders
