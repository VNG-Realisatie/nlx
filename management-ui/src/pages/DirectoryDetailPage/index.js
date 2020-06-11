// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func, string } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import DirectoryRepository from '../../domain/directory-repository'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'
import DirectoryDetailView from './components/DirectoryDetailView'
import DrawerHeader from './components/DrawerHeader'

const DirectoryDetailPage = ({ getService, parentUrl }) => {
  const { organizationName, serviceName } = useParams()
  const { t } = useTranslation()
  const history = useHistory()
  const { isReady, error, result: service } = usePromise(
    getService,
    organizationName,
    serviceName,
  )
  const close = () => history.push(parentUrl)

  const handleRequestAccess = () => console.log('request access')

  return (
    <Drawer noMask closeHandler={close}>
      {service && <DrawerHeader service={service} />}

      <Drawer.Content>
        {!isReady || (!error && !service) ? (
          <LoadingMessage />
        ) : error ? (
          <Alert variant="error" data-testid="error-message">
            {t('Failed to load the service.', {
              name: `${organizationName}/${serviceName}`,
            })}
          </Alert>
        ) : service ? (
          <DirectoryDetailView
            service={service}
            onRequestAccess={handleRequestAccess}
          />
        ) : null}
      </Drawer.Content>
    </Drawer>
  )
}

DirectoryDetailPage.propTypes = {
  getService: func,
  parentUrl: string,
}

DirectoryDetailPage.defaultProps = {
  getService: DirectoryRepository.getByName,
  parentUrl: '/directory',
}

export default DirectoryDetailPage
