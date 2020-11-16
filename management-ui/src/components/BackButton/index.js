// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string } from 'prop-types'
import { Link } from 'react-router-dom'
import { Button } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { StyledIconChevron } from './index.styles'

const BackButton = ({ to }) => {
  const { t } = useTranslation()
  return (
    <Button as={Link} to={to} variant="link" aria-label={t('Back')}>
      <StyledIconChevron inline />
      {t('Back')}
    </Button>
  )
}

BackButton.propTypes = { to: string.isRequired }

export default BackButton
