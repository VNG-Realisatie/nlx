// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { Suspense, useContext, useEffect } from 'react'
import { node } from 'prop-types'
import { ThemeProvider } from 'styled-components'
import {
  GlobalStyles as DSGlobalStyles,
  ToasterProvider,
} from '@commonground/design-system'
import UserContext from '../user-context'
import '@fontsource/source-sans-pro/latin.css'
import GlobalStyles from '../components/GlobalStyles'
import theme from '../theme'
import { useApplicationStore } from '../hooks/use-stores'
import { StyledContainer } from './index.styles'

const App = ({ children, ...props }) => {
  const { user, isReady } = useContext(UserContext)
  const applicationStore = useApplicationStore()

  useEffect(() => {
    const fetch = async () => {
      if (!isReady) return
      if (user === null) return
      if (applicationStore.isOrganizationInwaySet !== null) return

      const settings = await applicationStore.getGeneralSettings()

      applicationStore.updateOrganizationInway({
        isOrganizationInwaySet: !!settings.organizationInway,
      })
    }

    fetch()
  }, [user, isReady]) // eslint-disable-line react-hooks/exhaustive-deps

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
  children: node,
}

export default App
