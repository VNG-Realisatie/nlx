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
        <h1>{t('Welcome')}</h1>
        <p>{t('Log in to continue')}</p>
      </StyledContent>
    </StyledContainer>
  )
}

export default LoginPage
