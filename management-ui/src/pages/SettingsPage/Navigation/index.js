// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { StyledLink, StyledNav } from './index.styles'

const Navigation = () => {
  const { t } = useTranslation()

  return (
    <StyledNav>
      <StyledLink to={`general`} aria-label={t('General settings')}>
        {t('General')}
      </StyledLink>
    </StyledNav>
  )
}

export default Navigation
