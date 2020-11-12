// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect } from 'react'
import { shape, string } from 'prop-types'
import { observer } from 'mobx-react'
import { useParams, useHistory } from 'react-router-dom'
import {
  Alert,
  Drawer,
  StackedDrawer,
  withDrawerStack,
  useDrawerStack,
} from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import { directoryServicePropTypes } from '../../../models/DirectoryServiceModel'
import DirectoryDetailView from './components/DirectoryDetailView'
import DrawerHeader from './components/DrawerHeader'

const DirectoryDetailPage = ({ service, parentUrl }) => {
  const { t } = useTranslation()
  const history = useHistory()
  const { organizationName, serviceName } = useParams()
  const { showDrawer } = useDrawerStack()

  const afterHideHandler = () => {
    history.push(parentUrl)
  }

  useEffect(() => {
    showDrawer('directoryDetail')
  }, [service]) // eslint-disable-line react-hooks/exhaustive-deps

  return (
    <StackedDrawer id="directoryDetail" noMask afterHide={afterHideHandler}>
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
          <DirectoryDetailView service={service} />
        ) : (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the service', {
              name: `${organizationName}/${serviceName}`,
            })}
          </Alert>
        )}
      </Drawer.Content>
    </StackedDrawer>
  )
}

DirectoryDetailPage.propTypes = {
  service: shape(directoryServicePropTypes),
  parentUrl: string,
}

DirectoryDetailPage.defaultProps = {
  parentUrl: '/directory',
}

export default observer(withDrawerStack(DirectoryDetailPage))
