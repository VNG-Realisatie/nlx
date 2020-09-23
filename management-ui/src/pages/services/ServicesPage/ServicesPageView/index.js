// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { array } from 'prop-types'
import { useTranslation } from 'react-i18next'
import { Table } from '@commonground/design-system'

import EmptyContentMessage from '../../../../components/EmptyContentMessage'
import ServiceRow from './ServiceRow'

const ServicesPageView = ({ services }) => {
  const { t } = useTranslation()

  return services.length === 0 ? (
    <EmptyContentMessage>{t('There are no services yet.')}</EmptyContentMessage>
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
          <ServiceRow service={service} key={i} />
        ))}
      </tbody>
    </Table>
  )
}

ServicesPageView.propTypes = {
  services: array,
}

export default ServicesPageView
