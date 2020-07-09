// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useContext, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { ToasterContext } from '@commonground/design-system'
import { useHistory, useLocation, useRouteMatch } from 'react-router-dom'

import serviceActions from '../serviceActions'

const toastMessages = {
  [serviceActions.ADDED]: (t) => t('The service has been added.'),
  [serviceActions.EDITED]: (t) => t('The service has been updated.'),
  [serviceActions.REMOVED]: (t) => t('The service has been removed.'),
}

const ServiceToastManager = () => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const serviceDetailPageMatch = useRouteMatch('/services/:serviceName')
  const location = useLocation()
  const history = useHistory()

  useEffect(() => {
    if (!serviceDetailPageMatch) return

    const searchParams = new URLSearchParams(location.search)
    const lastAction = searchParams.get('lastAction')
    if (!lastAction) return

    const { serviceName, url } = serviceDetailPageMatch.params

    showToast({
      title: serviceName,
      body: toastMessages[lastAction](t),
      variant: 'success',
    })

    history.replace(lastAction === serviceActions.REMOVED ? '/services' : url)
  }, [history, location.search, serviceDetailPageMatch, showToast, t])

  return null
}

export default ServiceToastManager
