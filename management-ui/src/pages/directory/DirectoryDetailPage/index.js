// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { useParams, useNavigate } from 'react-router-dom'
import {
  Alert,
  Drawer,
  StackedDrawer,
  withDrawerStack,
  useDrawerStack,
} from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import {
  useDirectoryServiceStore,
  useOutwayStore,
} from '../../../hooks/use-stores'
import DirectoryDetailView from './components/DirectoryDetailView'
import DrawerHeader from './components/DrawerHeader'

const DirectoryDetailPage = () => {
  const { t } = useTranslation()
  const navigate = useNavigate()
  const { organizationSerialNumber, serviceName } = useParams()
  const { showDrawer } = useDrawerStack()
  const directoryServiceStore = useDirectoryServiceStore()
  const outwayStore = useOutwayStore()
  const [service, setService] = useState(
    directoryServiceStore.getService(organizationSerialNumber, serviceName),
  )

  useEffect(() => {
    const fetch = async () => {
      await outwayStore.fetchAll()

      try {
        const result = await directoryServiceStore.fetch(
          organizationSerialNumber,
          serviceName,
        )
        setService(result)
      } catch {}
    }

    fetch()

    showDrawer('directoryDetail')
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const navigateToParentUrl = () => {
    navigate('/directory')
  }

  return (
    <StackedDrawer id="directoryDetail" noMask afterHide={navigateToParentUrl}>
      {service ? (
        <DrawerHeader service={service} />
      ) : (
        <Drawer.Header
          as="header"
          title={serviceName}
          closeButtonLabel={t('Close')}
        />
      )}

      <Drawer.Content>
        {service ? (
          <DirectoryDetailView
            service={service}
            outways={outwayStore.outways}
          />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the service', {
              name: `${organizationSerialNumber}/${serviceName}`,
            })}
          </Alert>
        )}
      </Drawer.Content>
    </StackedDrawer>
  )
}

export default observer(withDrawerStack(DirectoryDetailPage))
