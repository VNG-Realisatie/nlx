// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import {
  IconArrowLeftRight,
  IconDirectory,
  IconServices,
  IconBarChart,
  IconTimeLine,
  IconUserCheck,
  IconSettings,
} from '../../icons'
import {
  Nav,
  StyledHomeLink,
  StyledLink,
  StyledIcon,
  StyledNLXManagementLogo,
} from './index.styles'

const PrimaryNavigation = () => {
  const { t } = useTranslation()
  return (
    <Nav aria-labelledby="nlx-home">
      <section>
        <StyledHomeLink
          to="/"
          id="nlx-home"
          title={t('NLX Dashboard homepage')}
          aria-label={t('Homepage')}
        >
          <StyledNLXManagementLogo />
        </StyledHomeLink>

        <StyledLink
          to="/inways-and-outways"
          aria-label={t('Inways and Outways page')}
        >
          <StyledIcon as={IconArrowLeftRight} size="x-large" />
          {t('Inways and Outways')}
        </StyledLink>

        <StyledLink to="/services" aria-label={t('Services page')}>
          <StyledIcon as={IconServices} size="x-large" />
          {t('Services')}
        </StyledLink>

        <StyledLink to="/directory" aria-label={t('Directory page')}>
          <StyledIcon as={IconDirectory} size="x-large" />
          {t('Directory')}
        </StyledLink>

        <StyledLink to="/finances" aria-label={t('Finances page')}>
          <StyledIcon as={IconBarChart} size="x-large" />
          {t('Finances')}
        </StyledLink>

        <StyledLink to="/audit-log" aria-label={t('Audit log page')}>
          <StyledIcon as={IconTimeLine} size="x-large" />
          {t('Logs')}
        </StyledLink>
      </section>

      <section>
        <StyledLink to="/orders" aria-label={t('Orders page')}>
          <StyledIcon as={IconUserCheck} size="x-large" />
          {t('Orders')}
        </StyledLink>
        <StyledLink to="/settings" aria-label={t('Settings page')}>
          <StyledIcon as={IconSettings} size="x-large" />
          {t('Settings')}
        </StyledLink>
      </section>
    </Nav>
  )
}

export default PrimaryNavigation
