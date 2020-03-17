// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

// API specification
// https://gitlab.com/commonground/nlx/nlx/-/blob/master/config-api/configapi/configapi.swagger.json

class ServiceRepository {
  static async getAll() {
    const result = await fetch(`/api/v1/services`)

    if (!result.ok) {
      throw new Error('failed to get all services')
    }

    const response = await result.json()
    return response.services
  }
}

export default ServiceRepository
