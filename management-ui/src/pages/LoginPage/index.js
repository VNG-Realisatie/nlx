// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { Button } from '@commonground/design-system'
import UserContext from '../../user-context'
import OrganizationName from '../../components/OrganizationName'
import {
  StyledContainer,
  StyledContent,
  StyledExternalLink,
  StyledNLXManagementLogo,
  StyledOrganization,
  StyledIconOrganization,
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
            <p>{t('Log in to continue')}</p>

            <StyledOrganization>
              <StyledIconOrganization />
              <OrganizationName data-testid="organizationName" />
            </StyledOrganization>

            <Button data-testid="login" as="a" href="/oidc/authenticate">
              {t('Log in with organization account')} <StyledExternalLink />
            </Button>
          </>
        ) : (
          <form method="POST" action="/oidc/logout">
            <Button data-testid="logout" type="submit">
              {t('Log out')}
            </Button>
          </form>
        )}
      </StyledContent>
    </StyledContainer>
  )
}

export default LoginPage
