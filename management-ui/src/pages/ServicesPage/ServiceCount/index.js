// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

import React from 'react'
import { number } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { StyledCount, StyledAmount } from './index.styles'

const ServiceCount = ({ count }) => {
  const { t } = useTranslation()

  return (
    <StyledCount>
      <StyledAmount>{count}</StyledAmount>
      <small>{t('Services')}</small>
    </StyledCount>
  )
}

ServiceCount.propTypes = {
  count: number.isRequired,
}

export default ServiceCount
