// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconKey } from '../../../../../icons'
import { StyledButtonWithIcon } from './index.styles'

const QuickAccessButton = (props) => {
  const { t } = useTranslation()

  return (
    <StyledButtonWithIcon size="small" variant="link" {...props}>
      <IconKey />
      {t('Request')}
    </StyledButtonWithIcon>
  )
}

QuickAccessButton.propTypes = {
  onClick: func.isRequired,
}

export default QuickAccessButton
