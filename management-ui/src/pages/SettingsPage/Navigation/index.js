// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { useRouteMatch } from 'react-router-dom'
import { StyledLink, StyledNav } from './index.styles'

const Navigation = () => {
  const { t } = useTranslation()
  const { path } = useRouteMatch('/settings')
  return (
    <StyledNav>
      <StyledLink to={`${path}/general`} aria-label={t('General settings')}>
        {t('General')}
      </StyledLink>

      <StyledLink to={`${path}/insight`} aria-label={t('Insight settings')}>
        {t('Insight')}
      </StyledLink>
    </StyledNav>
  )
}

export default Navigation
