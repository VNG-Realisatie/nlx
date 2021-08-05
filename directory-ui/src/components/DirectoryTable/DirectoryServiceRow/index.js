// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, string } from 'prop-types'
import Table from '../Table'
import StateIndicator from '../StateIndicator'

const DirectoryServiceRow = ({ service, ...props }) => {
  const {
    apiType,
    // contactEmailAddress,
    // documentationUrl,
    name,
    organization,
    status,
  } = service

  return (
    <Table.Tr
      to={`/directory/${organization}/${name}`}
      name={`${organization} - ${name}`}
      data-testid="directory-service-row"
      {...props}
    >
      <Table.Td>{organization}</Table.Td>
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
    contactEmailAddress: string.isRequired,
    documentationUrl: string,
    name: string.isRequired,
    organization: string.isRequired,
    status: string.isRequired,
  }),
}

export default DirectoryServiceRow
