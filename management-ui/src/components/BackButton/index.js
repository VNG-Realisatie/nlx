// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { StyledBackButton, StyledIconChevron } from './index.styles'

const BackButton = ({ to }) => {
  const { t } = useTranslation()
  return (
    <StyledBackButton to={to} aria-label={t('Back')}>
      <StyledIconChevron />
      {t('Back')}
    </StyledBackButton>
  )
}
BackButton.propTypes = { to: string }
export default BackButton
