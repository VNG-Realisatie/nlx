// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect } from 'react'
import { observer } from 'mobx-react'
import { instanceOf } from 'prop-types'
import { SectionGroup } from '../../../../../../components/DetailView'
import OutgoingOrderModel from '../../../../../../stores/models/OutgoingOrderModel'
import { useDirectoryServiceStore } from '../../../../../../hooks/use-stores'
import Status from './Status'
import Reference from './Reference'
import StartEndDate from './StartEndDate'
import AccessProofs from './AccessProofs'

const OrderDetailView = ({ order }) => {
  const directoryServiceStore = useDirectoryServiceStore()

  useEffect(() => {
    directoryServiceStore.fetchAll()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

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
      <AccessProofs accessProofs={order.accessProofs} />
    </SectionGroup>
  )
}

OrderDetailView.propTypes = {
  order: instanceOf(OutgoingOrderModel),
}

export default observer(OrderDetailView)
