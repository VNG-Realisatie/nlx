// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { observer } from 'mobx-react'
import { Route } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import PageTemplate from '../../../components/PageTemplate'
import { useDirectoryStore } from '../../../hooks/use-stores'
import LoadingMessage from '../../../components/LoadingMessage'

import DirectoryDetailPage from '../DirectoryDetailPage'
import DirectoryServiceCount from './components/DirectoryServiceCount'
import DirectoryPageView from './components/DirectoryPageView'

const DirectoryPage = () => {
  const { t } = useTranslation()
  const {
    services,
    selectService,
    isInitiallyFetched,
    error,
  } = useDirectoryStore()

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Directory')}
        description={
          <span data-testid="directory-description">
            {t('List of all available services')}
            {isInitiallyFetched && !error ? (
              <DirectoryServiceCount services={services} />
            ) : null}
          </span>
        }
      />

      {!isInitiallyFetched ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the directory.')}
        </Alert>
      ) : (
        <>
          <DirectoryPageView services={services} />
          <Route
            exact
            path="/directory/:organizationName/:serviceName"
            render={({ match }) => {
              return (
                services.length && (
                  <DirectoryDetailPage service={selectService(match.params)} />
                )
              )
            }}
          />
        </>
      )}
    </PageTemplate>
  )
}

export default observer(DirectoryPage)
