// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { array, func, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert, Button, Table } from '@commonground/design-system'
import { Link, Route } from 'react-router-dom'

import PageTemplate from '../../components/PageTemplate'
import usePromise from '../../hooks/use-promise'
import ServiceRepository from '../../domain/service-repository'
import ServiceDetailPage from '../ServiceDetailPage'
import EmptyContentMessage from '../../components/EmptyContentMessage'
import LoadingMessage from '../../components/LoadingMessage'
import AuthorizationMode from './AuthorizationMode'
import ServiceCount from './ServiceCount'
import { StyledActionsBar, StyledIconPlus } from './index.styles'

const ServiceRow = ({ name, authorizations, mode, ...props }) => (
  <Table.Tr
    to={`/services/${name}`}
    name={name}
    data-testid="service-row"
    {...props}
  >
    <Table.Td>{name}</Table.Td>
    <Table.Td>
      <AuthorizationMode authorizations={authorizations} mode={mode} />
    </Table.Td>
  </Table.Tr>
)

ServiceRow.propTypes = {
  name: string.isRequired,
  authorizations: array,
  mode: string.isRequired,
}

const ServicesPage = ({ getServices }) => {
  const { t } = useTranslation()
  const { isReady, error, result, reload } = usePromise(getServices)

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

      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the services.')}
        </Alert>
      ) : result != null && result.length === 0 ? (
        <EmptyContentMessage>
          {t('There are no services yet.')}
        </EmptyContentMessage>
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
