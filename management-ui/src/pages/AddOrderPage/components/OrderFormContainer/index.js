// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//

import React, { useEffect } from 'react'
import { observer } from 'mobx-react'
import { func } from 'prop-types'
import OrderForm from '../OrderForm'
import useStores from '../../../../hooks/use-stores'

const OrderFormContainer = ({ onSubmitHandler }) => {
  const { directoryServicesStore } = useStores()

  useEffect(() => {
    directoryServicesStore.fetchAll()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const serviceNames = directoryServicesStore.servicesWithAccess.map(
    (service) => ({
      service: service.serviceName,
      organization: service.organizationName,
    }),
  )

  return <OrderForm services={serviceNames} onSubmitHandler={onSubmitHandler} />
}

OrderFormContainer.propTypes = {
  onSubmitHandler: func.isRequired,
}

export default observer(OrderFormContainer)
