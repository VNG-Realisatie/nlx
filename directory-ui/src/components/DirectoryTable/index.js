// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, shape, string } from 'prop-types'
import EmptyContentMessage from '../EmptyContentMessage'
import Table from './Table'
import DirectoryServiceRow from './DirectoryServiceRow'

const DirectoryTable = ({ services, selectedServiceName }) => {
  return services.length === 0 ? (
    <EmptyContentMessage data-testid="directory-no-services">
      There are no services yet
    </EmptyContentMessage>
  ) : (
    <Table withLinks role="grid" data-testid="directory-services">
      <thead>
        <Table.TrHead>
          <Table.Th>Organization</Table.Th>
          <Table.Th>Service</Table.Th>
          <Table.Th>State</Table.Th>
          <Table.Th>API Type</Table.Th>
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

DirectoryTable.propTypes = {
  services: arrayOf(
    shape({
      apiType: string,
      contactEmailAddress: string,
      documentationUrl: string,
      name: string.isRequired,
      organization: string.isRequired,
      status: string.isRequired,
    }),
  ).isRequired,
  selectedServiceName: string,
}

export default DirectoryTable
