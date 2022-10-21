// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { useTranslation } from 'react-i18next'
import { instanceOf } from 'prop-types'
import StatusIcon from '../../../../StatusIcon'
import OutgoingOrderModel from '../../../../../../../stores/models/OutgoingOrderModel'
import { StyledContainer, StateDetail } from './index.styles'

const Status = ({ order, ...props }) => {
  const { t } = useTranslation()

  const Content = () => {
    if (order.revokedAt) {
      // Revoked order
      return (
        <>
          <StatusIcon active={false} />
          <StateDetail>
            <span>{t('Order is revoked')}</span>
          </StateDetail>
        </>
      )
    }

    if (order.validFrom > new Date()) {
      // Order that is not yet active
      return (
        <>
          <StatusIcon active={false} />
          <StateDetail>
            <span>{t('Order is not yet active')}</span>
          </StateDetail>
        </>
      )
    }

    if (order.validUntil < new Date()) {
      // Expired order
      return (
        <>
          <StatusIcon active={false} />
          <StateDetail>
            <span>{t('Order is expired')}</span>
          </StateDetail>
        </>
      )
    }

    // Active order
    return (
      <>
        <StatusIcon active />
        <StateDetail>
          <span>{t('Order is active')}</span>
        </StateDetail>
      </>
    )
  }

  return (
    <StyledContainer {...props}>
      <Content />
    </StyledContainer>
  )
}

Status.propTypes = {
  order: instanceOf(OutgoingOrderModel),
}

export default Status
