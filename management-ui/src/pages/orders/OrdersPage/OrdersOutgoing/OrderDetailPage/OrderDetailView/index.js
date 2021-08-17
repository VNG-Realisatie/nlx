// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { shape, string, instanceOf, func } from 'prop-types'
import { observer } from 'mobx-react'
import { SectionGroup } from '../../../../../../components/DetailView'
import { useConfirmationModal } from '../../../../../../components/ConfirmationModal'
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
        <Status order={order} revokeHandler={handleRevoke} />
        <StartEndDate
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
  order: shape({
    delegatee: string.isRequired,
    reference: string.isRequired,
    validFrom: instanceOf(Date).isRequired,
    validUntil: instanceOf(Date).isRequired,
  }).isRequired,
  revokeHandler: func.isRequired,
}

OrderDetailView.defaultProps = {}

export default observer(OrderDetailView)
