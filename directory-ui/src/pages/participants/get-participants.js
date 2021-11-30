// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { mapListParticipantsAPIResponse } from './map-list-participants-api-response'

const getParticipants = async () => {
  try {
    const response = await fetch(`/api/directory/participants`, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    const participants = await response.json()
    return mapListParticipantsAPIResponse(participants)
  } catch (e) {
    console.error('error fetching participants: ', e)
    throw e
  }
}

export default getParticipants
