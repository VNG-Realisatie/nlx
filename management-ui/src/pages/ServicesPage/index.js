// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { array, func, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert, Button } from '@commonground/design-system'
import { Link, Route, useLocation } from 'react-router-dom'
import PageTemplate from '../../components/PageTemplate'
import usePromise from '../../hooks/use-promise'
import ServiceRepository from '../../domain/service-repository'
import ServiceDetailPage from '../ServiceDetailPage'
import Table from './Table'
import AuthorizationMode from './AuthorizationMode'
import ServiceCount from './ServiceCount'
import {
  StyledActionsBar,
  StyledIconPlus,
  StyledLoadingMessage,
  StyledNoServicesMessage,
} from './index.styles'
import Spinner from './Spinner'

const ServiceRow = ({ name, authorizations, mode, ...props }) => {
  const location = useLocation()

  return (
    <Table.Tr to={`${location.pathname}/${name}`} name={name} {...props}>
      <Table.Td>{name}</Table.Td>
      <Table.Td>
        <AuthorizationMode authorizations={authorizations} mode={mode} />
      </Table.Td>
    </Table.Tr>
  )
}

ServiceRow.propTypes = {
  name: string.isRequired,
  authorizations: array,
  mode: string.isRequired,
}

const ServicesPage = ({ getServices }) => {
  const { t } = useTranslation()
  const { loading, error, result } = usePromise(getServices)

  return (
    <PageTemplate>
      <PageTemplate.Header title={t('Services')} />

      <StyledActionsBar>
        <ServiceCount
          count={result ? result.length : 0}
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

      {loading ? (
        <StyledLoadingMessage role="progressbar">
          <Spinner /> {t('Loading…')}
        </StyledLoadingMessage>
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the services.')}
        </Alert>
      ) : result != null && result.length === 0 ? (
        <StyledNoServicesMessage>
          {t('There are no services yet.')}
        </StyledNoServicesMessage>
      ) : result ? (
        <Table withLinks data-testid="services-list" role="grid">
          <thead>
            <Table.TrHead>
              <Table.Th>{t('Name')}</Table.Th>
              <Table.Th>{t('Authorization')}</Table.Th>
            </Table.TrHead>
          </thead>
          <tbody>
            {result.map((service, i) => (
              <ServiceRow
                name={service.name}
                authorizations={service.authorizationSettings.authorizations}
                mode={service.authorizationSettings.mode}
                key={i}
              />
            ))}
          </tbody>
        </Table>
      ) : null}
      <Route path="/services/:name">
        <ServiceDetailPage parentUrl="/services" />
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
