import React from 'react'
import { useTranslation } from 'react-i18next'
import {
  StyledIcon,
  StyledLink,
  StyledHomeLink,
  StyledNLXManagementLogo,
  StyledNav,
} from './index.styles'
import IconServices from './IconServices'

const PrimaryNavigation = () => {
  const { t } = useTranslation()
  return (
    <StyledNav>
      <StyledHomeLink
        to="/"
        title={t('NLX Dashboard homepage')}
        aria-label={t('Homepage')}
      >
        <StyledNLXManagementLogo />
      </StyledHomeLink>

      <StyledLink to="/services" aria-label={t('Services page')}>
        <StyledIcon as={IconServices} />
        {t('Services')}
      </StyledLink>
    </StyledNav>
  )
}

export default PrimaryNavigation
