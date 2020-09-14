// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import env from '../env'
import { fetchWithoutCaching, throwOnError } from './fetch-utils'

class ServiceRepository {
  static async getAll() {
    const response = await fetchWithoutCaching(
      `${env.managementApiBaseUrl}/v1/services`,
    )

    if (!response.ok) {
      throw new Error('unable to handle the request')
    }

    const result = await response.json()
    const services = result.services ? result.services : []
    return services.map((service) => {
      service.internal = !!service.internal
      service.inways = service.inways || []
      service.authorizationSettings = service.authorizationSettings || {}
      service.authorizationSettings.authorizations =
        service.authorizationSettings.authorizations || []
      return service
    })
  }

  static async create(service) {
    const response = await fetch(`${env.managementApiBaseUrl}/v1/services`, {
      method: 'POST',
      body: JSON.stringify(service),
    })

    throwOnError(response)

    return response.json()
  }

  static async update(name, service) {
    if (name !== service.name) {
      throw new Error('Changing the service name is not allowed')
    }

    const response = await fetch(
      `${env.managementApiBaseUrl}/v1/services/${name}`,
      {
        method: 'PUT',
        body: JSON.stringify(service),
      },
    )

    throwOnError(response)

    return response.json()
  }

  static async remove(service) {
    const response = await fetch(
      `${env.managementApiBaseUrl}/v1/services/${service.name}`,
      {
        method: 'DELETE',
      },
    )

    throwOnError(response)

    return response.json()
  }

  static async getByName(name) {
    const response = await fetchWithoutCaching(
      `${env.managementApiBaseUrl}/v1/services/${name}`,
    )

    throwOnError(response)

    const service = await response.json()
    service.internal = !!service.internal
    service.inways = service.inways || []
    service.authorizationSettings.authorizations =
      service.authorizationSettings.authorizations || []

    return service
  }
}

export default ServiceRepository
