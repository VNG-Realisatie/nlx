// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import { IconPencil } from '../../icons'

const EditButton = (props) => {
  const { t } = useTranslation()
  return (
    <Button variant="secondary" {...props}>
      <IconPencil inline />
      {t('Edit')}
    </Button>
  )
}

export default EditButton
