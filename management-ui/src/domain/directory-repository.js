// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { fetchWithoutCaching, throwOnError } from './fetch-utils'

function ensureLatestAccessRequest(service) {
  service.latestAccessRequest = service.latestAccessRequest || null
  return service
}

class DirectoryRepository {
  static async getAll() {
    const response = await fetchWithoutCaching(`/api/v1/directory/services`)

    throwOnError(response)

    const result = await response.json()
    return result.services || []
  }

  static async getByName(organizationName, serviceName) {
    const response = await fetchWithoutCaching(
      `/api/v1/directory/organizations/${organizationName}/services/${serviceName}`,
    )

    throwOnError(response)

    const result = await response.json()
    return ensureLatestAccessRequest(result)
  }
}

export default DirectoryRepository
