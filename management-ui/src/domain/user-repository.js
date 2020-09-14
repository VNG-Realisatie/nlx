// Copyright © VNG Realisatie 2020
// Licensed under the EUPL
//

import env from '../env'
import { fetchWithoutCaching, throwOnError } from './fetch-utils'

class UserRepository {
  static async getAuthenticatedUser() {
    const response = await fetchWithoutCaching(`${env.oidcBaseUrl}/me`)

    throwOnError(response)

    return await response.json()
  }
}

export default UserRepository
