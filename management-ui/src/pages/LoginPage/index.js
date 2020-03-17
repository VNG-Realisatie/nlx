import React from 'react'
import { useTranslation } from 'react-i18next'
import {
  StyledContainer,
  StyledContent,
  StyledSidebar,
  StyledNLXManagementLogo,
} from './index.styles'

const LoginPage = () => {
  const { t } = useTranslation()

  return (
    <StyledContainer>
      <StyledSidebar>
        <StyledNLXManagementLogo />
      </StyledSidebar>
      <StyledContent>
        <h1>{t('Welkom')}</h1>
        <p>{t('Login om verder te gaan.')}</p>
      </StyledContent>
    </StyledContainer>
  )
}

export default LoginPage
