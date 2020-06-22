// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { func } from 'prop-types'
import { Route } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'

import PageTemplate from '../../components/PageTemplate'
import DirectoryRepository from '../../domain/directory-repository'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'
import DirectoryDetailPage from '../DirectoryDetailPage'
import DirectoryServiceCount from './components/DirectoryServiceCount'
import DirectoryPageView from './components/DirectoryPageView'

const DirectoryPage = ({ getDirectoryServices }) => {
  const { t } = useTranslation()
  const { isReady, result: services, error } = usePromise(getDirectoryServices)

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
        <>
          <DirectoryPageView services={services} />
          <Route exact path="/directory/:organizationName/:serviceName">
            <DirectoryDetailPage parentUrl="/directory" />
          </Route>
        </>
      )}
    </PageTemplate>
  )
}

DirectoryPage.propTypes = {
  getDirectoryServices: func,
}

DirectoryPage.defaultProps = {
  getDirectoryServices: DirectoryRepository.getAll,
}

export default DirectoryPage
