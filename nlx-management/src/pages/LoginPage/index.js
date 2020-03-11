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
        <p>{t('Welkom')}</p>
      </StyledContent>
    </StyledContainer>
  )
}

export default LoginPage
