// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { Alert, Button } from '@commonground/design-system'
import { Link, Route, useParams } from 'react-router-dom'
import { observer } from 'mobx-react'
import usePolling from '../../../hooks/use-polling'
import PageTemplate from '../../../components/PageTemplate'
import ServiceDetailPage from '../ServiceDetailPage'
import LoadingMessage from '../../../components/LoadingMessage'
import { useServiceStore } from '../../../hooks/use-stores'
import { IconPlus } from '../../../icons'
import ServiceToastManager from './ServiceToastManager'
import ServiceCount from './ServiceCount'
import ServicesPageView from './ServicesPageView'
import { StyledActionsBar } from './index.styles'

const ServicesPage = () => {
  const { t } = useTranslation()
  const {
    isInitiallyFetched,
    services,
    error,
    getService,
    fetchStats,
  } = useServiceStore()
  const { name } = useParams()

  usePolling(fetchStats)

  return (
    <PageTemplate>
      <ServiceToastManager />

      <PageTemplate.Header title={t('Services')} />

      <StyledActionsBar>
        <ServiceCount
          count={services ? services.length : 0}
          data-testid="service-count"
        />
        <Button
          as={Link}
          to="/services/add-service"
          aria-label={t('Add service')}
        >
          <IconPlus inline />
          {t('Add service')}
        </Button>
      </StyledActionsBar>

      {!isInitiallyFetched ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the services')}
        </Alert>
      ) : (
        <>
          <ServicesPageView services={services} selectedServiceName={name} />
          <Route
            path="/services/:name"
            render={({ match }) => {
              const service = getService(match.params.name)

              if (service) {
                service.fetch()
              }

              return (
                services.length && (
                  <ServiceDetailPage parentUrl="/services" service={service} />
                )
              )
            }}
          />
        </>
      )}
    </PageTemplate>
  )
}

export default observer(ServicesPage)
