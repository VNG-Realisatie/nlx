// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { observer } from 'mobx-react'
import { object, func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Button } from '@commonground/design-system'
import StatusIcon from '../../../../StatusIcon'
import { StyledContainer, StateDetail } from './index.styles'

const Status = ({ order, revokeHandler }) => {
  const { t } = useTranslation()

  return (
    <StyledContainer>
      {order.revokedAt ? (
        <>
          <StatusIcon active={false} />
          <StateDetail>
            <span>{t('Order is revoked')}</span>
          </StateDetail>
        </>
      ) : order.validFrom > new Date() ? (
        <>
          <StatusIcon active={false} />
          <StateDetail>
            <span>{t('Order is not yet active')}</span>
          </StateDetail>
          <Button onClick={revokeHandler} aria-label={t('Revoke')}>
            {t('Revoke')}
          </Button>
        </>
      ) : order.validUntil < new Date() ? (
        <>
          <StatusIcon active={false} />
          <StateDetail>
            <span>{t('Order is expired')}</span>
          </StateDetail>
        </>
      ) : (
        <>
          <StatusIcon active />
          <StateDetail>
            <span>{t('Order is active')}</span>
          </StateDetail>
          <Button onClick={revokeHandler} aria-label={t('Revoke')}>
            {t('Revoke')}
          </Button>
        </>
      )}
    </StyledContainer>
  )
}

Status.propTypes = {
  order: object.isRequired,
  revokeHandler: func.isRequired,
}

export default observer(Status)
