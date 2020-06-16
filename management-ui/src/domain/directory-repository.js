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
}

export default DirectoryRepository
