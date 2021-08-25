// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { SectionGroup } from '../../../../../../components/DetailView'
import IncomingOrderModel from '../../../../../../stores/models/IncomingOrderModel'
import Status from './Status'
import Reference from './Reference'
import StartEndDate from './StartEndDate'
import Services from './Services'

const OrderDetailView = ({ order }) => (
  <SectionGroup>
    <Status data-testid="status" order={order} />
    <StartEndDate
      data-testid="start-end-date"
      validFrom={order.validFrom}
      validUntil={order.validUntil}
      revokedAt={order.revokedAt}
    />
    <Reference value={order.reference} />
    <Services services={order.services} />
  </SectionGroup>
)

OrderDetailView.propTypes = {
  order: instanceOf(IncomingOrderModel).isRequired,
}

OrderDetailView.defaultProps = {}

export default observer(OrderDetailView)
