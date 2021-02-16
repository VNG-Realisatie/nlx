// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array, string } from 'prop-types'
import { useTranslation } from 'react-i18next'
import Table from '../../../../components/Table'
import EmptyContentMessage from '../../../../components/EmptyContentMessage'
import ServiceRow from './ServiceRow'

const ServicesPageView = ({ services, selectedServiceName }) => {
  const { t } = useTranslation()

  return services.length === 0 ? (
    <EmptyContentMessage>{t('There are no services yet')}</EmptyContentMessage>
  ) : (
    <Table withLinks data-testid="services-list" role="grid">
      <thead>
        <Table.TrHead>
          <Table.Th>{t('Name')}</Table.Th>
          <Table.Th />
        </Table.TrHead>
      </thead>
      <tbody>
        {services.map((service, i) => (
          <ServiceRow
            service={service}
            key={i}
            selected={service.name === selectedServiceName}
          />
        ))}
      </tbody>
    </Table>
  )
}

ServicesPageView.propTypes = {
  services: array,
  selectedServiceName: string,
}

export default ServicesPageView
