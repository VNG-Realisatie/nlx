// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

class AccessGrantRepository {
  static async getByService(serviceName) {
    const response = await fetch(
      `/api/v1/access-grants/services/${serviceName}`,
    )

    throwOnError(response)

    const result = await response.json()
    return result.accessGrants || []
  }
}

export default AccessGrantRepository
