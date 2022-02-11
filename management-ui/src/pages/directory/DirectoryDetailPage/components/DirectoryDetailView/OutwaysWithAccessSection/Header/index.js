// Copyright © VNG Realisatie 2022
// Licensed under the EUPL
//

import React from 'react'
import { string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { IconKey } from '../../../../../../../icons'
import { StyledLabel, StyledDetailHeading } from './index.styles'

const Header = ({ label }) => {
  const { t } = useTranslation()

  return (
    <StyledDetailHeading>
      <IconKey />
      {t('Outways with access')}
      <StyledLabel>{label}</StyledLabel>
    </StyledDetailHeading>
  )
}

Header.propTypes = {
  label: string,
}

export default Header
