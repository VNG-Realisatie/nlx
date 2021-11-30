// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

export const mapListParticipantsAPIResponse = (response) => {
  if (response.participants) {
    return response.participants.map((participant) => ({
      organization: {
        name: participant.organization.name,
        serialNumber: participant.organization.serial_number || '',
      },
      createdAt: new Date(participant.createdAt),
      inwayCount: participant.statistics.inways,
      outwayCount: participant.statistics.outways,
      serviceCount: participant.statistics.services,
    }))
  }
  return []
}
