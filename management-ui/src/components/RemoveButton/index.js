// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { StyledBin } from './index.styles'

const RemoveButton = (props) => {
  const { t } = useTranslation()
  return (
    <Button variant="danger" {...props}>
      <StyledBin />
      {t('Remove')}
    </Button>
  )
}

export default RemoveButton
