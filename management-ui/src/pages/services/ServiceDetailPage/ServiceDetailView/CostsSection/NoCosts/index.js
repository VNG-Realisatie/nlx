// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { IconMoneyEuroCircleLine } from '../../../../../../icons'
import { StyledContainer, StyledLabel } from './index.styles'

const NoCosts = () => {
  const { t } = useTranslation()
  return (
    <StyledContainer>
      <IconMoneyEuroCircleLine />
      {t('Costs')}
      <StyledLabel>{t('Free')}</StyledLabel>
    </StyledContainer>
  )
}

export default NoCosts
