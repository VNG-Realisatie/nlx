// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { IconMoneyEuroCircleLine } from '../../../icons'
import { StyledLabel, StyledDetailHeading } from './index.styles'

const CollapsibleHeader = ({ label }) => {
  const { t } = useTranslation()

  return (
    <StyledDetailHeading>
      <IconMoneyEuroCircleLine />
      {t('Costs')}
      <StyledLabel>{label}</StyledLabel>
    </StyledDetailHeading>
  )
}

CollapsibleHeader.propTypes = {
  label: string,
}

export default CollapsibleHeader
