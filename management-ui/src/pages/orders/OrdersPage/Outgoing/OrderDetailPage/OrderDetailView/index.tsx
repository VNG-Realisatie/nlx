// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { SectionGroup } from '../../../../../../components/DetailView'
import OutgoingOrderModel from '../../../../../../stores/models/OutgoingOrderModel'
import Service from '../../../../../../types/Service'
import Status from './Status'
import Reference from './Reference'
import StartEndDate from './StartEndDate'
import Services from './Services'

interface OrderDetailViewProps {
  order: OutgoingOrderModel
}

const OrderDetailView: React.FC<OrderDetailViewProps> = ({ order }) => {
  return (
    <>
      <SectionGroup>
        <Status data-testid="status" order={order} />
        <StartEndDate
          data-testid="start-end-date"
          validFrom={order.validFrom}
          validUntil={order.validUntil}
          revokedAt={order.revokedAt}
        />
        <Reference value={order.reference} />
        <Services services={order.services as unknown as Service[]} />
      </SectionGroup>
    </>
  )
}

export default observer(OrderDetailView)
