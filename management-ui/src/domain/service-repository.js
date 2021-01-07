// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

class ServiceRepository {
  static async create(service) {
    const response = await fetch('/api/v1/services', {
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

    const response = await fetch(`/api/v1/services/${name}`, {
      method: 'PUT',
      body: JSON.stringify(service),
    })

    throwOnError(response)

    return response.json()
  }

  static async remove(service) {
    const response = await fetch(`/api/v1/services/${service.name}`, {
      method: 'DELETE',
    })

    throwOnError(response)

    return response.json()
  }
}

export default ServiceRepository
