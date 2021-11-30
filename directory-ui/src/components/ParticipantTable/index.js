// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import React from 'react'
import { arrayOf, shape, string, number, instanceOf } from 'prop-types'
import EmptyContentMessage from '../EmptyContentMessage'
import SearchSummary from '../SearchSummary'
import Table from './Table'
import ParticipantRow from './ParticipantRow'

const filterParticipantsByQuery = (participants, query) => {
  return participants.filter(
    (participant) =>
      participant.organization.name
        .toLowerCase()
        .includes(query.toLowerCase()) ||
      participant.organization.serialNumber
        .toLowerCase()
        .includes(query.toLowerCase()),
  )
}

const filterParticipants = (participants, query) => {
  return query ? filterParticipantsByQuery(participants, query) : participants
}

const ParticipantsTable = ({ participants, filterQuery }) => {
  const filteredParticipants = filterParticipants(participants, filterQuery)

  return filteredParticipants.length ? (
    <>
      <SearchSummary
        totalItems={participants.length}
        totalFilteredItems={filteredParticipants.length}
        itemDescription="deelnemer"
        itemPluralDescription="deelnemers"
      />
      <Table withLinks role="grid" data-testid="directory-participants">
        <Table.Thead>
          <Table.TrHead>
            <Table.Th>Organisatie</Table.Th>
            <Table.Th>Deelnemer sinds</Table.Th>
            <Table.Th>Services</Table.Th>
            <Table.Th>Inways</Table.Th>
            <Table.Th>Outways</Table.Th>
          </Table.TrHead>
        </Table.Thead>
        <tbody>
          {filteredParticipants.map((participant) => (
            <ParticipantRow
              key={`${participant.organization.serialNumber}`}
              participant={participant}
            />
          ))}
        </tbody>
      </Table>
    </>
  ) : (
    <EmptyContentMessage data-testid="directory-no-participants">
      Geen deelnemers gevonden
    </EmptyContentMessage>
  )
}

ParticipantsTable.propTypes = {
  participants: arrayOf(
    shape({
      organization: shape({
        name: string.isRequired,
        serialNumber: string.isRequired,
      }).isRequired,
      createdAt: instanceOf(Date),
      serviceCount: number.isRequired,
      inwayCount: number.isRequired,
      outwayCount: number.isRequired,
    }),
  ).isRequired,
  filterQuery: string,
}

ParticipantsTable.defaultProps = {
  filterQuery: '',
}

export default ParticipantsTable
