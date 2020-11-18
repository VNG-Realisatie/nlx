// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconKey } from '../../../../icons'
import { StyledButtonWithIcon } from './index.styles'

const AmountOfIncomingRequestsLabel = ({ count, ...props }) => {
  const { t } = useTranslation()

  return (
    <StyledButtonWithIcon size="small" variant="link" {...props}>
      <IconKey />
      {t('requestWithCount', { count: count })}
    </StyledButtonWithIcon>
  )
}

AmountOfIncomingRequestsLabel.propTypes = {
  count: number,
}

AmountOfIncomingRequestsLabel.defaultProps = {
  count: 0,
}

export default AmountOfIncomingRequestsLabel
