// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { instanceOf } from 'prop-types'
import { SectionGroup } from '../../../../../../components/DetailView'
import OutgoingOrderModel from '../../../../../../stores/models/OutgoingOrderModel'
import { useDirectoryServiceStore } from '../../../../../../hooks/use-stores'
import Status from './Status'
import Reference from './Reference'
import StartEndDate from './StartEndDate'
import Services from './Services'

const OrderDetailView = ({ order }) => {
  const directoryServiceStore = useDirectoryServiceStore()
  const services = order.accessProofs
    .map((accessProof) => {
      return directoryServiceStore.getService(
        accessProof.organization.serialNumber,
        accessProof.serviceName,
      )
    })
    .filter((service) => !!service)
    .filter((service, index, arr) => {
      return arr.indexOf(service) === index
    })

  return (
    <SectionGroup>
      <Status data-testid="status" order={order} />
      <StartEndDate
        data-testid="start-end-date"
        validFrom={order.validFrom}
        validUntil={order.validUntil}
        revokedAt={order.revokedAt}
      />
      <Reference value={order.reference} />
      <Services services={services} />
    </SectionGroup>
  )
}

OrderDetailView.propTypes = {
  order: instanceOf(OutgoingOrderModel),
}

export default observer(OrderDetailView)
