// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useContext, useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { Route, Routes, useMatch, useParams } from 'react-router-dom'
import { Alert, ToasterContext } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import LoadingMessage from '../../../components/LoadingMessage'
import PageTemplate from '../../../components/PageTemplate'
import { useDirectoryServiceStore } from '../../../hooks/use-stores'
import DirectoryDetailPage from '../DirectoryDetailPage'
import EnvironmentRepository from '../../../domain/environment-repository'
import usePolling from '../../../hooks/use-polling'
import DirectoryServiceCount from './components/DirectoryServiceCount'
import DirectoryPageView from './components/DirectoryPageView'

const DirectoryPage = () => {
  const [subjectSerialNumber, setSubjectSerialNumber] = useState('')
  const { name } = useParams()
  const { t } = useTranslation()
  const { showToast } = useContext(ToasterContext)
  const directoryServiceStore = useDirectoryServiceStore()
  const [isLoaded, setIsLoaded] = useState(false)
  const [isPollRequestRunning, setIsPollRequestRunning] = useState(false)
  const directoryDetailPageMatch = useMatch(
    '/directory/:organizationSerialNumber/:serviceName',
  )

  usePolling(() => {
    async function fetchData() {
      if (document.hidden) {
        return
      }

      if (directoryDetailPageMatch) {
        return
      }

      if (directoryServiceStore.services.length === 0) {
        return
      }

      if (isPollRequestRunning) {
        return
      }

      setIsPollRequestRunning(true)

      try {
        await directoryServiceStore.syncAllOutgoingAccessRequests()
      } catch (error) {
        let message = ''

        if (error.response) {
          const text = await error.response.text()
          const textAsJson = JSON.parse(text)
          message = textAsJson.details[0].metadata['organizations']
        }

        showToast({
          title: t('Failed to synchronize access with:'),
          body: message,
          variant: 'error',
        })
      }

      setIsPollRequestRunning(false)
    }

    fetchData()
  })

  useEffect(() => {
    async function fetchData() {
      await directoryServiceStore.fetchAll()

      if (directoryServiceStore.services.length !== 0) {
        directoryServiceStore.syncAllOutgoingAccessRequests()
      }

      const env = await EnvironmentRepository.getCurrent()

      setSubjectSerialNumber(env.organizationSerialNumber)
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
