// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

const STATE_PREFIX = 'STATE_'

export const reduceInwayStatesToStatus = (inways = []) => {
  let status = 'down'

  // Create an array of unique inway lowercase states
  const states = inways
    .map((inway) => {
      const state = inway.state.replace(STATE_PREFIX, '').toLowerCase()

      if (state === 'unspecified') return 'unknown'

      return state
    })
    .filter((element, index, array) => array.indexOf(element) === index)

  if (states.length === 1) {
    status = states[0]
  } else if (states.length > 1) {
    status = 'degraded'
  }

  return status
}

export const mapListServicesAPIResponse = (response) => {
  if (response.services) {
    return response.services.map((service) => ({
      organization: {
        name: service.organization.name,
        serialNumber: service.organization.serial_number || '',
      },
      name: service.name,
      status: reduceInwayStatesToStatus(service.inways),
      apiType: service.api_specification_type,
      contactEmailAddress: service.public_support_contact,
      documentationUrl: service.documentation_url,
      oneTimeCosts: (service.costs?.one_time || 0) / 100,
      monthlyCosts: (service.costs?.monthly || 0) / 100,
      requestCosts: (service.costs?.request || 0) / 100,
    }))
  }
  return []
}
