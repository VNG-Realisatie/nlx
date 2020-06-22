// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useState } from 'react'
import { func, string } from 'prop-types'
import { useParams, useHistory } from 'react-router-dom'
import { Alert, Drawer } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import DirectoryRepository from '../../domain/directory-repository'
import AccessRequestRepository from '../../domain/access-request-repository'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'
import DirectoryDetailView from './components/DirectoryDetailView'
import DrawerHeader from './components/DrawerHeader'

const DirectoryDetailPage = ({ getService, requestAccess, parentUrl }) => {
  const { t } = useTranslation()
  const history = useHistory()
  const { organizationName, serviceName } = useParams()
  const [isAccessRequested, setAccessRequested] = useState(false)

  const { isReady, error, result: service } = usePromise(
    getService,
    organizationName,
    serviceName,
  )

  const close = () => history.push(parentUrl)

  // Concious duplication with DirectoryPage - until confirm box gets replaced by modal
  const handleRequestAccess = async () => {
    const confirmed = window.confirm(
      t('The request will be sent to', { name: organizationName }),
    )

    if (confirmed) {
      try {
        setAccessRequested(true)
        await requestAccess(organizationName, serviceName)
      } catch (e) {
        console.error(e)
      }
    }
  }

  return (
    <Drawer noMask closeHandler={close}>
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
            onRequestAccess={handleRequestAccess}
            isAccessRequested={isAccessRequested}
          />
        ) : null}
      </Drawer.Content>
    </Drawer>
  )
}

DirectoryDetailPage.propTypes = {
  getService: func,
  requestAccess: func,
  parentUrl: string,
}

DirectoryDetailPage.defaultProps = {
  getService: DirectoryRepository.getByName,
  requestAccess: AccessRequestRepository.requestAccess,
  parentUrl: '/directory',
}

export default DirectoryDetailPage
