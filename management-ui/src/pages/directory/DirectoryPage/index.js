// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { Route, Routes, useParams } from 'react-router-dom'
import { Alert, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import LoadingMessage from '../../../components/LoadingMessage'
import PageTemplate from '../../../components/PageTemplate'
import { useDirectoryServiceStore } from '../../../hooks/use-stores'
import DirectoryDetailPage from '../DirectoryDetailPage'
import EnvironmentRepository from '../../../domain/environment-repository'
import DirectoryServiceCount from './components/DirectoryServiceCount'
import DirectoryPageView from './components/DirectoryPageView'

const DirectoryPage = () => {
  const [subjectSerialNumber, setSubjectSerialNumber] = useState('')
  const { name } = useParams()
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const directoryServiceStore = useDirectoryServiceStore()
  const [isLoaded, setIsLoaded] = useState(false)

  useEffect(() => {
    async function fetchData() {
      const env = await EnvironmentRepository.getCurrent()
      setSubjectSerialNumber(env.organizationSerialNumber)

      try {
        await directoryServiceStore.syncAllOutgoingAccessRequests()
      } catch (error) {
        showToast({
          title: t('Failed to synchronize access states'),
          variant: 'error',
        })
      }

      await directoryServiceStore.fetchAll()

      setIsLoaded(true)
    }

    fetchData()
  }, []) // eslint-disable-line react-hooks/exhaustive-deps

  const DirectoryCount = () => {
    if (isLoaded && !directoryServiceStore.error) {
      return <DirectoryServiceCount services={directoryServiceStore.services} />
    }
    return null
  }

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Directory')}
        description={
          <span data-testid="directory-description">
            {t('List of all available services')}
            <DirectoryCount />
          </span>
        }
      />

      {!isLoaded ? (
        <LoadingMessage />
      ) : directoryServiceStore.error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the directory')}
        </Alert>
      ) : (
        <>
          <DirectoryPageView
            managementSubjectSerialNumber={subjectSerialNumber}
            services={directoryServiceStore.services}
            selectedServiceName={name || ''}
          />

          <Routes>
            <Route
              path=":organizationSerialNumber/:serviceName"
              element={<DirectoryDetailPage />}
            />
          </Routes>
        </>
      )}
    </PageTemplate>
  )
}

export default observer(DirectoryPage)
