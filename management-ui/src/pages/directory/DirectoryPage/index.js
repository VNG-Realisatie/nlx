// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React, { createContext, useState, useEffect } from 'react'
import { func } from 'prop-types'
import { observer } from 'mobx-react'
import { Route } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import PageTemplate from '../../../components/PageTemplate'
import { useDirectoryStore } from '../../../hooks/use-stores'
import AccessRequestRepository from '../../../domain/access-request-repository'
import LoadingMessage from '../../../components/LoadingMessage'

import DirectoryDetailPage from '../DirectoryDetailPage'
import DirectoryServiceCount from './components/DirectoryServiceCount'
import DirectoryPageView from './components/DirectoryPageView'

export const AccessRequestContext = createContext()

const DEFAULT_REQUEST_SENT_STATE = {
  organizationName: '',
  serviceName: '',
}

const DirectoryPage = ({ requestAccess }) => {
  const { t } = useTranslation()
  const { getServices, services, isLoading, error } = useDirectoryStore()

  const [requestSentTo, setRequestSentTo] = useState(DEFAULT_REQUEST_SENT_STATE)

  useEffect(() => {
    getServices()
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [])

  const handleRequestAccess = async ({ organizationName, serviceName }) => {
    const confirmed = window.confirm(
      t('The request will be sent to', { name: organizationName }),
    )

    if (confirmed) {
      setRequestSentTo({ organizationName, serviceName })

      try {
        await requestAccess({ organizationName, serviceName })
        setRequestSentTo(DEFAULT_REQUEST_SENT_STATE)
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
            {!isLoading && !error ? (
              <DirectoryServiceCount services={services} />
            ) : null}
          </span>
        }
      />

      {isLoading ? (
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
            <DirectoryDetailPage parentUrl="/directory" />
          </Route>
        </AccessRequestContext.Provider>
      )}
    </PageTemplate>
  )
}

DirectoryPage.propTypes = {
  requestAccess: func,
}

DirectoryPage.defaultProps = {
  requestAccess: AccessRequestRepository.requestAccess,
}

export default observer(DirectoryPage)
