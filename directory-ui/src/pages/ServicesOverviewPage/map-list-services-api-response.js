// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

const VALID_STATES = ['unknown', 'up', 'down']

export const reduceInwayStatesToStatus = (inways = []) => {
    let status = 'down'

    // Create an array of unique inway lowercase states
    const states = inways
        .map((inway) => {
            const state = inway.state.toLowerCase()
            if (VALID_STATES.includes(state) === false) {
                return 'unknown'
            }
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

export const mapListServicesAPIResponse = (response) =>
    response && response.services
        ? response.services.map((service) => ({
              /* eslint-disable camelcase */
              organization: service.organization_name,
              name: service.service_name,
              status: reduceInwayStatesToStatus(service.inways),
              apiType: service.api_specification_type,
              contactEmailAddress: service.public_support_contact,
              /* eslint-enable camelcase */
          }))
        : []
