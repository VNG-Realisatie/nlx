// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { fetchWithoutCaching, throwOnError } from './fetch-utils'

class SettingsRepository {
  static async get() {
    const response = await fetchWithoutCaching(`/api/v1/settings`)

    throwOnError(response)

    let result = await response.json()
    result = result || {}

    if (typeof result.organizationInway === 'undefined') {
      result.organizationInway = ''
    }

    return result
  }

  static async update(settings) {
    if (typeof settings.organizationInway === 'undefined') {
      throw new Error(
        'The setting organizationInway must be specified',
      )
    }

    const response = await fetch(`/api/v1/settings`, {
      method: 'PUT',
      body: JSON.stringify(settings),
    })

    throwOnError(response)

    return null
  }
}

export default SettingsRepository
