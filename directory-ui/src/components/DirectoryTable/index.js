// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, shape, string, bool } from 'prop-types'
import EmptyContentMessage from '../EmptyContentMessage'
import SearchSummary from '../SearchSummary'
import Table from './Table'
import DirectoryServiceRow from './DirectoryServiceRow'

const filterServicesByOnlineStatus = (services) => {
  return services.filter((service) => service.status === 'up')
}

const filterServicesByQuery = (services, query) => {
  return services.filter(
    (service) =>
      service.organization.toLowerCase().includes(query.toLowerCase()) ||
      service.name.toLowerCase().includes(query.toLowerCase()),
  )
}

const filterServices = (services, query, filterByOnlineServices) => {
  const result = filterByOnlineServices
    ? filterServicesByOnlineStatus(services)
    : services

  return query ? filterServicesByQuery(result, query) : result
}

const DirectoryTable = ({
  services,
  selectedServiceName,
  filterQuery,
  filterByOnlineServices,
}) => {
  const filteredServices = filterServices(
    services,
    filterQuery,
    filterByOnlineServices,
  )

  return filteredServices.length ? (
    <>
      <SearchSummary
        totalServices={services.length}
        totalFilteredServices={filteredServices.length}
      />
      <Table withLinks role="grid" data-testid="directory-services">
        <Table.Thead>
          <Table.TrHead>
            <Table.Th>Organisatie</Table.Th>
            <Table.Th>Service</Table.Th>
            <Table.Th>Status</Table.Th>
            <Table.Th>API Type</Table.Th>
          </Table.TrHead>
        </Table.Thead>
        <tbody>
          {filteredServices.map((service) => (
            <DirectoryServiceRow
              key={`${service.organization}-${service.name}`}
              service={service}
              selected={service.name === selectedServiceName}
            />
          ))}
        </tbody>
      </Table>
    </>
  ) : (
    <EmptyContentMessage data-testid="directory-no-services">
      Geen services gevonden
    </EmptyContentMessage>
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
      status: string,
    }),
  ).isRequired,
  selectedServiceName: string,
  filterQuery: string,
  filterByOnlineServices: bool,
}

DirectoryTable.defaultProps = {
  filterQuery: '',
  filterByOnlineServices: false,
}

export default DirectoryTable
