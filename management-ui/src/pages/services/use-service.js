// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useEffect, useState } from 'react'
import { useServicesStore } from '../../hooks/use-stores'

const useService = (name) => {
  const [service, setService] = useState(null)
  const [error, setError] = useState('')
  const [isReady, setIsReady] = useState(false)

  const {
    isReady: servicesIsReady,
    services,
    error: servicesError,
    selectService,
    fetchServices,
  } = useServicesStore()

  useEffect(() => {
    setError(servicesError)
  }, [servicesError])

  useEffect(() => {
    if (!isReady && servicesIsReady) {
      const selectedService = selectService(name)
      if (selectedService) {
        selectedService.fetch()
        setService(selectedService)
      } else {
        setError('Service not found')
      }

      setIsReady(true)
    }
  }, [services]) // eslint-disable-line react-hooks/exhaustive-deps

  if (!servicesIsReady) {
    fetchServices()
  }

  return [service, error, isReady]
}

export default useService
