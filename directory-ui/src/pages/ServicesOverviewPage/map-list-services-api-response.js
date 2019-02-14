export const mapListServicesAPIResponse = response =>
  response && response.services ?
    response.services.map(service => ({
      organization: service['organization_name'],
      name: service['service_name'],
      status: service['inway_addresses'] ? 'online' : 'offline',
      documentationLink: service['documentation_url'],
      apiType: service['api_specification_type']
    })) : []
