// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { instanceOf } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { IconTimer } from '../../../../../../../icons'
import { StyledContainer, StyledLabel } from './index.styles'

const StartEndDate = ({ validFrom, validUntil, revokedAt }) => {
  const { t } = useTranslation()
  return (
    <StyledContainer>
      <IconTimer />
      {revokedAt ? (
        t('Revoked on date', { date: revokedAt })
      ) : (
        <>
          {t('Valid until date', { date: validUntil })}
          <StyledLabel>{t('Since date', { date: validFrom })}</StyledLabel>
        </>
      )}
    </StyledContainer>
  )
}

StartEndDate.propTypes = {
  validFrom: instanceOf(Date).isRequired,
  validUntil: instanceOf(Date).isRequired,
  revokedAt: instanceOf(Date),
}

export default StartEndDate
