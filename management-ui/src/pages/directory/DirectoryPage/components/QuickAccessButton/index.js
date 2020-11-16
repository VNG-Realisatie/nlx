// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconKey } from '../../../../../icons'
import { StyledButton } from './index.styles'

const QuickAccessButton = (props) => {
  const { t } = useTranslation()

  return (
    <StyledButton size="small" variant="link" {...props}>
      <IconKey inline />
      {t('Request')}
    </StyledButton>
  )
}

QuickAccessButton.propTypes = {
  onClick: func.isRequired,
}

export default QuickAccessButton
