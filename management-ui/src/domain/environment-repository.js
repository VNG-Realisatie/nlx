// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import { fetchWithCaching } from './fetch-utils'

class EnvironmentRepository {
  static async getCurrent() {
    const result = await fetchWithCaching('/api/v1/environment')

    if (!result.ok) {
      throw new Error('unable to handle the request')
    }

    return await result.json()
  }
}

export default EnvironmentRepository
