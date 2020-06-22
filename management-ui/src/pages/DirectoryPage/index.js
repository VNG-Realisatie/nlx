// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { createContext, useState } from 'react'
import { func } from 'prop-types'
import { Route } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import PageTemplate from '../../components/PageTemplate'
import usePromise from '../../hooks/use-promise'
import DirectoryRepository from '../../domain/directory-repository'
import AccessRequestRepository from '../../domain/access-request-repository'
import LoadingMessage from '../../components/LoadingMessage'

import DirectoryDetailPage from '../DirectoryDetailPage'
import DirectoryServiceCount from './components/DirectoryServiceCount'
import DirectoryPageView from './components/DirectoryPageView'

export const AccessRequestContext = createContext()

const DEFAULT_REQUEST_SENT_STATE = {
  organizationName: '',
  serviceName: '',
}

const DirectoryPage = ({ getDirectoryServices, requestAccess }) => {
  const { t } = useTranslation()
  const { isReady, result: services, error, reload } = usePromise(
    getDirectoryServices,
  )

  const [requestSentTo, setRequestSentTo] = useState(DEFAULT_REQUEST_SENT_STATE)

  const handleRequestAccess = async ({ organizationName, serviceName }) => {
    const confirmed = window.confirm(
      t('The request will be sent to', { name: organizationName }),
    )

    if (confirmed) {
      setRequestSentTo({ organizationName, serviceName })

      try {
        await requestAccess(organizationName, serviceName)
        reload()
      } catch (e) {
        console.error(e)
        setRequestSentTo(DEFAULT_REQUEST_SENT_STATE)
      }
    }
  }

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Directory')}
        description={
          <span data-testid="directory-description">
            {t('List of all available services')}
            {isReady && !error ? (
              <DirectoryServiceCount services={services} />
            ) : null}
          </span>
        }
      />

      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the directory.')}
        </Alert>
      ) : (
        <AccessRequestContext.Provider
          value={{ requestSentTo, handleRequestAccess }}
        >
          <DirectoryPageView
            services={services}
            handleRequestAccess={handleRequestAccess}
          />
          <Route exact path="/directory/:organizationName/:serviceName">
            <DirectoryDetailPage
              parentUrl="/directory"
              // isReloaded={isReloaded}
            />
          </Route>
        </AccessRequestContext.Provider>
      )}
    </PageTemplate>
  )
}

DirectoryPage.propTypes = {
  getDirectoryServices: func,
  requestAccess: func,
}

DirectoryPage.defaultProps = {
  getDirectoryServices: DirectoryRepository.getAll,
  requestAccess: AccessRequestRepository.requestAccess,
}

export default DirectoryPage
