// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, shape, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../../../components/Table'
import EmptyContentMessage from '../../../../../components/EmptyContentMessage'
import DirectoryServiceRow from '../DirectoryServiceRow'

const DirectoryPageView = ({ services, selectedServiceName }) => {
  const { t } = useTranslation()

  return services.length === 0 ? (
    <EmptyContentMessage data-testid="directory-no-services">
      {t('There are no services yet')}
    </EmptyContentMessage>
  ) : (
    <Table withLinks role="grid" data-testid="directory-services">
      <thead>
        <Table.TrHead>
          <Table.Th>{t('Organization')}</Table.Th>
          <Table.Th>{t('Service')}</Table.Th>
          <Table.Th>{t('State')}</Table.Th>
          <Table.Th>{t('API Type')}</Table.Th>
          <Table.Th>{t('Access')}</Table.Th>
        </Table.TrHead>
      </thead>
      <tbody>
        {services.map((service) => (
          <DirectoryServiceRow
            key={`${service.organizationName}-${service.serviceName}`}
            service={service}
            selected={service.serviceName === selectedServiceName}
          />
        ))}
      </tbody>
    </Table>
  )
}

DirectoryPageView.propTypes = {
  services: arrayOf(
    shape({
      organizationName: string.isRequired,
      serviceName: string.isRequired,
    }),
  ).isRequired,
  selectedServiceName: string,
}

export default DirectoryPageView
