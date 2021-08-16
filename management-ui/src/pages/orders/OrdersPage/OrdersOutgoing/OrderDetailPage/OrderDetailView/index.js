// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string, instanceOf } from 'prop-types'
import { observer } from 'mobx-react'
import { SectionGroup } from '../../../../../../components/DetailView'
import Reference from './Reference'
import StartEndDate from './StartEndDate'

const OrderDetailView = ({ order }) => {
  return (
    <SectionGroup>
      <StartEndDate validFrom={order.validFrom} validUntil={order.validUntil} />
      <Reference value={order.reference} />
    </SectionGroup>
  )
}

OrderDetailView.propTypes = {
  order: shape({
    delegatee: string.isRequired,
    reference: string.isRequired,
    validFrom: instanceOf(Date).isRequired,
    validUntil: instanceOf(Date).isRequired,
  }).isRequired,
}

OrderDetailView.defaultProps = {}

export default observer(OrderDetailView)
