// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useContext, useEffect } from 'react'
import { useTranslation } from 'react-i18next'
import { ToasterContext } from '@commonground/design-system'
import { useNavigate, useLocation, useMatch } from 'react-router-dom'

import serviceActions from '../serviceActions'

const getToastMessageForAction = (action, t) => {
  switch (action) {
    case serviceActions.ADDED:
      return t('The service has been added')
    case serviceActions.EDITED:
      return t('The service has been updated')
    case serviceActions.REMOVED:
      return t('The service has been removed')
    default:
      console.warn(
        `can not determine toast message, unknown action '${action}'`,
      )
      return ''
  }
}

const ServiceToastManager = () => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const serviceDetailPageMatch = useMatch('/services/:serviceName')
  const location = useLocation()
  const navigate = useNavigate()

  useEffect(() => {
    if (!serviceDetailPageMatch) return

    const searchParams = new URLSearchParams(location.search)
    const lastAction = searchParams.get('lastAction')
    if (!lastAction) return

    const { serviceName } = serviceDetailPageMatch.params
    const url = serviceDetailPageMatch.pathname

    showToast({
      title: serviceName,
      body: getToastMessageForAction(lastAction, t),
      variant: 'success',
      delay: lastAction === serviceActions.REMOVED ? 0 : 250,
    })

    navigate(lastAction === serviceActions.REMOVED ? '/services' : url, {
      replace: true,
    })
  }, [location.search, serviceDetailPageMatch, showToast, t, navigate])

  return null
}

export default ServiceToastManager
