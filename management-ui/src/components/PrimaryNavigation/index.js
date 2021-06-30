// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import {
  IconArrowLeftRight,
  IconDirectory,
  IconTimeLine,
  IconUserCheck,
  IconSettings,
} from '../../icons'
import { ReactComponent as IconServices } from '../../icons/services2.svg'
import {
  Nav,
  StyledHomeLink,
  StyledLink,
  StyledIcon,
  InwaysIcon,
  ServicesIcon,
  DirectoryIcon,
  BarChartIcon,
  TimeIcon,
  OrdersLink,
  SettingsIcon,
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

        <StyledLink to="/inways" aria-label={t('Inways page')}>
          <InwaysIcon as={IconArrowLeftRight} size="x-large" />
          {t('Inways')}
        </StyledLink>

        <StyledLink to="/services" aria-label={t('Services page')}>
          <ServicesIcon as={IconServices} size="x-large" />
          {t('Services')}
        </StyledLink>

        <StyledLink to="/directory" aria-label={t('Directory page')}>
          <DirectoryIcon as={IconDirectory} size="x-large" />
          {t('Directory')}
        </StyledLink>

        <StyledLink to="/finances" aria-label={t('Finances page')}>
          <BarChartIcon>
            <div />
            <div />
            <div />
          </BarChartIcon>
          {t('Finances')}
        </StyledLink>

        <StyledLink to="/audit-log" aria-label={t('Audit log page')}>
          <TimeIcon as={IconTimeLine} size="x-large" />
          {t('Logs')}
        </StyledLink>
      </section>

      <section>
        <OrdersLink to="/orders" aria-label={t('Orders page')}>
          <StyledIcon as={IconUserCheck} size="x-large" />
          {t('Orders')}
        </OrdersLink>
        <StyledLink to="/settings" aria-label={t('Settings page')}>
          <SettingsIcon as={IconSettings} size="x-large" />
          {t('Settings')}
        </StyledLink>
      </section>
    </Nav>
  )
}

export default PrimaryNavigation
