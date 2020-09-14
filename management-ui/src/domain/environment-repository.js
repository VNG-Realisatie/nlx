// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import env from '../env'
import { fetchWithCaching, throwOnError } from './fetch-utils'

class EnvironmentRepository {
  static async getCurrent() {
    const response = await fetchWithCaching(
      `${env.managementApiBaseUrl}/v1/environment`,
    )

    throwOnError(response)

    return await response.json()
  }
}

export default EnvironmentRepository
