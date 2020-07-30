// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, shape } from 'prop-types'
import { Table } from '@commonground/design-system'
import { useTranslation } from 'react-i18next'
import pick from 'lodash.pick'

import { directoryServicePropTypes } from '../../../../../models/DirectoryServiceModel'
import EmptyContentMessage from '../../../../../components/EmptyContentMessage'
import DirectoryServiceRow from '../DirectoryServiceRow'

const DirectoryPageView = ({ services }) => {
  const { t } = useTranslation()

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
          />
        ))}
      </tbody>
    </Table>
  )
}

DirectoryPageView.propTypes = {
  services: arrayOf(
    shape(pick(directoryServicePropTypes, ['organizationName', 'serviceName'])),
  ).isRequired,
}

export default DirectoryPageView
