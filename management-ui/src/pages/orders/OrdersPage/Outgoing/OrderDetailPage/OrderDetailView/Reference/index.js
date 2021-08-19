// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { IconBarcode } from '../../../../../../../icons'
import { StyledContainer, StyledLabel } from './index.styles'

const Reference = ({ value }) => {
  const { t } = useTranslation()
  return (
    <StyledContainer>
      <IconBarcode />
      {t('Reference')}
      <StyledLabel>{value}</StyledLabel>
    </StyledContainer>
  )
}

Reference.propTypes = {
  value: string.isRequired,
}

export default Reference
