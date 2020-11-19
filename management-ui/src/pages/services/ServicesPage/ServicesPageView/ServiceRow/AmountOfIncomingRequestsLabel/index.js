// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { number } from 'prop-types'
import { useTranslation } from 'react-i18next'

import { IconKey } from '../../../../../../icons'
import { StyledButton } from './index.styles'

const AmountOfIncomingRequestsLabel = ({ count, ...props }) => {
  const { t } = useTranslation()

  return (
    <StyledButton size="small" variant="link" {...props}>
      <IconKey inline />
      {t('requestWithCount', { count: count })}
    </StyledButton>
  )
}

AmountOfIncomingRequestsLabel.propTypes = {
  count: number,
}

AmountOfIncomingRequestsLabel.defaultProps = {
  count: 0,
}

export default AmountOfIncomingRequestsLabel
