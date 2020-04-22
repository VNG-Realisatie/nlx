// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

class ServiceRepository {
  static async getAll() {
    const result = await fetch(`/api/v1/services`)

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    const response = await result.json()
    return response.services ? response.services : []
  }

  static async create(service) {
    const result = await fetch('/api/v1/services', {
      method: 'POST',
      body: JSON.stringify(service),
    })

    if (result.status === 400) {
      throw new Error('invalid user input')
    }

    if (result.status === 403) {
      throw new Error('forbidden')
    }

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    return result.json()
  }

  static async remove(service) {
    const result = await fetch(`/api/v1/services/${service.name}`, {
      method: 'DELETE',
    })

    if (result.status === 404) {
      throw new Error('not found')
    }

    if (result.status === 403) {
      throw new Error('forbidden')
    }

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    return result.json()
  }

  static async getByName(name) {
    const result = await fetch(`/api/v1/services/${name}`)

    if (result.status === 404) {
      throw new Error('not found')
    }
    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    const service = await result.json()
    service.internal = !!service.internal
    service.inways = service.inways || []
    service.authorizationSettings.authorizations =
      service.authorizationSettings.authorizations || []

    return service
  }
}

export default ServiceRepository
