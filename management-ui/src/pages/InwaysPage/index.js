// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import React from 'react'
import { func, string, number } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Alert } from '@commonground/design-system'
import PageTemplate from '../../components/PageTemplate'
import usePromise from '../../hooks/use-promise'
import Table from '../../components/Table'
import InwayRepository from '../../domain/inway-repository'
import EmptyContentMessage from '../../components/EmptyContentMessage'
import LoadingMessage from '../../components/LoadingMessage'
import { StyledInwayIcon, StyledIconTd } from './index.styles'

const InwayRow = ({
  name,
  hostname,
  selfAddress,
  servicesCount,
  version,
  ...props
}) => (
  <Table.Tr {...props}>
    <StyledIconTd>
      <StyledInwayIcon />
    </StyledIconTd>
    <Table.Td>{name}</Table.Td>
    <Table.Td>{hostname}</Table.Td>
    <Table.Td>{selfAddress}</Table.Td>
    <Table.Td data-testid="services-count">{servicesCount}</Table.Td>
    <Table.Td>{version}</Table.Td>
  </Table.Tr>
)

InwayRow.propTypes = {
  name: string,
  hostname: string,
  selfAddress: string,
  servicesCount: number,
  version: string,
}

InwayRow.defaultProps = {
  servicesCount: 0,
}

const InwaysPage = ({ getInways }) => {
  const { t } = useTranslation()
  const { isReady, error, result } = usePromise(getInways)

  return (
    <PageTemplate>
      <PageTemplate.Header
        title={t('Inways')}
        description={t('Gateways to provide services.')}
      />

      {!isReady ? (
        <LoadingMessage />
      ) : error ? (
        <Alert variant="error" data-testid="error-message">
          {t('Failed to load the inways.')}
        </Alert>
      ) : result != null && result.length === 0 ? (
        <EmptyContentMessage>
          {t('There are no inways registered yet.')}
        </EmptyContentMessage>
      ) : result ? (
        <Table data-testid="inways-list" role="grid">
          <thead>
            <Table.TrHead>
              <Table.Th>{t('Type')}</Table.Th>
              <Table.Th>{t('Name')}</Table.Th>
              <Table.Th>{t('Hostname')}</Table.Th>
              <Table.Th>{t('Self address')}</Table.Th>
              <Table.Th>{t('Services')}</Table.Th>
              <Table.Th>{t('Version')}</Table.Th>
            </Table.TrHead>
          </thead>
          <tbody>
            {result.map((inway, i) => (
              <InwayRow
                name={inway.name}
                hostname={inway.hostname}
                selfAddress={inway.selfAddress}
                servicesCount={inway.services ? inway.services.length : 0}
                version={inway.version}
                key={i}
              />
            ))}
          </tbody>
        </Table>
      ) : null}
    </PageTemplate>
  )
}

InwaysPage.propTypes = {
  getInways: func,
}

InwaysPage.defaultProps = {
  getInways: InwayRepository.getAll,
}

export default InwaysPage
