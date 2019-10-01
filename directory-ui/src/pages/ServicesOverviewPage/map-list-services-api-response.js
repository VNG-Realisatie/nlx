// Copyright © VNG Realisatie 2018
// Licensed under the EUPL

export const mapListServicesAPIResponse = response =>
  response && response.services ?
    response.services.map(service => ({
      organization: service['organization_name'],
      name: service['service_name'],
      status: service['inway_addresses'] ? 'online' : 'offline',
      apiType: service['api_specification_type'],
      contactEmail: service['public_support_contact']
    })) : []
