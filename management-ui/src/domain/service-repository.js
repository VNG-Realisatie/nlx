// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

class ServiceRepository {
  static async getAll() {
    const result = await fetch(`/api/v1/services`, {
      // needed to prevent caching on IE 11
      // https://stackoverflow.com/questions/37755782/prevent-ie11-caching-get-call-in-angular-2/44561162#44561162
      headers: {
        'Cache-Control': 'no-cache',
        Pragma: 'no-cache',
        Expires: 'Sat, 01 Jan 2000 00:00:00 GMT',
      },
    })

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

  static async update(name, service) {
    if (name !== service.name) {
      throw new Error('Changing the service name is not allowed')
    }

    const result = await fetch(`/api/v1/services/${name}`, {
      method: 'PUT',
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
    const result = await fetch(`/api/v1/services/${name}`, {
      headers: {
        'Cache-Control': 'no-cache',
        Pragma: 'no-cache',
        Expires: 'Sat, 01 Jan 2000 00:00:00 GMT',
      },
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

    const service = await result.json()
    service.internal = !!service.internal
    service.inways = service.inways || []
    service.authorizationSettings.authorizations =
      service.authorizationSettings.authorizations || []

    return service
  }
}

export default ServiceRepository
