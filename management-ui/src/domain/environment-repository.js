// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { fetchWithCaching, throwOnError } from './fetch-utils'

class EnvironmentRepository {
  static async getCurrent() {
    const response = await fetchWithCaching('/api/v1/environment')

    throwOnError(response)

    return await response.json()
  }
}

export default EnvironmentRepository
