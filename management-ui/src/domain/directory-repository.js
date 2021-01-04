// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//
import { fetchWithoutCaching, throwOnError } from './fetch-utils'

class DirectoryRepository {
  static async getAll() {
    const response = await fetchWithoutCaching(`/api/v1/directory/services`)

    throwOnError(response)

    const result = await response.json()
    return result.services || []
  }
}

export default DirectoryRepository
