// Copyright © VNG Realisatie 2021
// Licensed under the EUPL
//
import { mapListServicesAPIResponse } from './map-list-services-api-response'

const getServices = async () => {
  try {
    const response = await fetch(`/api/directory/list-services`, {
      headers: {
        'Content-Type': 'application/json',
      },
    })
    const services = await response.json()
    return mapListServicesAPIResponse(services)
  } catch (e) {
    console.error('error fetching services: ', e)
    throw e
  }
}

export default getServices
