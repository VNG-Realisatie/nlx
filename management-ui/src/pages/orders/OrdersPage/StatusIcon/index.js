// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { bool } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { IconStateUp, IconStateDown } from '../../../../icons'
import { StyledWrapper } from './index.styles'

const getStatusIcon = (active, t) => {
  if (active) {
    return (
      <>
        <IconStateUp title={t('Active')} />
      </>
    )
  } else {
    return (
      <>
        <IconStateDown title={t('Inactive')} />
      </>
    )
  }
}

const StatusIcon = ({ active, ...props }) => {
  const { t } = useTranslation()

  return <StyledWrapper {...props}>{getStatusIcon(active, t)}</StyledWrapper>
}

StatusIcon.propTypes = {
  active: bool,
}

StatusIcon.defaultProps = {
  active: false,
}

export default StatusIcon
