import React, { useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from '@commonground/design-system'
import UserContext from '../../user-context/UserContext'
import {
  StyledContainer,
  StyledContent,
  StyledExternalLink,
  StyledNLXManagementLogo,
  StyledSidebar,
} from './index.styles'

const LoginPage = () => {
  const { t } = useTranslation()
  const { user } = useContext(UserContext)

  return (
    <StyledContainer>
      <StyledSidebar>
        <StyledNLXManagementLogo />
      </StyledSidebar>
      <StyledContent>
        <h1>{t('Welcome')}</h1>
        {!user ? (
          <>
            <p>{t('Log in to continue.')}</p>
            <Button id="login" as="a" href="/oidc/authenticate">
              {t('Log in with organization account')} <StyledExternalLink />
            </Button>
          </>
        ) : (
          <>
            <Button id="logout" as="a" href="/oidc/logout">
              {t('Log out')}
            </Button>
          </>
        )}
      </StyledContent>
    </StyledContainer>
  )
}

export default LoginPage
