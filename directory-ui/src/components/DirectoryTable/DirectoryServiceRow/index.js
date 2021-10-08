// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import Table from '../Table'
import StateIndicator from '../StateIndicator'

const DirectoryServiceRow = ({ service, ...props }) => {
  const { apiType, name, organization, status } = service

  return (
    <Table.Tr
      to={`/${organization.serialNumber}/${name}`}
      name={`${organization.name} - ${name}`}
      data-testid="directory-service-row"
      {...props}
    >
      <Table.MobileTd>
        <StateIndicator state={status} />
        <Table.MobileTdContent>
          <p>{organization.name}</p>
          <p>{name}</p>
          {apiType && <p>{apiType}</p>}
        </Table.MobileTdContent>
      </Table.MobileTd>

      <Table.Td>{organization.name}</Table.Td>
      <Table.Td>{name}</Table.Td>
      <Table.Td>
        <StateIndicator state={status} />
      </Table.Td>
      <Table.Td>{apiType}</Table.Td>
    </Table.Tr>
  )
}

DirectoryServiceRow.propTypes = {
  service: shape({
    apiType: string,
    name: string.isRequired,
    organization: shape({
      name: string.isRequired,
      serialNumber: string.isRequired,
    }).isRequired,
    status: string.isRequired,
  }),
}

export default DirectoryServiceRow
