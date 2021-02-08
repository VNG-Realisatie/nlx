// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { Suspense } from 'react'
import { func, node } from 'prop-types'
import { ThemeProvider } from 'styled-components'
import {
  GlobalStyles as DSGlobalStyles,
  ToasterProvider,
} from '@commonground/design-system'
import GlobalStyles from '../components/GlobalStyles'
import theme from '../theme'
import SettingsRepository from '../domain/settings-repository'
import { StyledContainer } from './index.styles'
import useInitializeApplicationStoreFromSettings from './use-initialize-application-store-from-settings'

const App = ({ getSettings, children, ...props }) => {
  useInitializeApplicationStoreFromSettings(getSettings)

  return (
    <StyledContainer {...props}>
      <ThemeProvider theme={theme}>
        <DSGlobalStyles />
        <GlobalStyles />

        {/* Suspense is required for XHR backend i18next */}
        <Suspense fallback={null}>
          <ToasterProvider>{children}</ToasterProvider>
        </Suspense>
      </ThemeProvider>
    </StyledContainer>
  )
}

App.propTypes = {
  getSettings: func,
  children: node,
}

App.defaultProps = {
  getSettings: SettingsRepository.getGeneralSettings,
}

export default App
