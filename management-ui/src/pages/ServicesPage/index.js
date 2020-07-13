// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { func } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert, Button } from '@commonground/design-system'
import { Link, Route } from 'react-router-dom'

import PageTemplate from '../../components/PageTemplate'
import usePromise from '../../hooks/use-promise'
import ServiceRepository from '../../domain/service-repository'
import ServiceDetailPage from '../ServiceDetailPage'

import LoadingMessage from '../../components/LoadingMessage'

import ServiceToastManager from './ServiceToastManager'
import ServiceCount from './ServiceCount'
import ServicesPageView from './ServicesPageView'
import { StyledActionsBar, StyledIconPlus } from './index.styles'

const ServicesPage = ({ getServices }) => {
  const { t } = useTranslation()
  const { isReady, error, result: services, reload } = usePromise(getServices)

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
          <StyledIconPlus />
          {t('Add service')}
        </Button>
      </StyledActionsBar>

      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the services.')}
        </Alert>
      ) : (
        <ServicesPageView services={services} />
      )}

      <Route path="/services/:name">
        <ServiceDetailPage parentUrl="/services" refreshHandler={reload} />
      </Route>
    </PageTemplate>
  )
}

ServicesPage.propTypes = {
  getServices: func,
}

ServicesPage.defaultProps = {
  getServices: ServiceRepository.getAll,
}

export default ServicesPage
