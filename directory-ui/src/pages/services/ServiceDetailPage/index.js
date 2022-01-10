// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { useLocation, useNavigate, useParams } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import getServices from '../ServicesPage/get-services'
import DirectoryDetailView from './components/DirectoryDetailView'
import DrawerHeader from './components/DrawerHeader'
import { StyledDrawer } from './index.styles'

const ServiceDetailPage = () => {
  const navigate = useNavigate()
  const { serviceName } = useParams()
  const location = useLocation()
  const [service, setService] = useState()

  useEffect(() => {
    const loadService = async () => {
      const services = await getServices()

      const activeService = services.find(
        (service) => service.name === serviceName,
      )

      setService(activeService)
    }
    loadService()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  const navigateToParentUrl = () => {
    const urlParams = new URLSearchParams(location.search)
    const query = urlParams.get('q')
    if (query) {
      navigate(`/?q=${encodeURIComponent(query)}`)
    } else {
      navigate(`/`)
    }
  }

  return (
    <StyledDrawer
      id="directoryDetail"
      noMask
      afterHide={navigateToParentUrl}
      closeHandler={navigateToParentUrl}
    >
      {service ? (
        <DrawerHeader service={service} />
      ) : (
        <Drawer.Header
          as="header"
          title={serviceName}
          closeButtonLabel="Close"
        />
      )}

      <Drawer.Content>
        {service ? (
          <DirectoryDetailView service={service} />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {`Kan de service '${serviceName}' niet vinden.`}
          </Alert>
        )}
      </Drawer.Content>
    </StyledDrawer>
  )
}

export default ServiceDetailPage
