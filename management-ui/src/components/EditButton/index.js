// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { StyledPencil } from './index.styles'

const EditButton = (props) => {
  const { t } = useTranslation()
  return (
    <Button variant="secondary" {...props}>
      <StyledPencil />
      {t('Edit')}
    </Button>
  )
}

export default EditButton
