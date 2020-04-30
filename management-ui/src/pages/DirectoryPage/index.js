// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'

import DirectoryRepository from '../../domain/directory-repository'
import usePromise from '../../hooks/use-promise'
import LoadingMessage from '../../components/LoadingMessage'
import DirectoryServiceCount from './DirectoryServiceCount'
import DirectoryServices from './DirectoryServices'

const DirectoryPage = ({ getDirectoryServices }) => {
  const { t } = useTranslation()
  const { isReady, result, error } = usePromise(getDirectoryServices)

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Directory')}
        description={
          <span data-testid="directory-description">
            {t('List of all available services')}
            {isReady && !error ? (
              <DirectoryServiceCount directoryServices={() => result} />
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
        <DirectoryServices directoryServices={() => result} />
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
