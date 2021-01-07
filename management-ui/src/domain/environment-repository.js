// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { throwOnError } from './fetch-utils'

class EnvironmentRepository {
  static async getCurrent() {
    const response = await fetch('/api/v1/environment')

    throwOnError(response)

    return await response.json()
  }
}

export default EnvironmentRepository
