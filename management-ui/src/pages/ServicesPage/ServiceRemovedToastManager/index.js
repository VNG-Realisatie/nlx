// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { useTranslation } from 'react-i18next'
import { useContext, useEffect } from 'react'
import { ToasterContext } from '@commonground/design-system'
import { useHistory, useLocation, useRouteMatch } from 'react-router-dom'

const ServiceRemovedToastManager = () => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const serviceDetailPageMatch = useRouteMatch('/services/:serviceName')
  const location = useLocation()
  const history = useHistory()

  useEffect(() => {
    if (!serviceDetailPageMatch) {
      return
    }

    const searchParams = new URLSearchParams(location.search)
    if (searchParams.get('removed') !== 'true') {
      return
    }

    const { serviceName } = serviceDetailPageMatch.params
    showToast({
      title: serviceName,
      body: t('The service has been removed.'),
      variant: 'success',
    })
    history.replace('/services')
  }, [serviceDetailPageMatch, location.search, showToast, t, history])

  return null
}

export default ServiceRemovedToastManager
