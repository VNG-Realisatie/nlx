// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import { ThemeProvider } from 'styled-components'
import { I18nextProvider } from 'react-i18next'
import { node } from 'prop-types'
import React from 'react'
import { ToasterProvider } from '@commonground/design-system'
import theme from '../theme'
import i18n from './i18nTestConfig'

const AllProviders = ({ children }) => (
  <ThemeProvider theme={theme}>
    <I18nextProvider i18n={i18n}>
      <ToasterProvider>{children}</ToasterProvider>
    </I18nextProvider>
  </ThemeProvider>
)

AllProviders.propTypes = {
  children: node,
}

export default AllProviders
