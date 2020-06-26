// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { useTranslation } from 'react-i18next'
import { useContext, useEffect, useState } from 'react'
import { ToasterContext } from '@commonground/design-system'
import { useHistory, useLocation, useRouteMatch } from 'react-router-dom'

const ServiceEditedToastManager = () => {
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const serviceDetailPageMatch = useRouteMatch('/services/:serviceName')
  const [isToastAdded, setIsToastAdded] = useState(false)
  const location = useLocation()
  const history = useHistory()

  useEffect(() => {
    if (!serviceDetailPageMatch) {
      return
    }

    if (isToastAdded) {
      return
    }

    const searchParams = new URLSearchParams(location.search)
    if (searchParams.get('edited') !== 'true') {
      return
    }

    const { serviceName, url } = serviceDetailPageMatch.params
    setIsToastAdded(true)
    showToast({
      title: serviceName,
      body: t('The service has been updated.'),
      variant: 'success',
    })
    history.replace(url)
  }, [
    serviceDetailPageMatch,
    isToastAdded,
    location.search,
    showToast,
    t,
    history,
  ])

  return null
}

export default ServiceEditedToastManager
