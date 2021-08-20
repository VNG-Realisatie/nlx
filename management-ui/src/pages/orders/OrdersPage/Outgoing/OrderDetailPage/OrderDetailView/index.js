// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { instanceOf, func } from 'prop-types'
import { observer } from 'mobx-react'
import { SectionGroup } from '../../../../../../components/DetailView'
import { useConfirmationModal } from '../../../../../../components/ConfirmationModal'
import OutgoingOrderModel from '../../../../../../stores/models/OutgoingOrderModel'
import Status from './Status'
import Reference from './Reference'
import StartEndDate from './StartEndDate'
import Services from './Services'

const OrderDetailView = ({ order, revokeHandler }) => {
  const { t } = useTranslation()

  const [ConfirmRevokeModal, confirmRevoke] = useConfirmationModal({
    okText: t('Revoke'),
    children: <p>{t('Do you want to revoke the order?')}</p>,
  })

  const handleRevoke = async () => {
    if (await confirmRevoke()) {
      revokeHandler(order)
    }
  }

  return (
    <>
      <SectionGroup>
        <Status
          data-testid="status"
          order={order}
          revokeHandler={handleRevoke}
        />
        <StartEndDate
          data-testid="start-end-date"
          validFrom={order.validFrom}
          validUntil={order.validUntil}
          revokedAt={order.revokedAt}
        />
        <Reference value={order.reference} />
        <Services services={order.services} />
      </SectionGroup>

      <ConfirmRevokeModal />
    </>
  )
}

OrderDetailView.propTypes = {
  order: instanceOf(OutgoingOrderModel).isRequired,
  revokeHandler: func.isRequired,
}

OrderDetailView.defaultProps = {}

export default observer(OrderDetailView)
