// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import React, { useContext } from 'react'
import { useTranslation } from 'react-i18next'
import { useLocation } from 'react-router-dom'
import { Button } from '@commonground/design-system'
import UserContext from '../../user-context'
import OrganizationName from '../../components/OrganizationName'
import {
  Wrapper,
  Content,
  StyledAlert,
  StyledIconExternalLink,
  StyledNLXManagementLogo,
  StyledOrganization,
  StyledIconOrganization,
  StyledSidebar,
} from './index.styles'

function useLoginErrorMessage() {
  const { t } = useTranslation()
  const loc = useLocation()

  const match = loc.hash.match(/#auth-(.*)/)

  if (match === null) {
    return null
  }

  switch (match[1]) {
    case 'fail':
      return t('Something went wrong when logging in. Please try again.')

    case 'missing-user':
      return t("User doesn't exist in NLX Management.")

    default:
      return null
  }
}

const LoginOIDCPage = () => {
  const { t } = useTranslation()
  const { user } = useContext(UserContext)
  const loginError = useLoginErrorMessage()

  return (
    <Wrapper>
      <StyledSidebar>
        <StyledNLXManagementLogo />
      </StyledSidebar>
      <Content>
        <h1>{t('Welcome')}</h1>

        {!user ? (
          <>
            <p>{t('Log in to continue')}</p>

            {loginError !== null && (
              <StyledAlert
                data-testid="login-error-message"
                variant="error"
                title={loginError}
              />
            )}

            <StyledOrganization>
              <StyledIconOrganization inline />
              <OrganizationName data-testid="organizationName" />
            </StyledOrganization>

            <Button data-testid="login" as="a" href="/oidc/authenticate">
              {t('Log in with organization account')} <StyledIconExternalLink />
            </Button>
          </>
        ) : (
          <form method="POST" action="/oidc/logout">
            <Button data-testid="logout" type="submit">
              {t('Log out')}
            </Button>
          </form>
        )}
      </Content>
    </Wrapper>
  )
}

export default LoginOIDCPage
