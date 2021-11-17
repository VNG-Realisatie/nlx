// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React, { useEffect, useState } from 'react'
import { observer } from 'mobx-react'
import { Route, useParams } from 'react-router-dom'
import { Alert } from '@commonground/design-system'
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
  const { t } = useTranslation()
  const { services, getService, isInitiallyFetched, error } =
    useDirectoryServiceStore()
  const { name } = useParams<{ name: string }>()

  useEffect(() => {
    const loadEnv = async () => {
      const env = await EnvironmentRepository.getCurrent()
      const organizationSubjectSerialNumber = env.organizationSerialNumber
      setSubjectSerialNumber(organizationSubjectSerialNumber)
    }
    loadEnv().catch(console.warn)
  })

  const DirectoryCount = () => {
    if (isInitiallyFetched && !error) {
      return <DirectoryServiceCount services={services} />
    }
    return null
  }

  const MainContent = () => {
    if (!isInitiallyFetched || !subjectSerialNumber) {
      return <LoadingMessage />
    }
    if (error) {
      return (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the directory')}
        </Alert>
      )
    }
    return (
      <>
        <DirectoryPageView
          managementSubjectSerialNumber={subjectSerialNumber}
          services={services}
          selectedServiceName={name}
        />
        <Route
          exact
          path="/directory/:organizationSerialNumber/:serviceName"
          render={({ match }) => {
            const service = getService(
              match.params.organizationSerialNumber,
              match.params.serviceName,
            )

            if (service) {
              service.fetch()
            }

            return services.length && <DirectoryDetailPage service={service} />
          }}
        />
      </>
    )
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
      <MainContent />
    </PageTemplate>
  )
}

export default observer(DirectoryPage)
