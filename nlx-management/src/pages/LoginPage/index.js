import React from 'react'
import { useTranslation } from 'react-i18next'
import {
  StyledContainer,
  StyledContent,
  StyledSidebar,
  StyledNLXLogo,
} from './index.styles'

const LoginPage = () => {
  const { t } = useTranslation()

  return (
    <StyledContainer>
      <StyledSidebar>
        <StyledNLXLogo />
        <br />
        Management
      </StyledSidebar>
      <StyledContent>
        <h2>{t('Welkom')}</h2>
        <p>{t('Login om verder te gaan.')}</p>
      </StyledContent>
    </StyledContainer>
  )
}

export default LoginPage
