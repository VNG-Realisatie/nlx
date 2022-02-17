// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { useTranslation } from 'react-i18next'
import { arrayOf, instanceOf, string } from 'prop-types'
import Table from '../../../../../components/Table'
import EmptyContentMessage from '../../../../../components/EmptyContentMessage'
import DirectoryServiceRow from '../DirectoryServiceRow'
import DirectoryServiceModel from '../../../../../stores/models/DirectoryServiceModel'

const DirectoryPageView = ({
  managementSubjectSerialNumber,
  services,
  selectedServiceName,
}) => {
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
            key={`${service.organization.serialNumber}-${service.serviceName}`}
            service={service}
            selected={service.serviceName === selectedServiceName}
            ownService={
              service.organization.serialNumber ===
              managementSubjectSerialNumber
            }
          />
        ))}
      </tbody>
    </Table>
  )
}

DirectoryPageView.propTypes = {
  managementSubjectSerialNumber: string,
  services: arrayOf(instanceOf(DirectoryServiceModel)),
  selectedServiceName: string,
}

export default DirectoryPageView
