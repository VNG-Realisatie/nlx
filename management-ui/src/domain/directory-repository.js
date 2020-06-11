// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { fetchWithoutCaching } from './fetch-utils'

class DirectoryRepository {
  static async getAll() {
    const result = await fetchWithoutCaching(`/api/v1/directory/services`)

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    const response = await result.json()
    response.services = response.services || []

    return response
  }

  static async getByName(organizationName, serviceName) {
    const result = await fetchWithoutCaching(
      `/api/v1/directory/organizations/${organizationName}/services/${serviceName}`,
    )

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

  static async requestAccess(organizationName, serviceName) {
    const result = await fetch(
      `/api/v1/directory/organizations/${organizationName}/services/${serviceName}/access-requests`,
      {
        method: 'POST',
      },
    )

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
}

export default DirectoryRepository
