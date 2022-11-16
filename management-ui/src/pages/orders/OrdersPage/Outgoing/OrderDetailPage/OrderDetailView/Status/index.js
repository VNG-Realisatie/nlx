// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import StatusIcon from '../../../../StatusIcon'
import OutgoingOrderModel from '../../../../../../../stores/models/OutgoingOrderModel'
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
  order: instanceOf(OutgoingOrderModel),
}

export default observer(Status)
