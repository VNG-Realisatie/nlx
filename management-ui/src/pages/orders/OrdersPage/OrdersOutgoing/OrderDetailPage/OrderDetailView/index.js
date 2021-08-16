// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { SectionGroup } from '../../../../../../components/DetailView'
import Reference from './Reference'

const OrderDetailView = ({ order }) => {
  return (
    <SectionGroup>
      <Reference value={order.reference} />
    </SectionGroup>
  )
}

OrderDetailView.propTypes = {
  order: shape({
    delegatee: string.isRequired,
    reference: string.isRequired,
  }).isRequired,
}

OrderDetailView.defaultProps = {}

export default observer(OrderDetailView)
