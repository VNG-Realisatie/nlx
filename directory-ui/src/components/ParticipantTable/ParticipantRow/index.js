// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { shape, number, string, instanceOf } from 'prop-types'
import dayjs from 'dayjs'
import Table from '../Table'

const ParticipantRow = ({ participant, ...props }) => {
  const { organization, createdAt, serviceCount, inwayCount, outwayCount } =
    participant

  return (
    <Table.Tr
      name={`${organization.name}`}
      data-testid="directory-participant-row"
      {...props}
    >
      <Table.Td>{organization.name}</Table.Td>
      <Table.Td>{dayjs(createdAt).format('D MMMM YYYY')}</Table.Td>
      <Table.Td>{serviceCount}</Table.Td>
      <Table.Td>{inwayCount}</Table.Td>
      <Table.Td>{outwayCount}</Table.Td>
    </Table.Tr>
  )
}

ParticipantRow.propTypes = {
  participant: shape({
    organization: shape({
      name: string.isRequired,
      serialNumber: string.isRequired,
    }).isRequired,
    createdAt: instanceOf(Date),
    serviceCount: number.isRequired,
    inwayCount: number.isRequired,
    outwayCount: number.isRequired,
  }),
}

export default ParticipantRow
