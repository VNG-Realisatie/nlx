// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { useTranslation } from 'react-i18next'
import { arrayOf, shape, string } from 'prop-types'
import React from 'react'
import { Table } from '@commonground/design-system'

import EmptyContentMessage from '../../../components/EmptyContentMessage'
import StatusIndicator from '../../../components/StatusIndicator'

const DirectoryServices = ({ directoryServices }) => {
  const { t } = useTranslation()
  const { services = [] } = directoryServices

  return services.length === 0 ? (
    <EmptyContentMessage data-testid="directory-no-services">
      {t('There are no services yet.')}
    </EmptyContentMessage>
  ) : (
    <Table withLinks role="grid" data-testid="directory-services">
      <thead>
        <Table.TrHead>
          <Table.Th>{t('Organization')}</Table.Th>
          <Table.Th>{t('Service')}</Table.Th>
          <Table.Th>{t('Status')}</Table.Th>
          <Table.Th>{t('API Type')}</Table.Th>
        </Table.TrHead>
      </thead>
      <tbody>
        {services.map((service, i) => {
          const {
            organizationName,
            serviceName,
            status,
            apiSpecificationType,
          } = service

          return (
            <Table.Tr
              key={i}
              to={`/directory/${organizationName}/${serviceName}`}
              name={`${organizationName} - ${serviceName}`}
              data-testid={`directory-service-row-${i}`}
            >
              <Table.Td>{organizationName}</Table.Td>
              <Table.Td>{serviceName}</Table.Td>
              <Table.Td>
                <StatusIndicator status={status} />
              </Table.Td>
              <Table.Td>{apiSpecificationType}</Table.Td>
            </Table.Tr>
          )
        })}
      </tbody>
    </Table>
  )
}

DirectoryServices.propTypes = {
  directoryServices: shape({
    services: arrayOf(
      shape({
        organizationName: string.isRequired,
        serviceName: string.isRequired,
        status: string.isRequired,
        apiSpecificationType: string,
      }),
    ).isRequired,
  }).isRequired,
}

export default DirectoryServices
