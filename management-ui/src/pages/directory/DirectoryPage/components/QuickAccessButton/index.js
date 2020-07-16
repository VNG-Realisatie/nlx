// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconKey } from '../../../../../icons'
import { AccessButton } from './index.styles'

const QuickAccessButton = (props) => {
  const { t } = useTranslation()

  return (
    <AccessButton {...props}>
      <IconKey />
      {t('Request')}
    </AccessButton>
  )
}

QuickAccessButton.propTypes = {
  onClick: func.isRequired,
}

export default QuickAccessButton
