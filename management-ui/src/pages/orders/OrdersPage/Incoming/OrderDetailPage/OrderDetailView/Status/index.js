// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { observer } from 'mobx-react'
import { object } from 'prop-types'
import { useTranslation } from 'react-i18next'
import StatusIcon from '../../../../StatusIcon'
import { StyledContainer, StateDetail } from './index.styles'

const Status = ({ order, ...props }) => {
  const { t } = useTranslation()

  return (
    <StyledContainer {...props}>
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
        </>
      )}
    </StyledContainer>
  )
}

Status.propTypes = {
  order: object.isRequired,
}

export default observer(Status)
